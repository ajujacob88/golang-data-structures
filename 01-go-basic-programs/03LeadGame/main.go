package main

import "fmt"

func main() {
	var noOfRounds, finalWinner, maxlead int = 0, 0, 0
	fmt.Println("Enter no of rounds: ")
	fmt.Scan(&noOfRounds)
	var playerAScoreSum, playerBScoreSum int = 0, 0

	for i := 1; i <= noOfRounds; i++ {
		var score int

		fmt.Print("Enter player 1 score - Round ", i, " : ")
		fmt.Scan(&score)
		playerAScoreSum += score
		fmt.Print("Enter player 2 score - Round ", i, " : ")
		fmt.Scan(&score)
		playerBScoreSum += score
		var lead, winner int

		if playerAScoreSum > playerBScoreSum {
			lead = playerAScoreSum - playerBScoreSum
			winner = 1
		} else if playerBScoreSum > playerAScoreSum {
			lead = playerBScoreSum - playerAScoreSum
			winner = 2
		}
		if lead > maxlead {
			maxlead = lead
			finalWinner = winner
		}

	}

	fmt.Println("The winner is player ", finalWinner, " and maximum lead is ", maxlead)
}
