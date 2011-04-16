package taglib

const Namespace = "m"

//
// <m:test/>
//
func Test(b Body, a map[string]interface{}) {
    fmt.Printf("this is hello world from Taglib")
}
