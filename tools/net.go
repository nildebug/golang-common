package tools

import (
	"errors"
	"net"
	"time"
)

// IsConnClosed
//
//	@Description: 判断conn连接是否关闭
//	@param conn
//	@return bool
func IsConnClosed(conn net.Conn) bool {
	var tempErr *net.OpError
	if errors.As(conn.SetReadDeadline(time.Time{}), &tempErr) && tempErr.Timeout() {
		return true
	}
	if errors.As(conn.SetWriteDeadline(time.Time{}), &tempErr) && tempErr.Timeout() {
		return true
	}
	if errors.As(conn.SetDeadline(time.Time{}), &tempErr) && tempErr.Timeout() {
		return true
	}
	return false
}
