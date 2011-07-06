package gon

type WebContext interface {
    WriteString(content string)
    GetParams()map[string]string
}
