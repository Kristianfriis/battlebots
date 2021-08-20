package main

import (
	"battlebots/pkg/models/completeRobot"
	"battlebots/pkg/repository/sqllite"
	"fmt"
)

func main() {
	repo := sqllite.NewService()

	testRobo := completeRobot.Robot{
		Id:       1,
		Name:     "Test Robo",
		Head:     123456789,
		Body:     789560498,
		RightArm: 180455608,
		LeftArm:  258007955,
		Legs:     668040988,
	}

	success, _ := repo.CreateRobot(testRobo)

	fmt.Println(success)
}
