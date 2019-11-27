package utils

import (
	"strconv"
	"strings"
)

func CountAdd(lines []string) map[string]int {
	var domainAndCountMap = map[string]int{}
	for _, line := range lines {
		domain, count := Split(line)
		domainAndCountMap[domain] += count
	}
	return domainAndCountMap
}

func Split(str string) (string, int) {
	domainAndCount := strings.Split(str, ", ")
	num, _ := strconv.Atoi(domainAndCount[1])
	return domainAndCount[0], num
}
