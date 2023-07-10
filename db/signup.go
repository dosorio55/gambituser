package db

import (
	"fmt"
	"github.com/dosorio55/gambituser/models"
	"github.com/dosorio55/gambituser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(data models.SignUp) error {
	fmt.Println("Register begin")

	err := DbConnect()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer Db.Close()

	query := fmt.Sprintf("INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%s', '%s', '%s', '%s')", data.Username, data.UserUUID, tools.GetTimeNow())

	fmt.Println("query: ", query)

	//the first parameter is the number of rows affected
	_, err = Db.Exec(query)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Register end")

	return nil
}
