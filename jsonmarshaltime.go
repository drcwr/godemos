package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*

/usr/lib/go15/src/encoding/json/encode.go:223
// Marshaler is the interface implemented by types that
// can marshal themselves into valid JSON.
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

time.Time has MarshalJSON

/usr/lib/go15/src/time/time.go:1238
// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (t Time) MarshalJSON() ([]byte, error)

*/

func main() {
	t := struct {
		N int
		time.Time
		// T time.Time
	}{
		5,
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
	}

	m, _ := json.Marshal(t)

	fmt.Printf("%s\n", m)
}

/*

"2020-12-20T00:00:00Z"

*/
