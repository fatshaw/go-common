package go_common

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRe(t *testing.T) {
	t.Run("wordAEqualsA", func(t *testing.T) {
		assert.Equal(t, lit("a")("a"), true)
	})

	t.Run("wordBNotEqualsA", func(t *testing.T) {
		assert.Equal(t, lit("b")("a"), false)
	})

	t.Run("wordAAEqualsAA", func(t *testing.T) {
		assert.Equal(t, lit("aa")("aa"), true)
	})

	t.Run("wordAAANotEqualsAA", func(t *testing.T) {
		assert.Equal(t, lit("aaa")("aa"), false)
	})

	t.Run("Sequence(ab,cd) == abcd", func(t *testing.T) {
		assert.Equal(t, sequence(lit("ab"), lit("cd"))("abcd"), true)
	})

	t.Run("Sequence(abc,cd) != abcd", func(t *testing.T) {
		assert.Equal(t, sequence(lit("abc"), lit("cd"))("abcd"), false)
	})

	t.Run("either(abc,cd) == abc", func(t *testing.T) {
		assert.Equal(t, either(lit("abc"), lit("cd"))("abc"), true)
	})

	t.Run("either(abc,cd) == cd", func(t *testing.T) {
		assert.Equal(t, either(lit("abc"), lit("cd"))("cd"), true)
	})

	t.Run("either(abc,cd) != abccd", func(t *testing.T) {
		assert.Equal(t, either(lit("abc"), lit("cd"))("abccd"), false)
	})

	t.Run("any() == abccd", func(t *testing.T) {
		assert.Equal(t, any()("abccd"), true)
	})

	t.Run("oneof(abc) == a", func(t *testing.T) {
		assert.Equal(t, oneof("abc")("a"), true)
	})


}
