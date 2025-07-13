package robotstxt

import (
	"fmt"
	"io"
	"net/http"
)

func ScanRobotsTxt(url string) {
	resp, err := http.Get(fmt.Sprintf("%s/robots.txt", url))

	if err != nil {
		fmt.Println("robots.txt request failed!", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("robots.txt Not Found")
		return
	}

	body, err := io.ReadAll(resp.Body)

	bodyStr := string(body)
	fmt.Println("robots.txt non-404 response:")
	fmt.Println(bodyStr)

	if resp.Body.Close() != nil {
		fmt.Println("robots.txt body reading closed failed!", err)
	}
}
