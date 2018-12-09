package main

import (
	"os"
)

func getInput() {

}

func firstChallenge() {

}

func secondChallenge() {

}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}
