package day02

import (
	"fmt"
	"strings"
	"github.com/elprogramadorgt/aoc_go/net"
)

const url = "https://adventofcode.com/2022/day/2/input"

var points = map[string]int{"A":1, "B":2, "C":3, "X":1, "Y":2, "Z":3}

var results = map[string]int{"A X": 3, "A Y":6, "B Y":3, "B Z": 6, "C X":6,  "C Z": 3 } 


var points2 = map[string]int{"A X": 3, "A Y": 1, "A Z": 2, "B X": 1, "B Y": 2, "B Z": 3, "C X": 2, "C Y":3, "C Z": 1}

var results2  = map[string]int{"Y":3, "Z": 6}


func initConfig() []string {
 	body, err := net.GetRemoteFile(url)

	if err != nil {
		fmt.Println("Error on connection")
	}

	bodyStr := strings.TrimSpace(string(body))
	
	dataArr := strings.Split(bodyStr, "\n")

	return dataArr


}

func SolutionDay02Exercise01(){


	dataArr := initConfig();

	totalScore := 0

	for _, v := range dataArr {

		totalScore += results[v] + points[string(v[2])]

 	}

	fmt.Println("The total is:", totalScore)
}

func SolutionDay02Exercise02(){

	dataArr := initConfig();
	
	totalScore := 0

	for _, v := range dataArr {
	
		totalScore += points2[v] +  results2[string(v[2])]
 	}

	fmt.Println("The total is:", totalScore)


}
