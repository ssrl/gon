package bean

import "launchpad.net/mgo"
import "app/domain/book"

type AppContext map[string]func()interface{}

var registry = make(AppContext)

func Registry() map[string]func()interface{} {
    return registry
}

type Context struct {
    name string
    function func()interface{}
    reply chan int
}

var ch chan *Context = make(chan *Context, 1)
func StartBeanServer() {
    go func(){
        for {
            ctx := <-ch
            registry[ctx.name] = ctx.function
            ctx.reply<- 1
        }
    }()
}

func bean(name string, f func()interface{}) {
    ctx := &Context{name, f, make(chan int, 1)}
    ch<- ctx
    <-ctx.reply
}

func GetBean(name string) interface{} {
    return registry[name]()
}

func Initialize() {
    StartBeanServer()
    bean("Session", func()interface{} {
        session, _ := mgo.Mongo("127.0.0.1")
        session.SetMode(mgo.Monotonic, true)
        return session
    })
    bean("BookService", func()interface{}{
        bookService := &book.BookService{GetBean("Session").(*mgo.Session)}
        return bookService
    })
}