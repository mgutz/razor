package razor

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCap(t *testing.T) {
	if Capitalize("hello") != "Hello" {
		t.Error()
	}
	if Capitalize("0hello") != "0hello" {
		t.Error()
	}
}

func TestLexer(t *testing.T) {
	text := "case do func var switch"
	lex := &Lexer{text, Tests}
	res, err := lex.Scan()
	if err != nil {
		t.Error(err)
	}
	if len(res) != 10 {
		t.Error("token number")
	}
	for i, x := range res {
		if i%2 == 0 && x.Type != KEYWORD {
			t.Error("KEYWORD", x)
		}
	}
	text = "case casex do do3 func func_ var var+ "
	lex = &Lexer{text, Tests}
	res, err = lex.Scan()
	if err != nil {
		t.Error(err)
	}
	if len(res) != 18 {
		t.Error(err)
	}
	for i, x := range res {
		if i == 0 || i == 4 || i == 8 || i == 12 || i == 14 {
			if x.Type != KEYWORD {
				t.Error("KEYWORD")
			}
		} else if x.Type == KEYWORD {
			t.Error("Should NOT KEYWORD", x)
		}
	}
}

func TestGenerate(t *testing.T) {
	casedir, _ := filepath.Abs(filepath.Dir("./cases/"))

	visit := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			name := strings.Replace(path, ".go.html", ".go", 1)
			cmp := strings.Replace(name, "/cases/", "/test/", -1)
			dirname := filepath.Dir(cmp)
			log := filepath.Join(dirname, "_"+filepath.Base(cmp))
			if !exists(dirname) {
				os.MkdirAll(dirname, 0755)
			}
			option := Option{}
			GenFile(path, log, option)
			if !exists(cmp) {
				t.Error("No cmp:", cmp)
			} else if !exists(log) {
				t.Error("No log:", log)
			} else {
				//compare the log file and cmp file
				_cmp, _e1 := ioutil.ReadFile(cmp)
				_log, _e2 := ioutil.ReadFile(log)
				if _e1 != nil || _e2 != nil {
					t.Error("Reading")
				} else if string(_cmp) != string(_log) {
					t.Error("MISMATCH:", log, cmp)
				} else {
					t.Log("PASS")
				}
			}
		}
		return nil
	}
	err := filepath.Walk(casedir, visit)
	if err != nil {
		t.Error("walk")
	}
}
