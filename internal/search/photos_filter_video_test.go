package search

import (
	"testing"

	"github.com/photoprism/photoprism/internal/form"
	"github.com/stretchr/testify/assert"
)

func TestPhotosQueryVideo(t *testing.T) {
	var f0 form.SearchPhotos

	f0.Query = "video:true"
	f0.Merged = true

	// Parse query string and filter.
	if err := f0.ParseQueryString(); err != nil {
		t.Fatal(err)
	}

	photos0, _, err := Photos(f0)

	if err != nil {
		t.Fatal(err)
	}
	assert.GreaterOrEqual(t, len(photos0), 3)

	t.Run("false > yes", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:yes"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), len(photos0))
		f.Query = "video:false"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos2, _, err2 := Photos(f)

		if err2 != nil {
			t.Fatal(err2)
		}
		assert.Greater(t, len(photos2), len(photos))
	})
	t.Run("StartsWithPercent", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"%gold\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("CenterPercent", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"I love % dog\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("EndsWithPercent", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"sale%\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("StartsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"&IlikeFood\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("CenterAmpersand", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Pets & Dogs\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("EndsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Light&\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("StartsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"'Family\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("CenterSingleQuote", func(t *testing.T) {
		var f form.SearchPhotos

		// Note: If the string in video starts with f/F, the txt package will assume it means false,
		f.Query = "video:\"Mother's Day\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("EndsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Ice Cream'\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("StartsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"*Forrest\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("CenterAsterisk", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"My*Kids\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("EndsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Yoga***\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("StartsWithPipe", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"|Banana\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("CenterPipe", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Red|Green\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("EndsWithPipe", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Blue|\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("StartsWithNumber", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"345 Shirt\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("CenterNumber", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Color555 Blue\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("EndsWithNumber", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Route 66\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("AndSearch", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Route 66 & Father's Day\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
	t.Run("OrSearch", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "video:\"Route %66 | *Father's Day\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), len(photos0))
	})
}
