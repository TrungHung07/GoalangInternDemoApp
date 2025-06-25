package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func generateTableMap() {
	e := entc.Generate("./ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureEntQL,
			gen.FeaturePrivacy,
			gen.FeatureModifier,
		},
	})

	if e != nil {
		log.Fatalf("eror when gen %v", e)
	}
}

func main() {
	generateTableMap()
}
