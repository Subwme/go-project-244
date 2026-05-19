package code

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedNested = `{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow: ` + `
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
}`

func TestGenDiffNestedJSON(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixtures/file1_nested.json",
		"testdata/fixtures/file2_nested.json",
		"stylish",
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedNested, result)
}

func TestGenDiffNestedYAML(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixtures/file1_nested.yml",
		"testdata/fixtures/file2_nested.yml",
		"stylish",
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedNested, result)
}

func TestGenDiffJSON(t *testing.T) {
	data, err := os.ReadFile("testdata/fixtures/result_json.json")
	assert.NoError(t, err)
	expected := strings.TrimRight(string(data), "\n")

	result, err := GenDiff(
		"testdata/fixtures/file1_nested.json",
		"testdata/fixtures/file2_nested.json",
		"json",
	)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGenDiffPlain(t *testing.T) {
	expected := "Property 'common.follow' was added with value: false\n" +
		"Property 'common.setting2' was removed\n" +
		"Property 'common.setting3' was updated. From true to null\n" +
		"Property 'common.setting4' was added with value: 'blah blah'\n" +
		"Property 'common.setting5' was added with value: [complex value]\n" +
		"Property 'common.setting6.doge.wow' was updated. From '' to 'so much'\n" +
		"Property 'common.setting6.ops' was added with value: 'vops'\n" +
		"Property 'group1.baz' was updated. From 'bas' to 'bars'\n" +
		"Property 'group1.nest' was updated. From [complex value] to 'str'\n" +
		"Property 'group2' was removed\n" +
		"Property 'group3' was added with value: [complex value]"

	result, err := GenDiff(
		"testdata/fixtures/file1_nested.json",
		"testdata/fixtures/file2_nested.json",
		"plain",
	)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
