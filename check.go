package main

import (
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func main() {
	ctx := cuecontext.New()

	// load schema.cue
	schemaBytes, err := os.ReadFile("schema.cue")
	if err != nil {
		panic(err)
	}
	schema := ctx.CompileBytes(schemaBytes, cue.Filename("schema.cue"))
	if err := schema.Err(); err != nil {
		panic(err)
	}

	// load data.json
	dataBytes, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	data := ctx.CompileBytes(dataBytes, cue.Filename("data.json"))
	if err := data.Err(); err != nil {
		panic(err)
	}

	// use #Def_1 from the schema
	schema = schema.LookupPath(cue.ParsePath("#Def_1"))
	if err := schema.Err(); err != nil {
		panic(err)
	}

	// unify the schema with the data and validate, like `cue vet`
	v := schema.Unify(data)
	if err := v.Err(); err != nil {
		panic(err)
	}
	if err := v.Validate(); err != nil {
		panic(err)
	}
}
