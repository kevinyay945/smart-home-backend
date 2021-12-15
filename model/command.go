package model

import (
	"fmt"
	"gorm.io/gorm"
	"smart-home-backend/lib/pq/schema"
)

type ICommand interface {
	Get() ([]schema.Command, error)
	Save(input *schema.Command) (schema.Command, error)
	UpdateOne(commandUuid string, command *schema.Command) (schema.Command, error)
	Delete(commandUuid string) (schema.Command, error)
}

type MCommand struct{
	db *gorm.DB
}

func NewOriginCommand(_db *gorm.DB) ICommand {
	output := new(MCommand)
	output.db = _db
	return output
}

func NewCommand() ICommand {
	return NewOriginCommand(db)
}

func (c *MCommand) Get() (Commands []schema.Command, err error) {
	err = c.db.Model(schema.Command{}).Find(&Commands).Error
	return
}

func (c *MCommand) Save(input *schema.Command) (schema.Command, error) {

	return *input, nil
}

func (c *MCommand) UpdateOne(commandUuid string, command *schema.Command) (schema.Command, error) {
	output := new(schema.Command)
	for _, command := range schema.AllCommand {
		if commandUuid == command.Uuid {
			output = &command
		}
	}
	return *output, nil
}

func (c *MCommand) Delete(commandUuid string) (schema.Command, error) {
	fmt.Println("Delete uuid", commandUuid)
	return schema.Command{}, nil
}

