package utils

import "regexp"

func OnlyNumbers(v string) string {
    re := regexp.MustCompile("[0-9]+")
    matches := re.FindAllString(v, -1)
    var result string
    for _, match := range matches {
        result += match
    }
    return result
}
