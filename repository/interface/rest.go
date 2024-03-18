package _interface

type Rest interface {
	PostRequest(endpoint string, requestData any) (string, error)
}
