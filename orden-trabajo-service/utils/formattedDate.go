package utils

import (
	"fmt"
	"time"
)

type FechaPersonalizada time.Time

func (f FechaPersonalizada) MarshalJSON() ([]byte, error) {
	t := time.Time(f)
	formatted := fmt.Sprintf("\"%s\"", t.Format("02/01/2006")) // dd/mm/yyyy
	return []byte(formatted), nil
}

func (f *FechaPersonalizada) UnmarshalJSON(b []byte) error {
	parsed, err := time.Parse(`"02/01/2006"`, string(b))
	if err != nil {
		return err
	}
	*f = FechaPersonalizada(parsed)
	return nil
}

func (f FechaPersonalizada) ToTime() time.Time {
	return time.Time(f)
}
