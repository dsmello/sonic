package files

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	t.Run("Test the read method.", func(t *testing.T) {

		tests := []string{
			"1- I want data",
			"2- multiline \n yep",
			"3- More lines \n\n\n\t skip",
			"4- Large line without break lines, yep I dont had breack lines",
		}

		for _, test := range tests {
			got := Read(strings.NewReader(test))
			assert.EqualValues(t, test, got)
		}

	})
}
