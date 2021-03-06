package command

import (
	"fmt"

	"github.com/yitsushi/go-commander"
	s "github.com/yitsushi/totp-cli/storage"
	"github.com/yitsushi/totp-cli/util"
)

// List structure is the representation of the list command
type List struct {
}

// Execute is the main function. It will be called on list command
func (c *List) Execute(opts *commander.CommandHelper) {
	storage := s.PrepareStorage()
	ns := opts.Arg(0)
	if len(ns) < 1 {
		for _, namespace := range storage.Namespaces {
			fmt.Printf("%s (Number of accounts: %d)\n", namespace.Name, len(namespace.Accounts))
		}

		return
	}

	namespace, err := storage.FindNamespace(ns)
	util.Check(err)

	for _, account := range namespace.Accounts {
		fmt.Printf("%s.%s\n", namespace.Name, account.Name)
	}
}

// NewList creates a new List command
func NewList(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &List{},
		Help: &commander.CommandDescriptor{
			Name:             "list",
			ShortDescription: "List all available namespaces or accounts under a namespace",
			Arguments:        "[namespace]",
			Examples: []string{
				"",
				"mynamespace",
			},
		},
	}
}
