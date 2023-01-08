package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/aquilax/escape"
)

type ListCommand struct {
	w  io.Writer
	fs *flag.FlagSet
}

func NewListCommand(w io.Writer) *ListCommand {
	return &ListCommand{
		w:  w,
		fs: flag.NewFlagSet("list", flag.ContinueOnError),
	}
}

func (c *ListCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *ListCommand) Run() error {
	fmt.Fprintln(c.w, "List of supported formatters:")
	for _, e := range escape.AllEscapers() {
		if _, err := fmt.Fprintf(c.w, "  * %s\n", e.Name()); err != nil {
			return err
		}
	}
	return nil
}

func (c *ListCommand) Name() string {
	return c.fs.Name()
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	output := os.Stdout

	commands := []Runner{
		NewListCommand(output),
	}

	subcommand := os.Args[1]

	for _, cmd := range commands {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
