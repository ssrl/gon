package bean

type AppContext interface {
    GetBean(name string) interface{}
}
