package main

import "fmt"

type data struct {
	things [3]*string

	player1_win int8
	playe2_win  int8
	draw        int8

	player1_name *string
	player2_name *string

	player1_input *string
	player2_input *string

	round int8
}

func container(Data *data) {
	initiate(Data)
	set_player(Data)

	var repeat string
	for {
		play_game(Data)

		fmt.Println(*Data.player1_input)
		fmt.Println(*Data.player2_input)

		if !check(Data) {
			fmt.Println("invalid input")
			continue
		}

		is_win(Data)

		fmt.Print("input: ")
		fmt.Scan(&repeat)
		if repeat != "y" {
			break
		}
	}

}

func initiate(Data *data) {
	for i := 0; i < len(Data.things); i++ {
		Data.things[i] = new(string)
	}
	*Data.things[0] = "rock"
	*Data.things[1] = "scissor"
	*Data.things[2] = "papper"

	Data.player1_input = new(string)
	Data.player2_input = new(string)

	Data.player1_name = nil
	Data.player2_name = nil

	Data.player1_win = 0
	Data.playe2_win = 0

	Data.draw = 0
	Data.round = 0
}

func set_player(Data *data) {
	Data.player1_name = new(string)
	Data.player2_name = new(string)

	fmt.Print("player 1 name: ")
	fmt.Scanln(Data.player1_name)

	fmt.Print("player 2 name: ")
	fmt.Scanln(Data.player2_name)
}

func play_game(Data *data) {
	fmt.Print("player 1 input: ")
	fmt.Scan(Data.player1_input)

	fmt.Print("player 2 input: ")
	fmt.Scan(Data.player2_input)
}

func check(Data *data) bool {
	var p1_valid bool = false //long declare
	p2_valid := false         //shorhand declare

	for i := 0; i < len(Data.things); i++ {
		if !p1_valid {
			if *Data.things[i] == *Data.player1_input {
				p1_valid = true
			} else {
				p1_valid = false
			}
		}

		if !p2_valid {
			if *Data.things[i] == *Data.player2_input {
				p2_valid = true
			} else {
				p2_valid = false
			}
		}

		if p1_valid && p2_valid {
			return true
		}
	}

	return false
}

func is_win(Data *data) {
	Data.round++

	if *Data.player1_input == *Data.things[0] {
		if *Data.player2_input == *Data.things[1] {
			fmt.Println("player 1 win")
			Data.player1_win++
		} else if *Data.player2_input == *Data.things[2] {
			fmt.Println("player 2 win")
			Data.playe2_win++
		}

		return

	} else if *Data.player1_input == *Data.things[1] {
		if *Data.player2_input == *Data.things[0] {
			fmt.Println("playe 2 win")
			Data.playe2_win++
		} else if *Data.player2_input == *Data.things[2] {
			fmt.Println("playe 1 win")
			Data.player1_win++
		}

		return

	} else if *Data.player1_input == *Data.things[2] {
		if *Data.player2_input == *Data.things[0] {
			fmt.Println("playe 1 win")
			Data.player1_win++
		} else if *Data.player2_input == *Data.things[1] {
			fmt.Println("playe 2 win")
			Data.player1_win++
		}

		return
	}
	fmt.Println("draw game")
}

func main() {
	var Data data
	container(&Data)

	fmt.Println("game is over")
	if Data.player1_win > Data.playe2_win {
		fmt.Println("player 1 is winner")
	} else if Data.player1_win < Data.playe2_win {
		fmt.Println("playe 2 is win")
	} else {
		fmt.Println("draw game")
	}
	fmt.Println("thank for playing")
}
