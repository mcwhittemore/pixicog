package pixicog

import (
  "bytes"
  "go/ast"
	"go/parser"
	"go/token"
  "go/format"
  "fmt"
  "os"
  "os/exec"
  "crypto/sha1"
  "io/ioutil"
)

func main() {
  msg, err := run(os.Args[1])
  fmt.Printf("%s", msg)
  if err != nil {
    fmt.Println("err", err)
    panic(1)
  }
}

func run(fn string) ([]byte, error) {
  mainFile, err := buildMainFile(fn)
  if err != nil {
    return nil, err
  }

  tmpName, err := saveTemp([]byte(mainFile))
  if err != nil {
    os.Remove(tmpName)
    return nil, err
  }

  msg, err := gorun(tmpName)
  os.Remove(tmpName)
  if err != nil {
    return nil, err
  }

  return msg, err
}

func buildMainFile(fn string) (string, error) {
  fset := token.NewFileSet()
  f, err := parser.ParseFile(fset, fn, nil, 0)
  if err != nil {
    return "", err
  }

  funcs := getFuncs(f, fset);

  var buf bytes.Buffer
  format.Node(&buf, fset, f)

  mainStr, err := buildMainFunc(funcs)
  if err != nil {
    return "", err
  }

  ctx := fmt.Sprintf("%s\n%s", buf.String(), mainStr)

  return ctx, nil
}

func buildMainFunc(funcs [][]string) (string, error) {
  params := ast.FieldList{token.NoPos,nil,token.NoPos}

  statments, err := buildStmts(funcs)
	if err != nil {
    return "", err
	}

  funcName := ast.NewIdent("main")
  funcType := ast.FuncType{token.NoPos,&params,nil}
  funcBody := ast.BlockStmt{token.NoPos,statments,token.NoPos}

  funcDecl := ast.FuncDecl{nil, nil, funcName, &funcType, &funcBody}

  fset := token.NewFileSet()

  var buf bytes.Buffer
  err = format.Node(&buf, fset, &funcDecl)
	if err != nil {
    return "", err
	}

  return buf.String(), nil
}

func buildStmts(funcs [][]string) (statments []ast.Stmt, err error) {

  for i := 0; i < len(funcs); i++ {
    logExpr, _ := parser.ParseExpr(fmt.Sprintf(`fmt.Println("%s -> %s")`, funcs[0][0], funcs[0][1]))
    statments = append(statments, &ast.ExprStmt{logExpr})

    fnExpr, _ := parser.ParseExpr(fmt.Sprintf("%s()", funcs[0][0]))
    statments = append(statments, &ast.ExprStmt{fnExpr})
  }

  return statments, nil
}

func getFuncs(file *ast.File, fset *token.FileSet) [][]string {
  var funcs [][]string

  for _, d := range file.Decls {
    if f, ok := d.(*ast.FuncDecl); ok {
      var oneFunc = make([]string, 2)
      oneFunc[0] = f.Name.Name
      oneFunc[1] = getFuncHash(f, fset)
      funcs = append(funcs, oneFunc)

    }
  }

  return funcs
}

func getFuncHash(fun *ast.FuncDecl, fset *token.FileSet) string {
  var buf bytes.Buffer
	format.Node(&buf, fset, fun)

  h := sha1.New()
  h.Write([]byte(buf.String()))
  return fmt.Sprintf("%x", h.Sum(nil))
}

func saveTemp(content []byte) (string, error) {
  tmpFile, err := ioutil.TempFile("", "pixicog-script*.go")
  if err != nil {
    return "", err
  }

  if _, err := tmpFile.Write(content); err != nil {
    return "", err
  }

  if err := tmpFile.Close(); err != nil {
    return "", err
  }

  return tmpFile.Name(), nil
}


func gorun(fn string) ([]byte, error) {
  cmd := exec.Command("go", "run", fn)
  msg, err := cmd.CombinedOutput()
  return msg, err
}
