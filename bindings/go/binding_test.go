package tree_sitter_org_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_org "github.com/tree-sitter/tree-sitter-org/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_org.Language())
	if language == nil {
		t.Errorf("Error loading Org mode grammar")
	}
}
