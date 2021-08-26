package sqllite

import (
	"battlebots/pkg/models/completeRobot"
	"fmt"
	"log"
	"time"
)

func (s *Service) CreateRobot(robo completeRobot.Robot) (success bool, err error) {
	sql := `INSERT INTO completeRobots (id,name,head,body,rightArm,leftArm,legs,createdAt) VALUES(?,?,?,?,?,?,?,?)`
	resp, err := s.DB.Exec(sql, robo.Id, robo.Name, robo.Head, robo.Body, robo.RightArm, robo.LeftArm, robo.Legs, time.Now().String())

	if err != nil {
		return false, err
	}
	last, _ := resp.LastInsertId()
	fmt.Printf("Robot Added, " + fmt.Sprint(last))
	return true, nil
}
func (s *Service) GetAllRobots() ([]completeRobot.Robot, error) {
	robots := []completeRobot.Robot{}
	err := s.DB.Select(&robots, "SELECT * FROM completeRobots")
	if err != nil {
		log.Print("ERROR ", err)
		return []completeRobot.Robot{}, err
	}
	return robots, nil
}
func (s *Service) GetSpecificRobot(id string) (completeRobot.Robot, error) {
	return completeRobot.Robot{}, nil
}
