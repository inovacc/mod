package buildcfg

import (
	"os"
	"runtime"
	"strings"
)

var DefaultGO386 = `sse2`
var DefaultGOAMD64 = `v1`
var DefaultGOARM = `7`
var DefaultGOARM64 = `v8.0`
var DefaultGOMIPS = `hardfloat`
var DefaultGOMIPS64 = `hardfloat`
var DefaultGOPPC64 = `power8`
var DefaultGORISCV64 = `rva20u64`
var defaultGOEXPERIMENT = ``
var defaultGO_EXTLINK_ENABLED = ``
var defaultGO_LDSO = ``
var version = `go1.24.1`
var defaultGOOS = runtime.GOOS
var defaultGOARCH = runtime.GOARCH
var DefaultGOFIPS140 = `off`
var DefaultGOEXPERIMENT = quote(os.Getenv("CGO_ENABLED"))

func quote(s string) string {
	const hex = "0123456789abcdef"
	var out strings.Builder
	out.WriteByte('"')
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 0x20 <= c && c <= 0x7E && c != '"' && c != '\\' {
			out.WriteByte(c)
		} else {
			out.WriteByte('\\')
			out.WriteByte('x')
			out.WriteByte(hex[c>>4])
			out.WriteByte(hex[c&0xf])
		}
	}
	out.WriteByte('"')
	return out.String()
}
