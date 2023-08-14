package repository

import (
	"testing"
)

func TestSortVimColorSchemesBySearchTermMatch(t *testing.T) {
	t.Run("should leave the list as is if it's already sorted correctly", func(t *testing.T) {
		repository := Repository{
			VimColorSchemes: []VimColorScheme{
				{Name: "first_term"},
				{Name: "second_term"},
			},
		}

		repository.SortVimColorSchemesBySearchTermMatch("first")

		if repository.VimColorSchemes[0].Name != "first_term" {
			t.Errorf("Expected first_term, got %s", repository.VimColorSchemes[0].Name)
		}
	})

	t.Run("should move the first match to first index", func(t *testing.T) {
		repository := Repository{
			VimColorSchemes: []VimColorScheme{
				{Name: "first_term"},
				{Name: "second_term"},
			},
		}

		repository.SortVimColorSchemesBySearchTermMatch("second")

		if repository.VimColorSchemes[0].Name != "second_term" {
			t.Errorf("Expected second_term, got %s", repository.VimColorSchemes[0].Name)
		}
	})
}
