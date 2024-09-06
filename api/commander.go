package api

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mrdhat/eth-txns/errors"
)

type Commander interface {
	Start(stop chan bool) error
}

type commander struct {
	parser Parser
}

func (c *commander) parseCommand(text string) (string, []string, error) {
	parts := strings.Split(text, " ")
	if len(parts) < 1 {
		return "", nil, errors.ErrInvalidCommand
	}
	return parts[0], parts[1:], nil
}

func (c *commander) handleError(err error) {
	// we don't need to kill the program here
	if err != nil {
		fmt.Println(err)
	}
}

func (c *commander) handleCommand(command string, args []string) {
	switch command {
	case "subscribe":
		if len(args) < 1 {
			c.handleError(errors.ErrInvalidCommand)
		}
		err := c.parser.Subscribe(args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Subscribed to address: ", args[0])
	case "get-transactions":
		if len(args) < 1 {
			c.handleError(errors.ErrInvalidCommand)
		}
		transactions := c.parser.GetTransactions(args[0])
		if len(transactions) == 0 || transactions == nil {
			fmt.Println("No transactions found for address ", args[0])
		} else {
			fmt.Println("Transactions for address ", args[0], ": ", transactions)
		}

	case "get-current-block":
		currentBlock := c.parser.GetCurrentBlock()
		fmt.Println("Current block: ", currentBlock)
	default:
		c.handleError(errors.ErrInvalidCommand)
	}
}

func (c *commander) Start(stop chan bool) error {
	for {
		select {
		case <-stop:
			return nil
		default:
			// read input from stdin
			reader := bufio.NewReader(os.Stdin)
			fmt.Println()
			fmt.Println("Available Commands:")
			fmt.Println("1. subscribe <address>")
			fmt.Println("2. get-transactions <address>")
			fmt.Println("3. get-current-block")
			fmt.Print("Enter Command: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			// remove the newline character
			text = strings.TrimSpace(text)
			// parse the input
			command, args, err := c.parseCommand(text)
			if err != nil {
				return err
			}
			c.handleCommand(command, args)
		}
	}

}

func NewCommander(parser Parser) Commander {
	return &commander{
		parser: parser,
	}
}
