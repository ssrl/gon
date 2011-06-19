package goom

type DataSource struct {
	
}

type QueryTemplate interface {
	Get(id string)interface{}
	FindAll()
}
