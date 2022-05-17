package browser

type Browser struct {
	Name      string
	LocalPath string
}

func (b *Browser) GetPath() string {
	return b.LocalPath
}

func (b *Browser) GetName() string {
	return b.Name
}

func ChromeExe() (path string) {
	chrome := Browser{Name: "Chrome", LocalPath: "D:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"}
	path = chrome.GetPath()
	return
}

var Chrome = ChromeExe()
