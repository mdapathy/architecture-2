//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mdapathy/architecture-2/server/forums"
	"github.com/mdapathy/architecture-2/server/users"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*ForumsApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from forums package.
		forums.Providers,
		// Add providers from users package.
		users.Providers,
		// Provide ForumsApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(ForumsApiServer), "Port", "ForumsHandler", "UsersHandler"),
	)
	return nil, nil
}
