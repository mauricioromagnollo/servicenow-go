package api

func (a *api) Put(urlPath string, data []byte) (int, []byte, error) {
	return a.request("PUT", urlPath, data)
}
