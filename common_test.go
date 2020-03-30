package liner

import (
	"os"
	"testing"
)

// LOCAL=true go test -v -timeout 30s github.com/peterh/liner -run "^(TestCompleter)$"
func TestCompleter(t *testing.T) {
	if local := os.Getenv("LOCAL"); local != "true" {
		t.Skip()
	}
	l := NewLiner()
	l.commonState.columns = 80 // make space to prevent internal error
	f := func(line string) []string {
		return []string{line + "-completed"}
	}
	l.SetCompleter(f)
	completed, pos, v, err := l.tabComplete([]rune("$"), []rune("line"), 4)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(completed, pos, v)
}

// LOCAL=true go test -v -timeout 30s github.com/peterh/liner -run "^(TestWordCompleter)$"
func TestWordCompleter(t *testing.T) {
	if local := os.Getenv("LOCAL"); local != "true" {
		t.Skip()
	}
	l := NewLiner()
	l.commonState.columns = 80 // make space to prevent internal error
	f := func(line string, pos int) (head string, completions []string, tail string) {
		t.Logf("line:%q,pos:%d", line, pos)
		return "Hello, ", []string{"world", "Word"}, "!!!"
	}
	l.SetWordCompleter(f)
	completed, pos, v, err := l.tabComplete([]rune("$"), []rune("Hello, wo!!!"), 9)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(completed, pos, v)
}

// LOCAL=true go test -v -timeout 30s github.com/peterh/liner -run "^(TestPosCompleter)$"
func TestPosCompleter(t *testing.T) {
	if local := os.Getenv("LOCAL"); local != "true" {
		t.Skip()
	}
	l := NewLiner()
	l.commonState.columns = 80 // make space to prevent internal error
	f := func(line string, pos int) (head string, completions []WordAndPos, tail string) {
		t.Logf("line:%q,pos:%d", line, pos)
		return "Hello, ", []WordAndPos{WordAndPos{Word: "world", Pos: 0}, WordAndPos{Word: "Word", Pos: 0}}, "!!!"
	}
	l.SetPosCompleter(f)
	completed, pos, v, err := l.tabComplete([]rune("$"), []rune("Hello, wo!!!"), 9)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(completed, pos, v)
}
