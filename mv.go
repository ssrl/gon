package mv

type Model map[string]interface{}
type View string
func (v View) String() string {
    return string(v)
}
