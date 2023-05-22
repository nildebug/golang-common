package tools

import (
	"math/rand"
	"time"
)

// RandomInt
//
//	@Description: 生成随机数
//	@param min
//	@param max
//	@return int
func RandomInt(min, max int) int {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 计算随机数的范围
	rangeSize := max - min + 1

	// 生成随机数
	randomInt := rand.Intn(rangeSize) + min

	return randomInt
}

// GetLocalIP
//
//	@Description: 获取本地IP地址
//	@return string
//	@return error
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}

	return "", fmt.Errorf("no local IP address found")
}

// GetLocalPort
//
//	@Description: 随机获取一个可用的端口
//	@return int
//	@return error
func GetLocalPort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}

// MD5
//
//	@Description: 获取MD5
//	@param str
//	@return string
func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
