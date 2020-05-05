package gaddag_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_CreateGraph(t *testing.T) {
	gaddagRoot, err := gaddag.CreateGraph("../exampleData/tiny_english.txt")

	var expected gaddag.Node
	test_utils.GetGoldenFileJSON(t, gaddagRoot, &expected, t.Name(), *update)

	assert.Equal(t, err, nil)
	assert.Equal(t, &expected, gaddagRoot)
}

func Test_IsWordValid(t *testing.T) {
	gaddagRoot, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")

	isOk, err := gaddagRoot.IsWordValid("w.or")
	assert.Equal(t, false, isOk)
	assert.Equal(t, errors.New("Word wor is not in dictionary"), err)
	isOk, err = gaddagRoot.IsWordValid("w.ord")
	assert.Equal(t, true, isOk)
	assert.Equal(t, nil, err)
	isOk, err = gaddagRoot.IsWordValid("w.ordX")
	assert.Equal(t, false, isOk)
	assert.Equal(t, errors.New("Word wordX is not in dictionary"), err)
	isOk, err = gaddagRoot.IsWordValid("w.orthless")
	assert.Equal(t, true, isOk)
	assert.Equal(t, nil, err)
	isOk, err = gaddagRoot.IsWordValid("w.orthleks")
	assert.Equal(t, false, isOk)
	assert.Equal(t, errors.New("Word worthleks is not in dictionary"), err)
	isOk, err = gaddagRoot.IsWordValid("ob.ss")
	assert.Equal(t, true, isOk)
	assert.Equal(t, nil, err)
}

func Benchmark_CreateGraph_2kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gaddag.CreateGraph("../exampleData/2k_english.txt")
	}
}

func Benchmark_CreateGraph_20kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gaddag.CreateGraph("../exampleData/20k_english.txt")
	}
}

func Benchmark_CreateGraph_280kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gaddag.CreateGraph("../exampleData/collins_official_scrabble_2019.txt")
	}
}
