package model

import (
	"fmt"
	"time"
)

type ICommand interface {
	Get() []Command
	Save(input *Command) Command
	UpdateOne(commandUuid string, command *Command) Command
	Delete(commandUuid string) Command
}

type MCommand struct{}

type Command struct {
	Uuid string `json:"uuid"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Url string `json:"url"`
}

var AllCommand []Command = []Command{{
	"bbbad37f-c6cd-47f7-907d-add1c4045559",
	time.Now(),
	time.Now(),
	"http://example.com",
},{
	"bbbad37f-c6cd-47f7-907d-add1c4045558",
	time.Now(),
	time.Now(),
	"http://example2.com",
}}

func NewCommand() *MCommand {
	output := new(MCommand)
	return output
}

func (c *MCommand) Get() []Command {
	return AllCommand
}

func (c *MCommand) Save(input *Command) Command {
	return *input
}

func (c *MCommand) UpdateOne(commandUuid string, command *Command) Command {
	output := new(Command)
	for _, command := range AllCommand {
		if commandUuid == command.Uuid {
			output = &command
		}
	}
	return *output
}

func (c *MCommand) Delete(commandUuid string) Command {
	fmt.Println("Delete uuid", commandUuid)
	return Command{}
}

