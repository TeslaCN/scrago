package util

import (
	"regexp"
	"strconv"
)

/*
function reurl(code) {
    var decoded = '';
    var c = csplit(code, 8).split('\r\n');
    for (var i = 0; i < (c.length - 1); i++) {
        decoded = decoded + String.fromCharCode((parseInt(c[i], 2) - 10).toString(10))
    }
    return decoded
}

function csplit(a, b, c) {
    b = parseInt(b, 10) || 76;
    c = c || '\r\n';
    if (b < 1) {
        return false
    }
    return a.match(new RegExp(".{0," + b + "}", "g")).join(c)
}
*/
func Decode(encoded string) string {
	length := len(encoded)
	decoded := ""
	for from := 0; from < length; from += 8 {
		to := from + 8
		if to > length {
			to = length
		}
		c, _ := strconv.ParseInt(encoded[from:to], 2, 8)
		decoded += string(c - 10)
	}
	return decoded
}

var jsRegex = regexp.MustCompile("\\$\\('#\\w+'\\)\\.attr\\('href','(magnet:\\?xt=urn:btih:)'\\+reurl\\('(\\w+)'\\)\\);?")

// $('#rid_70869938699').attr('href','magnet:?xt=urn:btih:'+reurl('0101101001101111'));
func ParseJs(js string) string {
	matches := jsRegex.FindStringSubmatch(js)
	if len(matches) < 3 {
		return ""
	}
	return matches[1] + Decode(matches[2])
}
