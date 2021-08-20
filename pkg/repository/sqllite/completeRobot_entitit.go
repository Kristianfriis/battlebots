package sqllite

import (
	"battlebots/pkg/models/completeRobot"
	"fmt"
	"time"
)

func (s *Service) CreateRobot(robo completeRobot.Robot) (success bool, err error) {
	sql := `INSERT INTO kid (id,name,head,body,rightArm,leftArm,legs,createdAt) VALUES(?,?,?,?,?,?,?,?)`
	resp, err := s.DB.Exec(sql, robo.Id, robo.Name, robo.Head, robo.Body, robo.RightArm, robo.LeftArm, robo.Legs, time.Now().String())

	if err != nil {
		return false, err
	}
	last, _ := resp.LastInsertId()
	fmt.Printf("Robot Added, " + fmt.Sprint(last))
	return true, nil
}
func (s *Service) GetAllRobots() ([]completeRobot.Robot, error) {
	return []completeRobot.Robot{}, nil
}
func (s *Service) GetSpecificRobot(id string) (completeRobot.Robot, error) {
	return completeRobot.Robot{}, nil
}
