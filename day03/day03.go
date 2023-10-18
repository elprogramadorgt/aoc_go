package day03

import (
	"fmt"
	"strings"
	 "slices"
	"github.com/elprogramadorgt/aoc_go/net"
)

const url = "https://adventofcode.com/2022/day/3/input"

var lettersMap = make(map[string]int)

func initConfig() []string {
	for i, letter := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		lettersMap[string(letter)] = i+1
	}
 	body, err := net.GetRemoteFile(url)

	if err != nil {
		fmt.Println("Error on connection")
	}

	bodyStr := strings.TrimSpace(string(body))
	
	dataArr := strings.Split(bodyStr, "\n")

	return dataArr


}

//letters := map[string]int{"a":1,"b":2,"c":3,"d":4,"e":5,"f":6,"g":7,"h":8,"i":9,"j":10,"k":11,"l":12,"m":13,
//"n":14,"o":15,"p":16,"q":17,"r":18,"s":19,"t":20,"u":21,"v":22,"w":23,"x":24,"y":25,"z":26}



func SolutionDay03Exercise01(){
	dataArr := initConfig()

	var repeated []string
	for _, v := range dataArr {
		half := len(v) / 2
		first := v[:half]
		second := v[half:]

		var repeatedLetter []string
		for _, item := range first {
			letter := string(item)
			if slices.Contains(strings.Split(second, ""), letter) {
				if !slices.Contains(repeatedLetter, letter) {
					repeated = append(repeated, letter)	
					repeatedLetter = append(repeatedLetter, letter)
				}
			}
		}
	}
	total := 0
	for _, v := range repeated {
		total += lettersMap[v]
	}
	fmt.Println(total)
}


func SolutionDay03Exercise02() {

	dataArr := initConfig()
	groupArr := []string{}
	var repeatedLetter []string
	for i, v := range dataArr {
		
		groupArr = append(groupArr, v)
		
        if(len(groupArr) == 3){

            fmt.Println(i, "i on 3")

			for _, item := range groupArr[0]{
			    letter := string(item)
                existsInG2 := strings.Contains(groupArr[1], letter)
                existsInG3 := strings.Contains(groupArr[2], letter)
                
                if existsInG2 && existsInG3 {
                    repeatedLetter = append(repeatedLetter, letter)
                    
                    break
                }
			}
			groupArr = nil
		}

	}

    fmt.Println(repeatedLetter)
    total := 0
    for _, letter := range repeatedLetter {
    
        total += lettersMap[letter]
    }

    fmt.Println(total)
//	for i:=3; i < len(dataArr); i+=3 {
//		for j := i; j >= i-3; j-- {
//			fmt.Println(dataArr[j])
//		}
//
//		fmt.Println("========",i)
//	}
		



}


















