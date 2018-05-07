package helper

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// GetRandomString generates random string
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// HashString hashes provided string
func HashString(source string) string {
	target := md5.New()
	io.WriteString(target, source)
	return fmt.Sprintf("%x", target.Sum(nil))
}
