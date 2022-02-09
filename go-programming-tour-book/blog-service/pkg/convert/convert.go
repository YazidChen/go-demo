package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	i, err := strconv.Atoi(s.String())
	return i, err
}

func (s StrTo) MustInt() int {
	i, _ := s.Int()
	return i
}

func (s StrTo) UInt32() (uint32, error) {
	i, err := strconv.Atoi(s.String())
	return uint32(i), err
}

func (s StrTo) MustUInt32() uint32 {
	i, _ := s.UInt32()
	return i
}
