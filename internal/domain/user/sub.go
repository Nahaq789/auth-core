package user

type Sub struct {
	value string
}

func (s Sub) Value() string {
	return s.value
}

func NewSub(sub string) *Sub {
	return &Sub{value: sub}
}

func (s Sub) String() string {
	return s.value
}
