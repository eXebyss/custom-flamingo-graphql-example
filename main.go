package main

//go:generate rm -f graphql/generated.go
//go:generate go run -tags graphql main.go graphql

import (
	projectgraphql "FlamingoGQLExample/graphql"
	"FlamingoGQLExample/todo"
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
)

func main() {
	flamingo.App([]dingo.Module{
		new(todo.Module),
		new(projectgraphql.Module),
		new(requestlogger.Module),
	})
}
