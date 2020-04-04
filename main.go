package main

import (
	"context"
	"fmt"
	"log"
	"mod/logic"
)




func main() {
	var cfg = logic.Cfg{
		Connection: "mongodb://localhost:27017",
		DB:         "game",
		Collection: "users",
	}

	//local
	err, str:= play(context.TODO(),"6666",cfg)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(str)

}

func play(ctx context.Context, userID string, cfg logic.Cfg) (error, string) {
	conn, err := logic.NewMongoClient(cfg.Connection, cfg.DB)
	if err != nil {
		return err, ""
	}

	defer func() {
		_ = conn.Disconnect()
	}()

	user, err := conn.FindUser(userID, ctx, cfg.Collection)
	if err != nil {
		return err, logic.UserNotExist
	}

	if user.Money == 0 {
		return nil, logic.BalanceZero
	}

	defer func() {
		conn.UpdateUserMoney(user.Id, user.Money, ctx, cfg.Collection)
	}()
	return nil,""

}
