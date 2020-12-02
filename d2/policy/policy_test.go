package policy_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/welociraptor/adventofcode-2020/v2/d2/policy"

	"testing"
)

func TestValidate(t *testing.T) {
	in := "4-5 z: zzdjz"
	b, err := policy.Validate(in)

	assert.False(t, b)
	assert.NoError(t, err)
}
