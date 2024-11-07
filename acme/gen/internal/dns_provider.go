package internal

import (
	"fmt"
	"github.com/BurntSushi/toml"
	mapset "github.com/deckarep/golang-set/v2"
	"go/ast"
	"go/constant"
	"go/types"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	packagePath              = "github.com/go-acme/lego/v4/providers/dns"
	newDefaultConfigFuncName = "NewDefaultConfig"
	newDNSProviderFuncName   = "NewDNSProvider"
)

var needIgnoreFunctions = map[string]mapset.Set[string]{
	"oraclecloud": mapset.NewSet("newConfigProvider"),
}

type DnsProvider struct {
	Name                 string
	DisplayName          string
	Description          string
	Imports              []string
	PkgPath              string
	EnvConstants         map[string]any
	JsonSchema           string
	NewDefaultConfigFunc string
	NewDNSProviderFunc   string
	ExtraFunctions       map[string]string
	CredentialsFields    []string
	AdditionalFields     []string
	dependencyFunctions  mapset.Set[string]
}

func GetAllDnsProviderPackages() []string {
	cfg := &packages.Config{
		Mode: packages.LoadImports,
	}
	pkgs, err := packages.Load(cfg, packagePath)
	if err != nil {
		logger.Fatal(err)
	}
	if len(pkgs) == 0 {
		logger.Fatalf("package %s not found", packagePath)
	}
	pkgImports := pkgs[0].Imports
	dnsProviders := make([]string, 0)
	for _, pkg := range pkgImports {
		if strings.HasPrefix(pkg.PkgPath, packagePath) {
			dnsProviders = append(dnsProviders, pkg.PkgPath)
		}
	}
	sort.Strings(dnsProviders)
	return dnsProviders
}

func getImport(pkgName string, imp *ast.ImportSpec, imports map[string]*packages.Package) string {
	if imp.Name != nil {
		if imp.Name.Name == pkgName {
			return ""
		}
		return imp.Name.Name + " " + imp.Path.Value
	}
	importPathValue := imp.Path.Value
	importPathContent := importPathValue[1 : len(importPathValue)-1]
	importPackage := imports[importPathContent]
	if strings.HasPrefix(importPathContent, "github.com/go-acme/lego/v4") {
		if !strings.Contains(importPathContent, "github.com/go-acme/lego/v4/platform/config/env") &&
			!strings.Contains(importPathContent, "internal") {

			return importPathValue
		}
	} else if !strings.HasSuffix(importPathContent, pkgName) && importPackage != nil && importPackage.Name != pkgName {
		return importPathValue
	}
	return ""
}

func getProviderInfo(pkg *packages.Package) *DnsProvider {
	var newDefaultConfigFunc, newDNSProviderFunc string
	var hasConfigStruct = false
	constants := make(map[string]any)
	importSets := mapset.NewSet[string]()
	funcMap := make(map[string]string)
	for _, s := range pkg.Syntax {
		for _, decl := range s.Decls {
			for _, imp := range s.Imports {
				impStr := getImport(pkg.Name, imp, pkg.Imports)
				if impStr != "" {
					importSets.Add(impStr)
				}
			}
			if fn, isFn := decl.(*ast.FuncDecl); isFn {
				fnName := fn.Name.Name
				pos := pkg.Fset.Position(fn.Pos())
				end := pkg.Fset.Position(fn.End())
				data, err := os.ReadFile(pos.Filename)
				if err != nil {
					continue
				}
				funcMap[fnName] = string(data[pos.Offset:end.Offset])
			}
			if genDecl, isGenDecl := decl.(*ast.GenDecl); isGenDecl {
				for _, spec := range genDecl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if _, ok := typeSpec.Type.(*ast.StructType); ok {
							if typeSpec.Name.Name == "Config" {
								hasConfigStruct = true
							}
						}
					}
				}
			}
		}
		scope := pkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			if obj, ok := obj.(*types.Const); ok {
				objName := obj.Name()
				objVal := obj.Val()
				switch objVal.Kind() {
				case constant.String:
					constants[objName] = objVal.String()
				case constant.Int:
					intVal, ok := constant.Int64Val(objVal)
					if ok {
						constants[objName] = intVal
					} else {
						logger.Warnf("Could not convert constant %s to int", objName)
					}
				case constant.Bool:
					boolVal := constant.BoolVal(objVal)
					constants[objName] = boolVal
				default:
					logger.Warnf("Unsupported constant type: %v", objVal.Kind())
				}
			}
		}
	}

	for funcName, funcCode := range funcMap {
		if funcName == newDefaultConfigFuncName {
			newDefaultConfigFunc = funcCode
			delete(funcMap, funcName)
		} else if funcName == newDNSProviderFuncName {
			newDNSProviderFunc = funcCode
			delete(funcMap, funcName)
		}
	}

	if hasConfigStruct && newDefaultConfigFunc != "" && newDNSProviderFunc != "" {
		goFiles := pkg.GoFiles
		if len(goFiles) == 0 {
			logger.Fatalf("no go files found for package %s", pkg.PkgPath)
		}
		path := filepath.Join(filepath.Dir(goFiles[0]), pkg.Name+".toml")
		docFileBytes, err := os.ReadFile(path)
		if err != nil {
			logger.Fatalf("could not load config in %s: %v", path, err)
		}
		docFile := string(docFileBytes)

		// in hurricane package, I found a typo, so I do this
		docFile = strings.Replace(docFile, "[Configuration.Addtional]", "[Configuration.Additional]", -1)

		m := model{}
		_, err = toml.Decode(docFile, &m)
		if err != nil {
			logger.Fatalf("could not decode config in %s: %v", path, err)
		}

		p := &DnsProvider{
			Name:                 pkg.Name,
			Imports:              importSets.ToSlice(),
			PkgPath:              pkg.PkgPath,
			EnvConstants:         constants,
			NewDefaultConfigFunc: newDefaultConfigFunc,
			NewDNSProviderFunc:   newDNSProviderFunc,
		}
		generateNewDefaultConfigFunc(p)
		generateNewDNSProviderFunc(p)
		processDnsProviderModel(p, m)
		return p
	}
	return nil
}

func GetDnsProvider(pkgPath string) (*DnsProvider, error) {
	cfg := &packages.Config{
		Mode: packages.LoadAllSyntax,
	}
	pkgs, err := packages.Load(cfg, pkgPath)
	if err != nil {
		return nil, err
	}
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("package %s not found", pkgPath)
	}
	pkg := pkgs[0]
	provider := getProviderInfo(pkg)
	if provider == nil {
		return nil, fmt.Errorf("could not get provider info for %s", pkgPath)
	}
	return provider, nil
}
