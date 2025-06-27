package commands

import (
	"github.com/felixsolom/gator/internal/config"
	"github.com/felixsolom/gator/internal/database"
)

type State struct {
	Db              *database.Queries
	PointerToConfig *config.Config
}
