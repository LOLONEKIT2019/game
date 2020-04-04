package logic

type Cfg struct {
	Connection string
	DB         string
	Collection string
}

type User struct {
	Id     string  `json:"_id" bson:"_id"`
	UserId string   `json:"user_id" bson:"user_id"`
	Money  int64 `json:"money" bson:"money"`
}

const BalanceZero = "BalanceZero"
const UserNotExist = "UserNotExist"
