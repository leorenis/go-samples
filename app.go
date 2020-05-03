package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const mentoringTimes = 5
const mentoringDelay = 5

func main() {

	showIntroduction()

	for {
		showMenu()

		command := readCommand()
		switch command {
		case 1:
			startMentoring()
		case 2:
			fmt.Println("Logs")
		case 0:
			os.Exit(0)
		default:
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	personName := "John Doe"
	appVersion := 1.1
	age := 28

	fmt.Println("The first Go lang's app written by mr.", personName, "age", age)
	fmt.Println("App Version: ", appVersion)
}

func showMenu() {
	fmt.Println("1- Start mentoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The command choosen:", command)
	fmt.Println("")
	return command
}

func startMentoring() {
	fmt.Println("Mentoring")

	/* portals := []string{
		"https://google.com",
		"https://facebook.com",
		"https://twitter.com",
	    "http://random-status-code.herokuapp.com"} */

	portals := readPortalsFile()

	for i := 0; i < mentoringTimes; i++ {
		for index, portal := range portals {
			fmt.Println("Passing for slice position", index, "was found portal", portal)
			doMentoringTest(portal)
		}
		time.Sleep(mentoringDelay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func doMentoringTest(portal string) {
	response, err := http.Get(portal)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("The site", portal, "its ok!")
	} else {
		fmt.Println("The site", portal, "its downtime. Status code: ", response.StatusCode)
	}
}

func readPortalsFile() []string {
	file, err := os.Open("resources/sites.txt")
	if err != nil {
		fmt.Println("Error: was not possible to read file.", err)
	}

	var portals []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		portals = append(portals, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return portals
}
