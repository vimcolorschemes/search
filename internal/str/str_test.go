package str

import "testing"

func TestNormalize(t *testing.T) {
	t.Run("should remove unsafe characters", func(t *testing.T) {
		if result := Normalize(`[]`); result != "" {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, "")
		}

		if result := Normalize(`()`); result != "" {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, "")
		}

		if result := Normalize(`{}`); result != "" {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, "")
		}

		if result := Normalize(`\`); result != "" {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, "")
		}

		if result := Normalize(`$`); result != "" {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, "")
		}

		if result := Normalize(`^`); result != "" {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, "")
		}
	})

	t.Run("should replace delimiting characters with spaces", func(t *testing.T) {
		if result := Normalize(`./.`); result != ". ." {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, ". .")
		}
	})

	t.Run("should escape unsafe but useful characters", func(t *testing.T) {
		if result := Normalize(`+`); result != `\+` {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, `\+`)
		}

		if result := Normalize(`|`); result != `\|` {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, `\|`)
		}

		if result := Normalize(`?`); result != `\?` {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, `\?`)
		}
	})

	t.Run("should trim the final result", func(t *testing.T) {
		if result := Normalize(`/. `); result != `.` {
			t.Errorf("Incorrect result for Normalize, got: %s, want: %s", result, `.`)
		}
	})
}
