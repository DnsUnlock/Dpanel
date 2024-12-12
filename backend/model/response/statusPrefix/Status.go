package statusPrefix

type StatusPrefix string

const (
	OK       StatusPrefix = "S00"
	NotFound StatusPrefix = "N00"
	ERROR    StatusPrefix = "E01"
)

func (s StatusPrefix) String() string {
	return string(s)
}
