package myhttp

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var (
	defaultHeaders = map[string]string{
		"User-Agent": "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)",
	}
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		},
		Timeout: 10 * time.Second,
	}
)

// ConcurrentRequests 并发发送HTTP请求
func ConcurrentRequests(urls []string, payload, flags string) ([]string, []bool, []error) {
	var (
		wg         sync.WaitGroup
		resultCh   = make(chan string, len(urls))
		containsCh = make(chan bool, len(urls))
		errorCh    = make(chan error, len(urls))
		results    []string
		contains   []bool
		errors     []error
	)

	wg.Add(len(urls))
	for _, url := range urls {
		go func(u string) {
			defer wg.Done()
			result, contain, err := GetRequest(u, payload, flags)
			resultCh <- result
			containsCh <- contain
			errorCh <- err
		}(url)
	}

	go func() {
		wg.Wait()
		close(resultCh)
		close(containsCh)
		close(errorCh)
	}()

	for result := range resultCh {
		results = append(results, result)
	}
	for contain := range containsCh {
		contains = append(contains, contain)
	}
	for err := range errorCh {
		errors = append(errors, err)
	}

	return results, contains, errors
}

// GetRequest 发送HTTP GET请求并检查回显中是否包含特定字符串
func GetRequest(url, payload, flags string) (string, bool, error) {
	urlParams, err := parseURLParams(url, payload)
	if err != nil {
		return "", false, fmt.Errorf("解析URL参数时出错: %v", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", false, fmt.Errorf("创建请求时出错: %v", err)
	}

	setHeaders(req)
	req.URL.RawQuery = buildQueryString(urlParams)
	//fmt.Println(req.URL)

	time.Sleep(time.Millisecond * 3)
	resp, err := client.Do(req)
	if err != nil {
		return "", false, fmt.Errorf("发出请求时出错: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", false, fmt.Errorf("读取响应时出错: %v", err)
	}

	return string(body), strings.Contains(string(body), flags), nil
}

func parseURLParams(rawURL string, payload string) (map[string]string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("解析URL时出错: %v", err)
	}

	urlParams := parsedURL.Query()
	params := make(map[string]string)

	for key := range urlParams {
		params[key] = payload
	}

	return params, nil
}

func setHeaders(req *http.Request) {
	for key, value := range defaultHeaders {
		req.Header.Set(key, value)
	}
}

func buildQueryString(params map[string]string) string {
	var queryStrings []string
	for key, value := range params {
		queryStrings = append(queryStrings, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
	}
	return strings.Join(queryStrings, "&")
}
