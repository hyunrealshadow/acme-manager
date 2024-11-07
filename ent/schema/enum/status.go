package enum

import "io"

type Status string

const (
	Pending   Status = "Pending"
	Succeeded Status = "Succeeded"
	Failed    Status = "Failed"
)

// Values provides list valid values for Enum.
func (Status) Values() (kinds []string) {
	for _, s := range []Status{Pending, Succeeded, Failed} {
		kinds = append(kinds, string(s))
	}
	return
}

// UnmarshalGQL Make entgo generated code happy
func (s Status) UnmarshalGQL(v any) error {
	panic("implement me")
}

// MarshalGQL Make entgo generated code happy
func (s Status) MarshalGQL(w io.Writer) {
	panic("implement me")
}
