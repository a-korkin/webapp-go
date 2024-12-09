package utils

import "strings"

func splitUri(uri string) []string {
	return strings.Split(strings.ToLower(uri), "/")
}

func GetResourcePath(uri string) string {
	tokens := splitUri(uri)
	if len(tokens) > 1 {
		return tokens[1]
	}
	return ""
}

func GetResourceId(uri string) string {
	tokens := splitUri(uri)
	if len(tokens) > 2 {
		return tokens[2]
	}
	return ""
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
