package config

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Commander is a contract for config commands
type Commander interface {
	// Command builds the config command
	Command() (string, error)
}

// SetPropCommand is a command to set common properties.
// See: https://lucene.apache.org/solr/guide/8_5/config-api.html#commands-for-common-properties
type SetPropCommand struct {
	prop string
	val  interface{}
}

// NewSetPropCommand is a factory for SetPropCommand
func NewSetPropCommand(prop string, val interface{}) Commander {
	return SetPropCommand{prop: prop, val: val}
}

func (c SetPropCommand) Command() (string, error) {
	m := map[string]interface{}{c.prop: c.val}
	b, err := json.Marshal(m)
	if err != nil {
		return "", errors.Wrap(err, "marshal command")
	}

	return `"set-property": ` + string(b), nil
}

// UnsetPropCommand is a command to unset common properties.
// See: https://lucene.apache.org/solr/guide/8_5/config-api.html#commands-for-common-properties
type UnsetPropCommand struct {
	prop string
}

// NewUnsetPropCommand is a factory for UnsetPropCommand
func NewUnsetPropCommand(prop string) Commander {
	return UnsetPropCommand{prop: prop}
}

func (c UnsetPropCommand) Command() (string, error) {
	return fmt.Sprintf(`"unset-property": %q`, c.prop), nil
}

// CommandType is a component command type
type CommandType string

// Basic commands for components
const (
	AddRequestHandler         CommandType = "add-requesthandler"
	UpdateRequestHandler      CommandType = "update-requesthandler"
	DeleteRequestHandler      CommandType = "delete-requesthandler"
	AddSearchComponent        CommandType = "add-searchcomponent"
	UpdateSearchComponent     CommandType = "update-searchcomponent"
	DeleteSearchComponent     CommandType = "delete-searchcomponent"
	AddInitParams             CommandType = "add-initparams"
	UpdateInitParams          CommandType = "update-initparams"
	DeleteInitParams          CommandType = "delete-initparams"
	AddQueryResponseWriter    CommandType = "add-queryresponsewriter"
	UpdateQueryResponseWriter CommandType = "update-queryresponsewriter"
	DeleteQueryResponseWriter CommandType = "delete-queryresponsewriter"
)

// ComponentCommand is a component command.
// See: https://lucene.apache.org/solr/guide/8_5/config-api.html#commands-for-handlers-and-components
type ComponentCommand struct {
	CommandType CommandType
	Body        map[string]interface{}
}

// NewComponentCommand is a factory for component command
func NewComponentCommand(commandType CommandType, body map[string]interface{}) Commander {
	return &ComponentCommand{
		CommandType: commandType,
		Body:        body,
	}
}

func (c *ComponentCommand) Command() (string, error) {
	b, err := json.Marshal(c.Body)
	if err != nil {
		return "", errors.Wrap(err, "marshal command body")
	}

	return fmt.Sprintf(`"%s": `+string(b), c.CommandType), nil
}