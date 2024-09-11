package api

func (a *api) Get(urlPath string) (int, []byte, error) {
	return a.request("GET", urlPath, nil)
}
