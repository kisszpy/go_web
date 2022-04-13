package req

import "errors"

type LoginReq struct {
	Username string

	Password string
}

func (l LoginReq) CheckArguments() error {
	if l.Username == "" {
		return errors.New("用户名不能为空")
	}
	if l.Password == "" {
		return errors.New("密码不能为空")
	}
	return nil
}
