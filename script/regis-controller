package main

import (
    // "os"
    "fmt"
    "flag"
    "go/ast"
    "go/parser"
    "go/token"
    //"go/printer"
    //"container/vector"
)

var c = map[string]string{}
var a = map[string]string{}

func parseFile(filename string) {
    var fset = token.NewFileSet()
    astf, error := parser.ParseFile(fset, filename, nil, 0)
    if error != nil  {
        return
    }
    pname := astf.Name
    c[filename[0:len(filename)-3]] = pname.String()
    //fmt.Printf("%s\n", pname)

    for _, d:= range astf.Decls {
        if fd, ok:= d.(*ast.FuncDecl); ok {
            name := fd.Name.String()
            action := fmt.Sprintf("%s/%c%s", pname, 'a'+(name[0]-'A'),name[1:len(name)])
            a[action] = pname.String() + "." + name
        }
    }
}

func genInit() {
    fmt.Printf("package controller\n\n")
    fmt.Printf("import \"reflect\"\n")
    for k,_ := range c {
        fmt.Printf("import \"%s\"\n", k)
    }
    fmt.Printf(`
var Controllers = map[string]*reflect.FuncValue{}

func register(s string, c interface{}) {
    v := reflect.NewValue(c).(*reflect.FuncValue)
    Controllers[s] = v
}

func init() {
`)
    for k,v := range a {
        fmt.Printf("    register(\"%s\", %s)\n", k, v)
    }
    fmt.Printf("}\n",)
}

func main() {
    flag.Parse()
    narg := flag.NArg()
    if narg != 1 {
        fmt.Printf("-\n")
        return
    }
    filename := flag.Arg(0)
    parseFile(filename)
    // fmt.Printf("%s\n", c)
    // fmt.Printf("%s\n", a)
    genInit()
}

