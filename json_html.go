package jsonhtml

import (
    "encoding/json"
    "strings"
    "errors"
    "strconv"
    "io/ioutil"
    "bytes"
    "fmt"
    "os"
)


func Convert(data []byte) (s string, err error) {
    var v interface{}

    err = json.Unmarshal(data, &v)
    if err != nil {
        return
    }
    s, err = Encode(v)
    return
}

func ConvertString(data string) (string, error) {
    return Convert([]byte(data))
}

func ConvertFile(filename string) (s string, err error) {
    var file *os.File

    file, err = os.Open(filename)
    if err != nil {
        return
    }
    defer file.Close()

    var data []byte
    data, err = ioutil.ReadAll(file)
    if err != nil {
        return
    }
    return Convert(data)
}

func Encode(v interface{}) (s string, err error) {
    buf := new(bytes.Buffer)

    err = encodeValue(buf, v)
    s   = buf.String()

    return
}

var filter_string = strings.NewReplacer("<", "&lt;", ">", "&gt;")

func filterString(s string) string {
    s = filter_string.Replace(s)
    return s
}

func encodeValue(buf *bytes.Buffer, v interface{}) (err error) {
    switch t := v.(type) {
        default:
            err = fmt.Errorf("Unexpected type %T", t)
        case bool:
            if t {
                buf.WriteString("true")
            } else {
                buf.WriteString("false")
            }
        case nil:
            buf.WriteString("NULL")
        case int, int64, int32, int16, int8:
            strValue := fmt.Sprintf("%v", t)
            buf.WriteString(strValue)
        case float64:
            strValue := strconv.FormatFloat(t, 'f', -1, 64)
            buf.WriteString(strValue)
        case float32:
            strValue := strconv.FormatFloat(float64(t), 'f', -1, 32)
            buf.WriteString(strValue)
        case string:
            buf.WriteString(filterString(t))
        case []interface{}:
            buf.WriteString("<ul>")
            for _, value := range t {
                buf.WriteString("<li>")
                if err = encodeValue(buf, value); err != nil {
                    return
                }
                buf.WriteString("</li>")
            }
            buf.WriteString("</ul>")

        case map[string]interface{}:
            for key, value := range t {
                tag, tag_end, err := encodeTag(key)
                if err != nil {
                    return err
                }
                buf.WriteString(tag)
                if err = encodeValue(buf, value); err != nil {
                    return err
                }

                buf.WriteString(tag_end)
            }
    }
    return
}

// Example:
//  p.class#id
//  p.class1.class2
//  p.class1.class2#id
func encodeTag(s string) (tag, tag_end string, err error) {
    s = strings.Trim(s, " \n\t\r\u00a0")

    if len(s) == 0 {
        err = errors.New("No tag name")
        return
    }
    var tag_name string
    var class    []string
    var id       string

    if i := strings.IndexAny(s, "#"); i > -1 {
        id = s[i+1:]
        s  = s[:i]
    }
    class = strings.Split(s, ".")
    tag_name = class[0]
    class = class[1:]

    tag_end = "</"+tag_name+">"
    tag     = "<"+tag_name

    if len(id) > 0 {
        tag += ` id="`+id+`"`
    }
    if len(class) > 0 {
        tag += ` class="`+strings.Join(class, " ")+`"`
    }
    tag += ">"
    return
}
