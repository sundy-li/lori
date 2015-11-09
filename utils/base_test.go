package utils

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamedRegexp(t *testing.T) {
	var url = "/aa/bb/cc"
	var reg = regexp.MustCompile(`/(?P<aa>\S+)/(?P<bb>\S+)/(?P<cc>\S+)`)

	ng, flag := NamedRegexpGroup(url, reg)

	assert.NotNil(t, ng)

	assert.True(t, flag)

	fmt.Printf("%#v\n", ng)

}
