package main

import (
	"encoding/json"
	"fmt"
	"funct"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Expect one argument")
		return
	}
	taille := 0
	OS := runtime.GOOS
	if OS == "windows" {
		taille = 1
	}
	Game := GameData{}
	hang_pos := funct.PosHangman()
	if len(os.Args) > 2 {
		if os.Args[2] == "--StartWith" {
			filename := os.Args[3]
			data, _ := ioutil.ReadFile(filename)
			json.Unmarshal(data, &Game)
			fmt.Println("Welcome back, you have", Game.Attempts, "attemps remaining!")
			fmt.Println(funct.ShowWord(Game.Word))
			os.Remove("save.txt")
			funct.Play(Game.Attempts, Game.Solution, Game.Word, hang_pos, Game.Count_line, Game.Gamemod_game)
		} else if os.Args[2] == "--hard" {
			Game.Solution = funct.ChooseWord()
			Game.Gamemod_game = "hard"
			Game.Word, Game.Attempts = funct.InitGame(Game.Solution, "hard")
			var nouvmot string
			for i := 0; i < len(Game.Solution)-taille; i++ {
				nouvmot += string(Game.Solution[i])
			}
			fmt.Println("Good Luck, you have", Game.Attempts, " attempts.")
			mott := funct.ShowWord(Game.Word)
			fmt.Println(mott)
			funct.Play(Game.Attempts, nouvmot, Game.Word, hang_pos, 0, Game.Gamemod_game)
		}
	} else if len(os.Args) <= 2 {
		funct.Welcome()
		Game.Solution = funct.ChooseWord()
		Game.Gamemod_game = "normal"
		var nouvmot string
		for i := 0; i < len(Game.Solution)-taille; i++ {
			nouvmot += string(Game.Solution[i])
		}
		Game.Word, Game.Attempts = funct.InitGame(Game.Solution, "normal")
		fmt.Println("Good Luck, you have", Game.Attempts, " attempts.")
		mott := funct.ShowWord(Game.Word)
		fmt.Println(mott)
		funct.Play(Game.Attempts, nouvmot, Game.Word, hang_pos, 0, Game.Gamemod_game)
	}
}

type GameData struct {
	Solution     string
	Attempts     int
	Count_line   int
	Word         []string
	Letters_used []string
	Gamemod_game string
}
