package sub

import (
	"errors"
	"regexp"
)

type Sub struct {
	value string
}

func NewSub(s string) (*Sub, error) {
	if err := validFormat(s); err != nil {
		return nil, err
	}
	return &Sub{value: s}, nil
}

func (s Sub) String() string {
	return s.value
}

func (s Sub) Value() string {
	return s.value
}

func validFormat(uuid string) error {
	pattern := "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	match := re.MatchString(uuid)
	if match {
		return nil
	}
	error := errors.New("Invalid UUID format: must be in the format 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx'")
	return error
}
