package modules

func Crawl(urlQueue *Queue, scrapes [][]string) {
	for _, url := range scrapes {
		urlQueue.Enqueue(url[1])
	}
}
