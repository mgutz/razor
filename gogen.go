package razor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var GorazorNamespace = `"github.com/mgutz/razor"`

//------------------------------ Compiler ------------------------------ //
const (
	CMKP = iota
	CBLK
	CSTAT
)

func getValStr(e interface{}) string {
	switch v := e.(type) {
	case *Ast:
		return v.TagName
	case Token:
		if !(v.Type == AT || v.Type == AT_COLON) {
			return v.Text
		}
		return ""
	default:
		panic(e)
	}
}

type Part struct {
	ptype int
	value string
}

type Compiler struct {
	ast      *Ast
	buf      string //the final result
	layout   string
	firstBLK int
	params   string
	parts    []Part
	imports  map[string]bool
	options  Option
	dir      string
	file     string
	result   string
}

func (self *Compiler) addPart(part Part) {
	if len(self.parts) == 0 {
		self.parts = append(self.parts, part)
		return
	}
	last := &self.parts[len(self.parts)-1]
	if last.ptype == part.ptype {
		last.value += part.value
	} else {
		self.parts = append(self.parts, part)
	}
}

func (self *Compiler) genPart() {
	res := ""
	for _, p := range self.parts {
		if p.ptype == CMKP && p.value != "" {
			// do some escapings
			p.value = strings.Replace(p.value, `\n`, `\\n`, -1)
			p.value = strings.Replace(p.value, "\n", `\n`, -1)
			p.value = strings.Replace(p.value, `"`, `\"`, -1)
			for strings.HasSuffix(p.value, "\\n") {
				p.value = p.value[:len(p.value)-2]
			}
			if p.value != "\\n" && p.value != "" {
				res += "_buffer.WriteString(\"" + p.value + "\")\n"
			}
		} else if p.ptype == CBLK {
			if strings.HasPrefix(p.value, "{") &&
				strings.HasSuffix(p.value, "}") {
				p.value = p.value[1 : len(p.value)-2]
			}
			res += p.value + "\n"
		} else {
			res += p.value
		}
	}
	self.buf = res
}

func makeCompiler(ast *Ast, options Option, input string) *Compiler {
	dir := filepath.Base(filepath.Dir(input))
	file := Capitalize(strings.Replace(filepath.Base(input), gz_extension, "", 1))
	cp := &Compiler{ast: ast, buf: "",
		layout: "", firstBLK: 0,
		params: "()", parts: []Part{},
		imports: map[string]bool{},
		options: options,
		dir:     dir,
		file:    file,
	}
	return cp
}

func (cp *Compiler) visitBLK(child interface{}, ast *Ast) {
	blk := getValStr(child)
	cp.addPart(Part{CBLK, blk})
}

func (cp *Compiler) visitMKP(child interface{}, ast *Ast) {
	cp.addPart(Part{CMKP, getValStr(child)})
}

// First block contains imports and parameters, specific action for layout,
// NOTE, layout have some conventions.
func (cp *Compiler) visitFirstBLK(blk *Ast) {
	pre := cp.buf
	cp.buf = ""

	first := ""
	backup := cp.parts
	cp.parts = []Part{}
	cp.visitAst(blk)
	cp.genPart()
	first, cp.buf = cp.buf, pre
	cp.parts = backup

	isImport := false
	lines := strings.SplitN(first, "\n", -1)
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if strings.HasPrefix(l, "import") || strings.HasPrefix(l, "+import") {
			isImport = true
			continue
		}
		if l == ")" {
			isImport = false
			continue
		}

		if isImport {
			cp.imports[l] = true
		} else if strings.HasPrefix(l, "+func") {
			vname := l[5:]
			cp.params = vname
		} else if strings.HasPrefix(l, "+return ") {
			vname := l[8:]
			vname = strings.Replace(vname, "VIEW", "_buffer", -1)
			cp.result = strings.Replace(vname, "SECTIONS", "_sections", -1)
		} else if l != "" {
			cp.addPart(Part{CSTAT, l + "\n"})
		}
	}

}

func (cp *Compiler) visitExp(child interface{}, parent *Ast, idx int, isHomo bool) {
	start := ""
	end := ""
	ppNotExp := true
	ppChildCnt := len(parent.Children)
	//pack := cp.dir
	//htmlEsc := cp.options["htmlEscape"]
	if parent.Parent != nil && parent.Parent.Mode == EXP {
		ppNotExp = false
	}
	val := getValStr(child)

	if ppNotExp && idx == 0 {
		start = "_buffer.WriteSafe(" + start
	}
	if ppNotExp && idx == ppChildCnt-1 {
		end += ")\n"
	}

	v := start
	if val == "raw" {
		v += end
	} else {
		v += val + end
	}
	cp.addPart(Part{CSTAT, v})
}

func (cp *Compiler) visitAst(ast *Ast) {
	cp.imports[GorazorNamespace] = true

	switch ast.Mode {
	case MKP:
		cp.firstBLK = 1
		for _, c := range ast.Children {
			if _, ok := c.(Token); ok {
				cp.visitMKP(c, ast)
			} else {
				cp.visitAst(c.(*Ast))
			}
		}
	case BLK:
		if cp.firstBLK == 0 {
			cp.firstBLK = 1
			cp.visitFirstBLK(ast)
		} else {
			for _, c := range ast.Children {
				if _, ok := c.(Token); ok {
					cp.visitBLK(c, ast)
				} else {
					cp.visitAst(c.(*Ast))
				}
			}
		}
	case EXP:
		cp.firstBLK = 1
		nonExp := ast.hasNonExp()
		for i, c := range ast.Children {
			if _, ok := c.(Token); ok {
				cp.visitExp(c, ast, i, !nonExp)
			} else {
				cp.visitAst(c.(*Ast))
			}
		}
	case PRG:
		for _, c := range ast.Children {
			cp.visitAst(c.(*Ast))
		}
	}
}

