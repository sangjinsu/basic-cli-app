package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]
	A greeter application which prints the name you entered <integer> number of times.
`, os.Args[0])

func main() {
	reader := bufio.NewReader(os.Stdin)

	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		printUsage()
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Println(err)
		printUsage()
		os.Exit(1)
	}
	err = runCmd(reader, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type config struct {
	numTimes   int
	printUsage bool
}

func printUsage() {
	fmt.Println(usageString)
}

func getName(r io.Reader) (string, error) {
	message := "이름을 작성하세요. 작성한 뒤 엔터를 누르세요."
	fmt.Println(message)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("이름을 작성하지 않았습니다")
	}
	return name, nil
}

func parseArgs(args []string) (config, error) {
	numTimes := 0
	var err error
	c := config{}
	if len(args) < 1 {
		return c, errors.New("invalid number of arguments")
	}

	if args[0] == "-h" || args[0] == "-help" {
		c.printUsage = true
		return c, nil
	}

	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes

	return c, nil
}

func validateArgs(c config) error {
	if c.numTimes <= 0 {
		return errors.New("1개 이상의 숫자가 필요합니다")
	}
	return nil
}

func runCmd(r io.Reader, c config) error {
	if c.printUsage {
		printUsage()
		return nil
	}
	name, err := getName(r)
	if err != nil {
		return err
	}

	greetUser(c, name)
	return nil
}

func greetUser(c config, name string) {
	msg := fmt.Sprintf("Nice to meet you %s", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Println(msg)
	}
}
