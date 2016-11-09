package jsonhtml

import (
    "testing"
)

func Test1(t *testing.T) {
    js := `
    [
      {
        "span": "Title #1",
        "content": [
          {
            "p": "Example 1",
            "header": "header 1"
          }
        ]
      },
      {"div": "div 1"}
    ]
    `
    val, err := ConvertString(js)
    if err != nil {
        t.Errorf("Error convert json to html, %q", err.Error())
    }
    if val != `<ul><li><span>Title #1</span><content><ul><li><p>Example 1</p><header>header 1</header></li></ul></content></li><li><div>div 1</div></li></ul>` {
        t.Errorf("Expected another value %q", val)
    }
}

func Test2(t *testing.T) {
    js := `
    {
        "p":"hello1"
    }
    `
    val, err := ConvertString(js)
    if err != nil {
        t.Errorf("Error convert json to html, %q", err.Error())
    }
    if val != `<p>hello1</p>` {
        t.Errorf("Expected another value %q", val)
    }
}

func Test3(t *testing.T) {
    js := `
    {
        "p.my-class#my-id": "hello",
        "p.my-class1.my-class2":"example<a>asd</a>"
    }
    `
    val, err := ConvertString(js)
    if err != nil {
        t.Errorf("Error convert json to html, %q", err.Error())
    }
    if val != `<p id="my-id" class="my-class">hello</p><p class="my-class1 my-class2">example&lt;a&gt;asd&lt;/a&gt;</p>` {
        t.Errorf("Expected another value %q", val)
    }
}