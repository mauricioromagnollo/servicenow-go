package api

func (a *api) Patch(urlPath string, data []byte) (int, []byte, error) {
	return a.request("PATCH", urlPath, data)
}
