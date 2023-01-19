package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input     io.Reader
	output    io.Writer
	wordCount bool
}

type option func(*counter) error

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return counter{}, err
		}
	}
	return c, nil
}

func (c *counter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.input)

	for scanner.Scan() {
		lines++
	}

	return lines
}

func Lines() int {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return c.Lines()
}

func Words() int {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return c.Words()
}

func (c *counter) Words() int {
	words := 0
	scanner := bufio.NewScanner(c.input)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words++
	}

	return words
}

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil input reader")
		}
		c.output = output
		return nil
	}
}

func FromArgs(args []string) option {
	return func(c *counter) error {
		fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		wordCount := fset.Bool("w", false, "Count words instead of lines")

		fset.SetOutput(c.output)

		err := fset.Parse(args)
		if err != nil {
			return err
		}

		c.wordCount = *wordCount

		args = fset.Args()
		if len(args) < 1 {
			return nil
		}

		f, err := os.Open(args[0])
		if err != nil {
			return err
		}

		c.input = f

		return nil
	}
}

func RunCLI() {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)

	fmt.Println(c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if c.wordCount {
		fmt.Println(c.Words())
	} else {
		fmt.Println(c.Lines())
	}
}
