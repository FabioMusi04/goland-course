package main

import (
	"fmt"
	"guiApp/words"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("My Windows Custom UI")
	myWindow.Resize(fyne.NewSize(600, 200))

	timer := 15
	timeFinished := false
	timerLabel := widget.NewLabel("00:00")
	timerLabel.Alignment = fyne.TextAlignLeading
	timerLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}

	score := 0
	scoreLabel := widget.NewLabel(fmt.Sprintf("Score: %d", score))
	scoreLabel.Alignment = fyne.TextAlignTrailing
	scoreLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}

	wordSelected := words.GetRandomWord()
	wordSelectedSwapped := words.MixLetters(wordSelected)

	wordLabel := widget.NewLabel(wordSelectedSwapped)
	wordLabel.Alignment = fyne.TextAlignCenter
	wordLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}

	guessedWordsLabel := widget.NewLabel("Guessed Words: ")
	guessedWordsLabel.Alignment = fyne.TextAlignCenter
	guessedWordsLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}

	guessedWords := []string{}
	guessedWordsList := widget.NewList(
		func() int {
			return len(guessedWords)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(guessedWords[i])
		},
	)

	guessing := false
	inputLabel := widget.NewEntry()
	inputLabel.SetMinRowsVisible(2)
	inputLabel.PlaceHolder = "Enter the word here"
	inputLabel.OnSubmitted = func(s string) {
		if s == wordSelected && !timeFinished {
			guessing = true

			score++
			scoreLabel.SetText(fmt.Sprintf("Score: %d", score))

			guessedWords = append(guessedWords, s)
			guessedWordsList.Refresh()

			wordSelected = words.GetRandomWord()
			wordSelectedSwapped = words.MixLetters(wordSelected)

			wordLabel.SetText(wordSelectedSwapped)
			inputLabel.SetText("")

			timerLabel.SetText("00:00")
			timer = 15
		}
		guessing = false
	}

	scrollableGuessedWordsList := container.NewVScroll(guessedWordsList)
	scrollableGuessedWordsList.Resize(fyne.NewSize(200, 400))

	restartButton := widget.NewButton("Restart", func() {
		timeFinished = false
		timer = 15
		timerLabel.SetText(fmt.Sprintf("00:%02d", timer))

		score = 0
		scoreLabel.SetText(fmt.Sprintf("Score: %d", score))

		wordSelected = words.GetRandomWord()
		wordSelectedSwapped = words.MixLetters(wordSelected)
		wordLabel.SetText(wordSelectedSwapped)

		guessedWords = []string{}
		guessedWordsList.Refresh()

		guessing = false

		go startTimer(timerLabel, &timer, &timeFinished, scoreLabel, &score, &guessing)
	})

	myWindow.SetContent(
		container.New(
			layout.NewBorderLayout(
				timerLabel,
				scoreLabel,
				nil,
				nil,
			),
			timerLabel,
			scoreLabel,
			container.New(
				layout.NewGridLayout(1),
				wordLabel,
				guessedWordsLabel,
				scrollableGuessedWordsList,
				inputLabel,
				restartButton,
			),
		))

	go startTimer(timerLabel, &timer, &timeFinished, scoreLabel, &score, &guessing)

	myWindow.ShowAndRun()
}

func startTimer(timerLabel *widget.Label, timer *int, timeFinished *bool, scoreLabel *widget.Label, score *int, guessing *bool) {
	for !*timeFinished && !*guessing {
		time.Sleep(1 * time.Second)
		*timer--
		timerLabel.SetText(fmt.Sprintf("00:%02d", *timer))
		if *timer == 0 {
			*timeFinished = true
			scoreLabel.SetText(fmt.Sprintf("Game Over! Score: %d", *score))
		}
	}
}
