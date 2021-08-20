package completeRobot

type Robot struct {
	Id        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Head      int64  `json:"head" db:"head"`
	Body      int64  `json:"body" db:"body"`
	RightArm  int64  `json:"rightArm" db:"rightArm"`
	LeftArm   int64  `json:"leftArm" db:"leftArm"`
	Legs      int64  `json:"legs" db:"legs"`
	CreatedAt string `json:"createdAt" db:"createdAt"`
}
