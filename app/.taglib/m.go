package taglib

type Namespace string

type MyTagLib struct {
    Namespace "m"
    AppContext
    Writer
}

//
// <m:test/>
//
func (m *MyTagLib) Test(b Body, attr Attributes) {
    m.Write("this is my tag lib.")
    b.call()
    for k,v := range attr {

    }    
}
