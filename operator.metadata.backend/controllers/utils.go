package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}
func RetError(ctx *gin.Context, code int, msg interface{}, data interface{}) error {
	if !isNil(data) {
		ctx.JSON(http.StatusOK, gin.H{"code": code, "message": msg, "data": data})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
	}

	return fmt.Errorf("[ERROR]: %v", gin.H{"code": code, "cause": msg})
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}

func CreatePath(path string) bool {
	isExist, _ := PathExists(path)
	if isExist {
		return true
	}
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return false
	}

	return true
}

func FileToSha256(path string) (string, error) {
	hash := sha256.New()
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("====open erro:%v\n", err)
	}
	defer f.Close()
	if _, err := io.Copy(hash, f); err != nil {
		fmt.Printf("=======copy erro:%v\n", err)
	}
	s := hash.Sum(nil)
	return fmt.Sprintf("%x", s), err
}

func CheckImageCode(src []byte) bool {

	bytes := bytesToHexString(src)
	if strings.Index(bytes, strings.ToLower("FFD8FF")) == 0 {
		/// jpg
	} else if strings.Index(bytes, strings.ToLower("89504E47")) == 0 {
		/// png
	} else if strings.Index(bytes, strings.ToLower("49492A00")) == 0 {
		/// tif
	} else if strings.Index(bytes, strings.ToLower("47494638")) == 0 {
		/// gif
	} else if strings.Index(bytes, strings.ToLower("424DC001")) == 0 {
		/// bmp
	} else if strings.Index(bytes, strings.ToLower("38425053")) == 0 {
		/// psd
	} else if strings.Index(bytes, strings.ToLower("000001BA")) == 0 || strings.Index(bytes, strings.ToLower("000001B3")) == 0 {
		///MPEG
	} else {
		return false
	}
	return true
}

func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	i, length := 100, len(src)
	if length < i {
		i = length
	}
	for j := 0; j < i; j++ {
		sub := src[j] & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return strings.ToLower(res.String())
}
