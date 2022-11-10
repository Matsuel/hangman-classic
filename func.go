package funct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func ChooseWord() string {
	name := os.Args[1]
	body, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	list := []string{}
	hold := ""
	for _, m := range string(body) {
		if m != 10 {
			hold = hold + string(m)
		} else {
			if hold != "" {
				list = append(list, hold)
				hold = ""
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	list = append(list, hold)
	lent := rand.Intn(len(list))
	return list[lent]
}
func InitGame(word, mod_game string) ([]string, int) {
	os := runtime.GOOS
	taille := 0
	if os == "windows" {
		taille = 1
	}
	mot := []string{}
	for i := 0; i < len(word)-taille; i++ {
		mot = append(mot, "_")
	}
	if mod_game == "hard" {
		var letterreveal int
		for i := 0; i < (len(word)/3)-1; i++ {
			letterreveal = rand.Intn(len(mot))
			mot[letterreveal] = string(word[letterreveal])
		}
	} else {
		var letterreveal int
		for i := 0; i < (len(word)/2)-1; i++ {
			letterreveal = rand.Intn(len(mot))
			mot[letterreveal] = string(word[letterreveal])
		}
	}
	return mot, 10
}

func Play(attempts int, word string, word_array []string, hang_pos []string, count int, mod_game string) {
	OScount := 0
	OS := runtime.GOOS
	if OS == "windows" {
		OScount = 8
	} else if OS == "linux" || OS == "darwin" {
		OScount = 7
	}
	//count := 0
	var present bool
	var letter string
	letters_list_used := []string{}
	cpt := 0
	for word != TabtoStr(word_array) {
		if attempts <= 0 {
			fmt.Println()
			for i := len(hang_pos) - 8; i < len(hang_pos)-1; i++ {
				fmt.Println(hang_pos[i])
			}
			PrintWinLoose(false, word)
			return
		} else if mod_game == "normal" {
			present = false
			fmt.Print("Choose: ")
			fmt.Scan(&letter)
			if Accent(letter) {
				letter = AccentToLetters(letter)
			} else if letter == "STOP" {
				Save(attempts, count, word, word_array, letters_list_used, mod_game)
				return
			}
			if IsUse(letter, letters_list_used) {
				present = true
			}
			if len(letter) > 1 {
				if letter == word {
					Bim()
					fmt.Println("Congrats !")
					return
				} else {
					attempts--
					count += OScount
				}
			}
			for i := 0; i < len(word); i++ {
				if string(word[i]) == letter {
					word_array[i] = letter
					present = true
				}
			}
		} else if mod_game != "normal" {
			if cpt > 3 {
				attempts--
				fmt.Println("You have already choose 3 vowels ", attempts, " attempts remaining")
			}
			present = false
			fmt.Print("Choose: ")
			fmt.Scan(&letter)
			if Accent(letter) {
				letter = AccentToLetters(letter)
			} else if letter == "STOP" {
				Save(attempts, count, word, word_array, letters_list_used, mod_game)
				return
			}
			if IsUse(letter, letters_list_used) {
				present = true
			}
			if len(letter) > 1 {
				if letter == word {
					Bim()
					fmt.Println("Congrats !")
					return
				} else {
					attempts--
					count += OScount
				}
			}
			if IsVoyelle(letter) {
				cpt += 1
			}
			for i := 0; i < len(word); i++ {
				if string(word[i]) == letter {
					word_array[i] = letter
					present = true
				}
			}
		}
		if !present {
			attempts--
			if attempts >= 1 {
				fmt.Println("Not present in the word, ", attempts, " attempts remaining")
				fmt.Println()
				//letters_list_used = append(letters_list_used, letter)
				for num := count; num < count+OScount; num++ {
					fmt.Println(hang_pos[num])
				}
			}
			count += OScount
			//count += 8
		}
		if present && IsUse(letter, letters_list_used) {
			if mod_game == "hard" {
				attempts--
				fmt.Println("Letters already used ", attempts, " attempts remaining")
				for num := count; num < count+OScount; num++ {
					fmt.Println(hang_pos[num])
				}
				count += OScount
			} else {
				//attempts--
				PrintLetterUse(letters_list_used)
				fmt.Println()
			}
		}
		fmt.Println(TabtoStr(word_array))
		if !IsUse(letter, letters_list_used) {
			letters_list_used = append(letters_list_used, letter)
		}
	}
	PrintWinLoose(true, word)
}

func ShowWord(word []string) string {
	var motstr string
	for _, ch := range word {
		motstr += "" + ch
	}
	return motstr
}
func TabtoStr(word []string) string {
	str := ""
	for _, ch := range word {
		str += ch
	}
	return str
}
func PosHangman() []string {
	hang_pos := []string{}
	bod, err := ioutil.ReadFile("ascii/hangman.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	hold2 := ""
	for _, d := range string(bod) {
		if d != 10 {
			hold2 = hold2 + string(d)
		} else {
			if hold2 != "" {
				hang_pos = append(hang_pos, hold2)
				hold2 = ""
			}
		}
	}
	hang_pos = append(hang_pos, hold2)
	return hang_pos
}

func IsUse(letter string, letter_list []string) bool {
	for i := 0; i < len(letter_list); i++ {
		if letter == letter_list[i] {
			return true
		}
	}
	return false
}

func PrintLetterUse(letter_use []string) {
	word := "Letters already used : "
	if len(letter_use) == 0 {
		word += "None"
		fmt.Println(word)
		return
	} else {
		for i := 0; i < len(letter_use)-1; i++ {
			word += letter_use[i] + " "
		}
		word += letter_use[len(letter_use)-1]
		fmt.Println(word)
		return
	}
}

func PrintWinLoose(b bool, tofind string) {
	if b == true {
		fmt.Println("Congrats !")
		Bim()
		return
	} else {
		word := "You loose ! The word you have to find was : "
		word += tofind
		OhSnap()
		fmt.Println(word)
		return
	}
}

func AccentToLetters(letter string) string {
	rep := ""
	letter_rune := []rune(letter)
	if letter_rune[0] >= rune(232) && letter_rune[0] <= rune(235) {
		rep = "e"
	} else if letter_rune[0] >= rune(224) && letter_rune[0] <= rune(230) {
		rep = "a"
	} else if letter_rune[0] >= rune(236) && letter_rune[0] <= rune(240) {
		return "i"
	} else if letter_rune[0] == rune(240) || letter_rune[0] >= rune(242) && letter_rune[0] <= rune(248) {
		return "o"
	} else if letter_rune[0] >= rune(249) && letter_rune[0] <= rune(252) {
		return "u"
	}
	return rep
}

func Accent(letter string) bool {
	return []rune(letter)[0] > 128
}

type data struct {
	Attempts   int
	word       string
	word_array string
}

func Itoa(i int) string {
	itoa := ""
	for i != 0 {
		ch := i % 10
		i /= 10
		itoa += string(ch + '0')
	}
	res := ""
	for i := len(itoa) - 1; i >= 0; i-- {
		res += string(itoa[i])
	}
	return res
}

func Save(a int, count int, w string, m []string, letters_list_used []string, mod_game string) {
	filename := "save.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err == nil {
			if err == nil {
				Game := GameData{w, a, count, m, letters_list_used, mod_game}
				data, err := json.Marshal(Game)
				if err == nil {
					ioutil.WriteFile("save.txt", data, 0644)
				} else {
					fmt.Println(err)
				}
			}
		}
	} else {
		if err == nil {
			Game := GameData{w, a, count, m, letters_list_used, mod_game}
			data, err := json.Marshal(Game)
			if err == nil {
				ioutil.WriteFile("save.txt", data, 0644)
			} else {
				fmt.Println(err)
			}
		}
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

func OhSnap() {
	content, err := ioutil.ReadFile("ascii/ohsnap.txt")

	if err == nil {
		fmt.Printf(string(content))
	}
	fmt.Println()
}

func Bim() {
	content, err := ioutil.ReadFile("ascii/bim.txt")

	if err == nil {
		fmt.Printf(string(content))
	}
	fmt.Println()
}

func IsVoyelle(letter string) bool {
	if letter == "a" || letter == "e" || letter == "i" || letter == "o" || letter == "u" || letter == "y" {
		return true
	}
	return false
}

func Welcome() {
	content, err := ioutil.ReadFile("ascii/welcome.txt")

	if err == nil {
		fmt.Printf(string(content))
	}
	fmt.Println()
}
