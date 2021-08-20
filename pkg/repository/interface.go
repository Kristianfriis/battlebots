package repository

import "battlebots/pkg/models/completeRobot"

type Repository interface {
	GetDatabase() *interface{}
	// GetAllParts()
	// GetSpecificPart(id string)
	CreateRobot(robo completeRobot.Robot) (success bool, err error)
	GetAllRobots() ([]completeRobot.Robot, error)
	GetSpecificRobot(id string) (completeRobot.Robot, error)
}
