package php

import "bytes"

// ChunkSplit - Split a string into smaller chunks
func ChunkSplit(str string, chunkLen int, end string) string {
	if chunkLen <= 0 {
		return str
	}
	var buf bytes.Buffer
	l := len(str)
	for i := 0; i < l; i += chunkLen {
		if chunkLen > l-i {
			chunkLen = l - i
		}
		buf.WriteString(str[i : i+chunkLen])
		buf.WriteString(end)
	}
	return buf.String()
}
