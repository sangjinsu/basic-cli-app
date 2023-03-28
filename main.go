package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	c, err := parseArgs(writer, os.Args[1:])
	if err != nil {
		if err != flag.ErrHelp {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	if err = validateArgs(c); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = runCmd(reader, c); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type config struct {
	numTimes int
	name     string
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

func parseArgs(writer *bufio.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(writer)
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	fs.Usage = func() {
		usageString := `
A greeter application which prints the name you entered a specified number of times.

Usage of %s: <option> [name]`
		fmt.Fprintf(os.Stderr, usageString+"\n", fs.Name())
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "-n int\n")
		fmt.Fprintf(os.Stderr, "\tNumber of times to greet\n")
		fs.PrintDefaults()
	}
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() > 1 {
		return c, errors.New("잘못된 위치 인수입니다")
	}
	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}
	return c, nil
}

func validateArgs(c config) error {
	if c.numTimes <= 0 {
		return errors.New("1 이상의 숫자가 필요합니다")
	}
	return nil
}

func runCmd(r io.Reader, c config) error {
	if len(c.name) == 0 {
		var err error
		c.name, err = getName(r)
		if err != nil {
			return fmt.Errorf("이름을 가져오는 중에 오류가 발생했습니다: %w", err)
		}
	}
	greetUser(c)
	return nil
}

func greetUser(c config) {
	msg := fmt.Sprintf("만나서 반갑습니다 %s", c.name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Println(msg)
	}
}
