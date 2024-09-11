package api

func (a *api) Post(urlPath string, data []byte) (int, []byte, error) {
	return a.request("POST", urlPath, data)
}
