package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	printSegitiga(5)
	fmt.Println("Generate Password :", genPass("abcd", "strong"))
	lamaPenerbangan(7)
}

// TASK I
func printSegitiga(n int) {
	for x := 1; x <= n; x++ {
		for y := 0; y <= x; y++ {
			fmt.Print(" ")
		}
		for z := n; z >= (2*x)-n; z-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// TASK II
func genPass(password string, level string) string {
	var letterLow string = "0123456789QWERTYUIOPASDFGHJKLZXCVBNM"
	var letterMedium string = "qwertyuiopasdfghjklzxcvbnm" + letterLow
	var letterStrong string = "(~!#$%^&*()_-+?<>|[]{}:" + letterMedium

	var passwordLength int = len(password) + 6

	rand.Seed(time.Now().UTC().UnixNano())

	if level == "low" {
		for i := len(password); i < passwordLength; i++ {
			randLetterLow := rand.Intn(len(letterLow))
			password = password + string(letterLow[randLetterLow])
		}
		return (password)
		// fmt.Println("Password Generate :", password)
	} else if level == "medium" {
		for i := len(password); i < passwordLength; i++ {
			randLetterMedium := rand.Intn(len(letterMedium))
			password = password + string(letterMedium[randLetterMedium])
		}
		// fmt.Println("Password Generate :", password)
		return (password)
	} else if level == "strong" {
		for i := len(password); i < passwordLength; i++ {
			randLetterStrong := rand.Intn(len(letterStrong))
			password = password + string(letterStrong[randLetterStrong])
		}
		// fmt.Println("Password Generate :", password)
		return (password)
	} else {
		return ("Wrong Input")
	}
}

// TASK III
func lamaPenerbangan(n int) {
	var dataMovie = []int{1, 7, 3, 4, 8, 9}

	for i, item1 := range dataMovie {
		for j, item2 := range dataMovie {
			if i == j {
				continue
			}
			if i != j {
				if dataMovie[i]+dataMovie[j] == n {
					fmt.Println("Durasi film :", item1, "dan", item2)
					return
				}
			}
		}
	}
}
