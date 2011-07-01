package goom

type DataStore interface {
    DB(name string) Database
}

type Database interface {
    C(name string) DataSet
}

type Params map[string]interface{}

type Getter interface {
    Get(id string)     interface{}
}
type Inserter interface {
    Insert(p interface{})  (interface{}, bool)
}

type GoomTemplate interface {
    Query(stmt string) interface{}
    FindOne(p Params)  interface{}
    Get(id string)     interface{}
    Insert(p interface{})  (interface{}, bool)
}

type DataSet interface {
    Query(stmt string) interface{}
    FindOne(p Params)  interface{}
    Get(id string)     interface{}
    Insert(p interface{})  (interface{}, bool)
}
