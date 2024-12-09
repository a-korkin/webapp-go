package utils

import "strings"

func GetResourcePath(uri string) string {
	return strings.Split(strings.ToLower(uri), "/")[1]
}

func GetQueryParams(queryStr string) map[string]string {
	result := make(map[string]string, 0)
	if queryStr != "" {
		for _, str := range strings.Split(queryStr, "&") {
			arr := strings.Split(str, "=")
			result[arr[0]] = arr[1]
		}
	}
	return result
}
