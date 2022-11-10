<!-- PROJECT LOGO -->
<p align="center">
  <img src="./img/hangman.png" />
</p>
<br />
<div align="center">
    

  <h1 align="center">HangMan project in golang</h1>

  <p align="center">
    <br />
  </p>
</div>



<!-- ABOUT THE PROJECT -->
### Built With

![Golang](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/768px-Go_Logo_Blue.svg.png?20191207190041)

Two gamemodes, hard or normal

</br>

<!-- GETTING STARTED -->
## Getting Started

Here we have created a golang version of the famous hangman game. You have 10 attempts to find a word, some letters will be revealed at the start, the objective is to find them all in order to win the game.

### Features available

<ul>
  <li>Save the game in a file and restart after</li>
  <li>Hard mode or normal mode (with more difficulties).</li>
  <li>Ascii Art (welcome, bim, or oh snap).</li>
</ul>


### Prerequisites

You need to have the go language installed beforehand

[Golang](https://go.dev/dl/)

### Installation

_To use the project you should clone the repo._

Clone the repo
   ```sh
   git clone https://ytrack.learn.ynov.com/git/lmatheo/hangman-classic.git
   ```


<!-- RUN THE PROJECT -->
### Run the project
```
cd hangman-classic
```
</br>

```
go run /main/main.go wordlist/words.txt
```
```words.txt , words2.txt or words3.txt```
##### If you want to start the game in hard mode:

```
go run /main/main.go wordlist/words.txt --hard
```

##### And if you want to restart from a game save:

```
go run ./main/main.go ./wordlist/words.txt --StartWith save.txt
```
### Commands in game 


##### Submit one letter or one word
##### STOP the game, save in a file
