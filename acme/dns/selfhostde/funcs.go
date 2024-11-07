// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/selfhostde
// License: MIT

package selfhostde

import (
	"errors"
	"fmt"
	"github.com/go-acme/lego/v4/providers/dns/selfhostde"
	"strings"
)

func parseRecordsMapping(raw string) (map[string]*selfhostde.Seq, error) {
	raw = strings.ReplaceAll(raw, " ", "")

	if raw == "" {
		return nil, errors.New("empty mapping")
	}

	acc := map[string]*selfhostde.Seq{}

	for {
		index, err := safeIndex(raw, lineSep)
		if err != nil {
			return nil, err
		}

		if index != -1 {
			name, seq, err := parseLine(raw[:index])
			if err != nil {
				return nil, err
			}

			acc[name] = seq

			// Data for the next iteration.
			raw = raw[index+1:]

			continue
		}

		name, seq, errP := parseLine(raw)
		if errP != nil {
			return nil, errP
		}

		acc[name] = seq

		return acc, nil
	}
}

func parseLine(line string) (string, *selfhostde.Seq, error) {
	idx, err := safeIndex(line, recordSep)
	if err != nil {
		return "", nil, err
	}

	if idx == -1 {
		return "", nil, fmt.Errorf("missing %q: %s", recordSep, line)
	}

	name, rawIDs := line[:idx], line[idx+1:]

	var ids []string
	var count int

	for {
		idx, err = safeIndex(rawIDs, recordSep)
		if err != nil {
			return "", nil, err
		}

		if count == 2 {
			return "", nil, fmt.Errorf("too many record IDs for one domain: %s", line)
		}

		if idx != -1 {
			ids = append(ids, rawIDs[:idx])
			count++

			// Data for the next iteration.
			rawIDs = rawIDs[idx+1:]

			continue
		}

		ids = append(ids, rawIDs)

		return name, selfhostde.NewSeq(ids...), nil
	}
}

func safeIndex(v, sep string) (int, error) {
	index := strings.Index(v, sep)
	if index == 0 {
		return 0, fmt.Errorf("first char is %q: %s", sep, v)
	}

	if index == len(v)-1 {
		return 0, fmt.Errorf("last char is %q: %s", sep, v)
	}

	return index, nil
}
