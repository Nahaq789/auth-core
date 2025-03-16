package auth

type VerifyCode struct {
	code string
}

func NewVerifyCode(code string) *VerifyCode {
	return &VerifyCode{code: code}
}

func (c *VerifyCode) Code() string {
	return c.code
}
