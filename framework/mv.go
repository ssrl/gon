package mv

type Model map[string]interface{}
type View string

func (this View) String() string {
    return string(this)
}
