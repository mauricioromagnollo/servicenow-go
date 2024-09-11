package api

func (a *api) Delete(urlPath string) (int, []byte, error) {
	return a.request("DELETE", urlPath, nil)
}
