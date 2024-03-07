package main

import "fmt"

type Player struct {
	Username string `json:"username"`
	Level    uint   `json:"level"`
	Status   string `json:"status"`
	Class    string `json:"class"`
}

func (p *Player) LevelUp() {
	p.Level++
}

func main() {
	player1 := Player{
		Username: "player1",
		Level:    1,
		Status:   "active",
		Class:    "warrior",
	}

	player1.LevelUp()

	fmt.Println(player1.Level)
}
