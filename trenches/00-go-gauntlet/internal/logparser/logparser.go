package logparser

import (
	"errors"
	"strings"
)

var ErrMalformed = errors.New("malformed log line")

// Line format: LEVEL|k=v,k=v|message
type Entry struct {
	Level  string
	Fields map[string]string
	Msg    string
}

func (e *Entry) Set(k, v string) {
	e.Fields[k] = v
}

func Parse(line string) (*Entry, error) {
	parts := strings.Split(line, "|")
	e := &Entry{Level: parts[0], Msg: parts[2]}
	if parts[1] != "" {
		e.Fields = map[string]string{}
		for _, kv := range strings.Split(parts[1], ",") {
			pair := strings.SplitN(kv, "=", 2)
			e.Fields[pair[0]] = pair[1]
		}
	}
	return e, nil
}

func ParseAll(lines []string) ([]*Entry, error) {
	var out []*Entry
	for _, l := range lines {
		e, _ := Parse(l)
		out = append(out, e)
	}
	return out, nil
}

func Pointers(entries []Entry) []*Entry {
	var out []*Entry
	for _, e := range entries {
		out = append(out, &e)
	}
	return out
}
