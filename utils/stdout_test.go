package utils

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// captureStdout gives what f writes to stdout.
func captureStdout(t *testing.T, f func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	out := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		out <- string(b)
	}()
	old := os.Stdout
	os.Stdout = w
	func() {
		defer func() {
			os.Stdout = old
			w.Close()
		}()
		f()
	}()
	s := <-out
	r.Close()
	return s
}

// BL-270: the utils never write to stdout, where a script's output goes.

func TestIndexedAccessWritesNothingToStdout(t *testing.T) {
	out := captureStdout(t, func() {
		v, i, path, err := GetVarComponents("$list[0]")
		assert.NoError(t, err)
		assert.Equal(t, "$list", v)
		assert.Equal(t, 0, i)
		assert.Equal(t, "", path)

		v, i, path, err = GetVarComponents("$list[2].a[1].b")
		assert.NoError(t, err)
		assert.Equal(t, "$list", v)
		assert.Equal(t, 2, i)
		assert.Equal(t, "a.1.b", path)

		_, _, _, err = GetVarComponents("$list[x]")
		assert.Error(t, err)

		i, _, err = GetFirstIndexAndPath("$list.3.a")
		assert.NoError(t, err)
		assert.Equal(t, 3, i)
		_, _, err = GetFirstIndexAndPath("$list.x")
		assert.Error(t, err)
	})
	assert.Empty(t, out)
}

func TestJSONPathTokensWriteNothingToStdoutInDebug(t *testing.T) {
	s := NewJSONUtils()
	s.Dbg = true
	expr := "$list[0].token"
	out := captureStdout(t, func() {
		ok, path := s.GetJSONVariableExpressionTokens(&expr)
		assert.True(t, ok)
		assert.NotEmpty(t, *path)
	})
	assert.Empty(t, out)
}

func TestHTMLAndProcessHelpersWriteNothingToStdout(t *testing.T) {
	// no tasklist on the PATH: the failure is logged, not printed
	t.Setenv("PATH", t.TempDir())
	var logged bytes.Buffer
	log.SetOutput(&logged)
	defer log.SetOutput(os.Stderr)
	out := captureStdout(t, func() {
		_, err := GetJSONFromHtml("<html><body>BL270</body></html>")
		assert.NoError(t, err)
		assert.False(t, IsProcessRunningWindows("BL270.exe"))
	})
	assert.Empty(t, out)
	assert.Contains(t, logged.String(), "tasklist")
	assert.NotContains(t, logged.String(), "BL270")
}
