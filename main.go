package main

import (
	"fmt"

	"github.com/zulubit/mimi/pkg/mimi"
)

func main() {

	fmt.Println("Mimi is alive.")

	mimi := mimi.MimiExtend{}

	mimi.Start()
}

// TODO:
// 1. figure out how to merge in seo - build seo struct, make sure config files have seo fields, reserve fields in page-config, in the template, check what fields exist and conditionally merge them in. Prefer the page config over the global one.
// 2. Add core package (Core should hold a Mimi struct, that serves as a facade for routes, template functions and logger)
// 4. Implement a post/page loop template function. Fun.
// 5. Provide a very basic admin dash.
