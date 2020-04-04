package logic

import "errors"

func (u *User) UpMoney(value int64) {
	u.Money = u.Money + value
}

func (u *User) DownMoney(value int64) error {
	if u.Money < value {
		return errors.New("Not enough money")
	}

	u.Money = u.Money - value
}
