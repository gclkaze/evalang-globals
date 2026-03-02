package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiDimExpressionLeavesReplaced(t *testing.T) {
	assert.Equal(t, "", ReplaceBracketsWithDotsInExpr(""))
	assert.Equal(t, "$json", ReplaceBracketsWithDotsInExpr("$json"))
	assert.Equal(t, "$json.y.z.e", ReplaceBracketsWithDotsInExpr("$json.y.z.e"))
	assert.Equal(t, "$json.1.2.3.y.z.100.1.2.z.1.2.4", ReplaceBracketsWithDotsInExpr("$json[1][2][3].y.z[100][1][2].z[1][2][4]"))
	assert.Equal(t, "$json.a.10.y.z.100.1.2", ReplaceBracketsWithDotsInExpr("$json.a[10].y.z[100][1][2]"))
	assert.Equal(t, "$json.a.10.y.z.100.1.2.x.y.z", ReplaceBracketsWithDotsInExpr("$json.a[10].y.z[100][1][2].x.y.z"))
	//	assert.Equal(t, "$json. a .10 . y .z .100 .1.2.x.y.z", ReplaceBracketsWithDotsInExpr("$json. a [10] . y .z [100] [1] [2] .x.y.z"))
}

func TestIndexAndPathExtraction(t *testing.T) {
	res, path, err := GetFirstIndexAndPath("$first.2")
	assert.Equal(t, 2, res)
	assert.Equal(t, "", path)
	assert.Nil(t, err)

	res, path, err = GetFirstIndexAndPath("$first.60.1.2.x.y.10")
	assert.Equal(t, 60, res)
	assert.Nil(t, err)
	assert.Equal(t, "1.2.x.y.10", path)

}
