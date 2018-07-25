package dumper

import (
	"bytes"
	"fmt"
	"strings"
)

// MysqlDumper struct
type MysqlDumper struct{}

// Dump MySQL query
func (m *MysqlDumper) Dump(in []byte) string {
	if in[4] != 0x03 {
		return ""
	}
	buff := bytes.NewBuffer(in)
	_, _ = buff.Read(make([]byte, 4))
	str, _ := buff.ReadString(0x00)
	query := strings.Trim(str, "\x00")
	return fmt.Sprintf("%s\n", query)
}
