package day01

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/elprogramadorgt/aoc_go/net"
)

const url = "https://adventofcode.com/2022/day/1/input"

func SolutionDay01Exercise01() {

	body, err := net.GetRemoteFile(url)

	if err != nil {
		fmt.Println("Error on connection")
	}
	bodyStr := strings.TrimSpace(string(body))
	elfCaloriesArr, err := groupCalories(bodyStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(slices.Max(elfCaloriesArr))

}

func SolutionDay01Exercise02() {

	body, err := net.GetRemoteFile(url)

	if err != nil {
		fmt.Println("Error on connection")
	}
	bodyStr := strings.TrimSpace(string(body))
	elfCaloriesArr, err := groupCalories(bodyStr)
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(elfCaloriesArr)

	topThree := elfCaloriesArr[len(elfCaloriesArr)-3:]

	fmt.Println(sum(topThree))

}

func groupCalories(dataStream string) ([]int, error) {

	elfs := strings.Split(dataStream, "\n\n")

	var elfCaloriesArr []int

	for _, elf := range elfs {
		total_calories := 0
		for _, caloriesStr := range strings.Split(elf, "\n") {
			calories, err := strconv.Atoi(caloriesStr)

			if err != nil {
				return nil, err
			}

			total_calories += calories
		}
		elfCaloriesArr = append(elfCaloriesArr, total_calories)
	}

	return elfCaloriesArr, nil

}

func sum(elements []int) (int, error) {
	total := 0
	for _, element := range elements {
		total += element
	}

	return total, nil
}
