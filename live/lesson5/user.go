package lesson5

type UserClone interface {
	Clone()
}

type User struct{}

func (u *User) Clone() {

}
