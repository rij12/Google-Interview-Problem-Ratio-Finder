package conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConverstionFromTextFile(t *testing.T) {

	r := ReadConverstionFromTextFile("../data.txt")
	print(r)

}

func TestCreateNodeFromString(t *testing.T) {

	testString := "centimeter, inch, 2.540"
	origin, dest, raito := CreateNodeFromString(testString)

	assert.Equal(t, "centimeter", origin)
	assert.Equal(t, "inch", dest)
	assert.Equal(t, 2.540, raito)

}
