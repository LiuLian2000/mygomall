package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

var allowedDrift int64 = 60

// key暂时就放这了
var key = "MYGOMALL"

// GenerateHMAC 根据密钥和时间戳生成 HMAC

func GenerateHMAC(message string, timestamp int64) string {
	data := fmt.Sprintf("%d%s", timestamp, message)
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// 验证 HMAC
func VerifyHMAC(message string, timestamp int64, receivedHMAC string) bool {
	currentTimestamp := time.Now().Unix()

	// 检查时间戳,放重放攻击
	if abs(currentTimestamp-timestamp) > allowedDrift {
		return false
	}

	expectedHMAC := GenerateHMAC(message, timestamp)

	return hmac.Equal([]byte(expectedHMAC), []byte(receivedHMAC))
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
