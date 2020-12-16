package biz

type UserUseCase interface {
	Add(user *User) error
	GetOne(userId int) *User
}

type UserInfoRepo interface {
	Add(u User) error
}

func NewUserUseCase(repo UserInfoRepo) *userUseCase {
	return &userUseCase{Repo: repo}
}

type User struct {
	UserId       int
	Name         string
	Introduction string
}

type userUseCase struct {
	Repo UserInfoRepo
}

func (u *userUseCase) Add(userInfo *User) error {
	return u.Repo.Add(*userInfo)
}

func (u *userUseCase) GetOne(userId int) *User {
	return nil
}