// TODO, this is dirty now
func (cp *Compiler) processLayout() {
	lines := strings.SplitN(cp.buf, "\n", -1)
	out := ""
	sections := []string{}
	scope := 0
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if strings.HasPrefix(l, "section") && strings.HasSuffix(l, "{") {
			name := l
			name = strings.TrimSpace(name[7 : len(name)-1])
			out += "\n" + name + " := func() *razor.SafeBuffer {\n"
			out += "_buffer := razor.NewSafeBuffer()\n"
			scope = 1
			sections = append(sections, name)
		} else if scope > 0 {
			if strings.HasSuffix(l, "{") {
				scope++
			} else if strings.HasSuffix(l, "}") {
				scope--
			}
			if scope == 0 {
				out += "return _buffer\n}\n"
				//scope = 0
			} else {
				out += l + "\n"
			}
		} else {
			out += l + "\n"
		}
	}
	cp.buf = out

	if len(sections) > 0 {
		cp.buf += "_sections := make(razor.Sections)\n"
		for _, section := range sections {
			cp.buf += fmt.Sprintf("_sections[\"%s\"] = %s()\n", section, section)
		}
	}

	if cp.result != "" {
		cp.buf += "_buffer = " + cp.result
	}
	cp.buf += "\n return _buffer\n}\n"

}

func (cp *Compiler) visit() {
	cp.visitAst(cp.ast)
	cp.genPart()

	pack := cp.dir
	fun := cp.file

	//cp.imports[`"bytes"`] = true
	head := "// DO NOT EDIT! Auto-generated by github.com/mgutz/razor\n\n"
	head += "package " + pack + "\n import (\n"
	for k := range cp.imports {
		head += k + "\n"
	}

	// adds comment to appease golint
	head += "\n)\n// " + fun + " is generated\nfunc " + fun + cp.params
	head += ` *razor.SafeBuffer {
				_buffer := razor.NewSafeBuffer()
				locals := razor.Locals
				if locals != nil {
					// avoids not declared error if locals is not used
				}
			`
	cp.buf = head + cp.buf
	cp.processLayout()
}

func run(path string, Options Option) (*Compiler, error) {

	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	text := string(content)

	lex := &Lexer{text, Tests}

	res, err := lex.Scan()
	if err != nil {
		return nil, err
	}

	//DEBUG
	if Options["Debug"] != nil {
		fmt.Println("------------------- TOKEN START -----------------")
		for _, elem := range res {
			elem.P()
		}
		fmt.Println("--------------------- TOKEN END -----------------\n")
	}

	parser := &Parser{&Ast{}, nil, res, []Token{}, false, UNK}
	err = parser.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//DEBUG
	if Options["Debug"] != nil {
		fmt.Println("--------------------- AST START -----------------")
		parser.ast.debug(0, 20)
		fmt.Println("--------------------- AST END -----------------\n")
		if parser.ast.Mode != PRG {
			panic("TYPE")
		}
	}
	cp := makeCompiler(parser.ast, Options, path)
	cp.visit()
	return cp, nil
}

func generate(path string, output string, Options Option) error {
	cp, err := run(path, Options)
	if err != nil || cp == nil {
		panic(err)
	}
	err = ioutil.WriteFile(output, []byte(cp.buf), 0644)
	cmd := exec.Command("gofmt", "-s", "-w", output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("gofmt: ", err)
		return err
	}
	if Options["Debug"] != nil {
		content, _ := ioutil.ReadFile(output)
		fmt.Println(string(content))
	}
	return err
}

//------------------------------ API ------------------------------
const (
	go_extension = ".go"
	gz_extension = ".go.html"
)

// Generate from input to output file,
// gofmt will trigger an error if it fails.
func GenFile(input string, output string, options Option) error {
	//fmt.Printf("input=%s output=%s\n", input, output)
	outdir := filepath.Dir(output)
	if !exists(outdir) {
		os.MkdirAll(outdir, 0775)
	}
	return generate(input, output, options)
}

// Generate from directory to directory, Find all the files with extension
// of .go.html and generate it into target dir.
func GenFolder(indir string, outdir string, options Option) (err error) {
	if !exists(indir) {
		return errors.New("Input directory does not exsits")
	} else {
		if err != nil {
			return err
		}
	}
	//Make it
	if !exists(outdir) {
		os.MkdirAll(outdir, 0775)
	}

	incdir_abs, _ := filepath.Abs(indir)
	outdir_abs, _ := filepath.Abs(outdir)

	paths := []string{}

	visit := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			//Just do file with extension .go.html
			if !strings.HasSuffix(path, gz_extension) {
				return nil
			}
			filename := filepath.Base(path)
			if strings.HasPrefix(filename, ".#") {
				return nil
			}
			paths = append(paths, path)
		}
		return nil
	}

	fun := func(path string, res chan<- string) {
		//adjust with the abs path, so that we keep the same directory hierarchy
		input, _ := filepath.Abs(path)
		output := strings.Replace(input, incdir_abs, outdir_abs, 1)
		output = strings.Replace(output, gz_extension, go_extension, -1)
		err := GenFile(path, output, options)
		if err != nil {
			res <- fmt.Sprintf("%s -> %s", path, output)
			os.Exit(2)
		}
		res <- fmt.Sprintf("%s -> %s", path, output)
	}

	err = filepath.Walk(indir, visit)
	runtime.GOMAXPROCS(runtime.NumCPU())
	result := make(chan string, len(paths))

	for w := 0; w < len(paths); w++ {
		go fun(paths[w], result)
	}
	for i := 0; i < len(paths); i++ {
		<-result
	}
	return
}