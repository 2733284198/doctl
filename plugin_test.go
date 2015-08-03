package doit

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_argSlicer(t *testing.T) {
	cases := []struct {
		input    []string
		expected [][]string
	}{
		{
			input:    []string{"foo=1", "bar=2"},
			expected: [][]string{{"foo", "1"}, {"bar", "2"}},
		},
	}

	for _, c := range cases {
		got := argSlicer(c.input)
		assert.Equal(t, c.expected, got)
	}
}

func Test_loadPlugins(t *testing.T) {
	dir, err := ioutil.TempDir(os.TempDir(), "lp")
	assert.NoError(t, err)

	defer os.Remove(dir)

	d := []byte{}
	err = ioutil.WriteFile(filepath.Join(dir, "doit-plugin-test"), d, 0644)
	err = ioutil.WriteFile(filepath.Join(dir, "doit-not-a-plugin"), d, 0644)
	assert.NoError(t, err)

	ogPath := defaultPluginPaths
	defer func(p []string) {
		defaultPluginPaths = p
	}(ogPath)

	defaultPluginPaths = []string{dir}

	plugins := loadPlugins()
	assert.Equal(t, 1, len(plugins))
}

func TestPluginSummary(t *testing.T) {
	ogPluginFactory := pluginFactory
	defer func() {
		pluginFactory = ogPluginFactory
	}()
	pluginFactory = func(path string) Command {
		return NewMockCommand(path)
	}

	p := newPlugin("/bin", "test-plugin")
	_, err := p.Summary()
	assert.NoError(t, err)
}

func TestPluginStart(t *testing.T) {
	ogPluginFactory := pluginFactory
	defer func() {
		pluginFactory = ogPluginFactory
	}()
	pluginFactory = func(path string) Command {
		return NewMockCommand(path)
	}

	p := newPlugin("/bin", "doit-plugin-test")
	err := p.Exec("55")
	assert.NoError(t, err)

	err = p.Kill()
	assert.NoError(t, err)
}
