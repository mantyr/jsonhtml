# JsonHtml - генератор Html на основе валидного JSON объекта

[![Build Status]    (https://travis-ci.org/mantyr/jsonhtml.svg?branch=master)]      (https://travis-ci.org/mantyr/jsonhtml)
[![GoDoc]           (https://godoc.org/github.com/mantyr/jsonhtml?status.png)]      (https://godoc.org/github.com/mantyr/jsonhtml)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)] (LICENSE.md)
[![Go Report Card]  (https://goreportcard.com/badge/github.com/mantyr/jsonhtml)]    (https://goreportcard.com/report/github.com/mantyr/jsonhtml)

## Installation

    $ go get github.com/mantyr/jsonhtml

## Examples:

```Bash
go run ./*.go ./testdata/file1.json 
<ul><li><h3>Title #1</h3><div>Hello, World 1!</div></li><li><h3>Title #2</h3><div>Hello, World 2!</div></li></ul>

go run ./*.go ./testdata/file1.json ./testdata/file2.json 
<ul><li><h3>Title #1</h3><div>Hello, World 1!</div></li><li><h3>Title #2</h3><div>Hello, World 2!</div></li></ul>

go run ./*.go ./testdata/file1.json ./testdata/file2.json ./testdata/file3.json ./testdata/file4.json 
<ul><li><h3>Title #1</h3><div>Hello, World 1!</div></li><li><h3>Title #2</h3><div>Hello, World 2!</div></li></ul>
<ul><li><span>Title #1</span><content><ul><li><p>Example 1</p><header>header 1</header></li></ul></content></li><li><div>div 1</div></li></ul>
<p>hello1</p>
<p id="my-id" class="my-class">hello</p><p class="my-class1 my-class2">example&lt;a&gt;asd&lt;/a&gt;</p>


go run ./*.go no_file.json
2016/11/09 08:10:33 Convert error file no_file.json open no_file.json: no such file or directory
```

```GO
package main

import (
    "github.com/mantyr/jsonhtml"
    "flag"
    "fmt"
    "log"
)

func main() {
    flag.Parse()
    for _, file_address := range flag.Args() {
        s, err := jsonhtml.ConvertFile(file_address)

        if err != nil {
            log.Println("Convert error file", file_address, err.Error())
            continue
        }
        fmt.Println(s)
    }
}
```

```GO
package main

import (
    "github.com/mantyr/jsonhtml"
    "fmt"
)

func main() {
    js := `
    {
        "p.my-class#my-id": "hello",
        "p.my-class1.my-class2":"example<a>asd</a>"
    }
    `
    val, err := ConvertString(js)

    fmt.Println(val, err) // `<p id="my-id" class="my-class">hello</p><p class="my-class1 my-class2">example&lt;a&gt;asd&lt;/a&gt;</p>`
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr