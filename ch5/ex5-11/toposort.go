package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string {
	"algorithms": { "data structures" },
	"calculus": { "linear algebra" },
	"compilers": { 
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": { "discrete math" },
	"databases": { "data structures" },
	"discrete math": { "intro to programming " },
	"formal languages": { "discrete math" },
	"networks": { "operating systems" },
	"operating systems": {
		"data structures",
		"computer organization",
	},
	"programming languages": {
		"data structures", 
		"computer organization",
	},
	"linear algebra": { "calculus" },
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	if cyclesExist(m) {
		fmt.Printf("\nWarning: Cycles exist!\n\n")
	}

	var order []string
	seen := make(map[string]bool)
	var visitAll func(items [] string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func cyclesExist(m map[string][]string) bool {
	for k, v := range m {
		for _, prereq := range v {
			for _, prereq2 := range m[prereq] {
				if k == prereq2 {
					return true
				}
			}
		}
	}
	return false
}
