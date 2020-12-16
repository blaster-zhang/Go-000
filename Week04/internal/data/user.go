package data

import "Go000/Week04/internal/biz"

type UserRepo struct {
}

func (u *UserRepo) Add(user biz.User) error {
	//插入数据库
	return nil
}
