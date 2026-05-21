package code

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var expectedFlat = `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

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

var expectedFlatPlain = `Property 'follow' was removed
Property 'proxy' was removed
Property 'timeout' was updated. From 50 to 20
Property 'verbose' was added with value: true`

var expectedNestedPlain = "Property 'common.follow' was added with value: false\n" +
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

func readFixture(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile("testdata/fixtures/" + name)
	require.NoError(t, err)
	return strings.TrimRight(string(data), "\n")
}

// Stylish — flat files

func TestGenDiffFlatJSON(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixtures/file1.json",
		"testdata/fixtures/file2.json",
		"stylish",
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedFlat, result)
}

func TestGenDiffFlatYAML(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixtures/file1.yml",
		"testdata/fixtures/file2.yml",
		"stylish",
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedFlat, result)
}

// Stylish — nested files

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

// Plain format

func TestGenDiffFlatPlain(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixtures/file1.json",
		"testdata/fixtures/file2.json",
		"plain",
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedFlatPlain, result)
}

func TestGenDiffNestedPlain(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixtures/file1_nested.json",
		"testdata/fixtures/file2_nested.json",
		"plain",
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedNestedPlain, result)
}

func TestGenDiffNestedPlainYAML(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixtures/file1_nested.yml",
		"testdata/fixtures/file2_nested.yml",
		"plain",
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedNestedPlain, result)
}

// JSON format

func TestGenDiffJSON(t *testing.T) {
	expected := readFixture(t, "result_json.json")
	result, err := GenDiff(
		"testdata/fixtures/file1_nested.json",
		"testdata/fixtures/file2_nested.json",
		"json",
	)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGenDiffNestedJSONYAML(t *testing.T) {
	expected := readFixture(t, "result_json.json")
	result, err := GenDiff(
		"testdata/fixtures/file1_nested.yml",
		"testdata/fixtures/file2_nested.yml",
		"json",
	)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

// Error cases

func TestGenDiffUnknownFormat(t *testing.T) {
	_, err := GenDiff(
		"testdata/fixtures/file1.json",
		"testdata/fixtures/file2.json",
		"xml",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown format")
}

func TestGenDiffFileNotFound(t *testing.T) {
	_, err := GenDiff(
		"testdata/fixtures/nonexistent.json",
		"testdata/fixtures/file2.json",
		"stylish",
	)
	assert.Error(t, err)
}

func TestGenDiffInvalidJSON(t *testing.T) {
	_, err := GenDiff(
		"testdata/fixtures/invalid.json",
		"testdata/fixtures/file2.json",
		"stylish",
	)
	assert.Error(t, err)
}
