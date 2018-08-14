package dumper

import (
	"bytes"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

// MysqlDumper struct
type MysqlDumper struct {
	Logger *zap.Logger
}

// Dump query of MySQL
func (m *MysqlDumper) Dump(in []byte) error {
	if in[4] != 0x03 {
		return nil
	}
	buff := bytes.NewBuffer(in)
	_, _ = buff.Read(make([]byte, 4))
	str, _ := buff.ReadString(0x00)
	query := strings.Trim(str, "\x00")
	m.Logger.Info(fmt.Sprintf("%s", query))
	return nil
}