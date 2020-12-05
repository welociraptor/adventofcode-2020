package passport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateHeight(t *testing.T) {
	assert.True(t, validateHeight("60in"))
	assert.True(t, validateHeight("180cm"))
	assert.False(t, validateHeight("190in"))
	assert.False(t, validateHeight("190"))
}

func TestValidateHairColor(t *testing.T) {
	assert.True(t, validateHairColor("#123abc"))
	assert.False(t, validateHairColor("#123abz"))
	assert.False(t, validateHairColor("123abc"))
}

func TestValidateEyeColor(t *testing.T) {
	assert.True(t, validateEyeColor("oth"))
	assert.False(t, validateEyeColor("lolll"))
}

func TestValidatePassportNumber(t *testing.T) {
	assert.True(t, validatePassportNumber("123456789"))
}
