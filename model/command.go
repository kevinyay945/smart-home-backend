package model

import (
	"gorm.io/gorm"
	"smart-home-backend/lib/pg/schema"
)

type ICommand interface {
	Get() ([]schema.Command, error)
	Save(input *schema.Command) (schema.Command, error)
	UpdateOne(commandUuid string, command *schema.Command) (schema.Command, error)
	Delete(commandUuid string) error
}

type MCommand struct {
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
	err := c.db.Save(input).Error
	return *input, err
}

func (c *MCommand) UpdateOne(commandUuid string, command *schema.Command) (schema.Command, error) {
	err := c.db.Model(&schema.Command{}).
		Updates(command).
		Where("uuid = ?", commandUuid).Error
	return *command, err
}

func (c *MCommand) Delete(commandUuid string) error {
	err := c.db.Model(&schema.Command{}).Delete("uuid = ?", commandUuid).Error
	return err
}
