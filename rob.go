// ro パッケージは []byte, string 間の低コストな相互キャストを提供します.
//
// ro パッケージから返された値は Read Only に扱う必要があります.
// 書き込みを行った場合 access violation が生じる場合があります.
package rob

import "unsafe"

// string を []byte に変換します.
func Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// []byte を string に変換します.
func String(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}
