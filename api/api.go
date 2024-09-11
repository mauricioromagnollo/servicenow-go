package api

// API is the interface that wraps the basic methods to interact with the ServiceNow API.
type API interface {
	Get(urlPath string) (int, []byte, error)
	Post(urlPath string, data []byte) (int, []byte, error)
	Put(urlPath string, data []byte) (int, []byte, error)
	Delete(urlPath string) (int, []byte, error)
	Patch(urlPath string, data []byte) (int, []byte, error)
}
