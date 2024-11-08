package api

type Endpoint interface {
	Get() (interface{}, error)
}