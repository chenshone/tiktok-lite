package util

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/golang-jwt/jwt/v4"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"path/filepath"
	"time"
)

// MyClaims 自定义的payload
type MyClaims struct {
	jwt.RegisteredClaims        // 内置的payload字段
	UserID               string `json:"user_id"` // 自定义的payload字段
}

type JWT struct {
}

var key = []byte("qwouidfoqiwr23fioqw") // 用于签名的密钥

func (j *JWT) GenerateToken(msg string, ttl int64) string {
	now := time.Now()
	// 创建一个新的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		UserID: msg,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(24*ttl) * time.Hour)), // 过期时间，ttl天后过期
			IssuedAt:  jwt.NewNumericDate(now),                                        // 签发时间
			NotBefore: jwt.NewNumericDate(now),                                        // 生效时间
			Issuer:    "tiktok-lite",                                                  // 签发人
		},
	})
	tokenString, _ := token.SignedString(key) // 签名, 返回token字符串
	return tokenString
}

func (j *JWT) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		log.Println(err)
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		log.Println("invalid token")
		return "", errors.New("invalid token")
	}
	return claims.UserID, nil
}

// Mkdir 定义一个创建文件目录的方法
func Mkdir(basePath string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006/01/02/")
	folderPath := filepath.Join(basePath, folderName)
	//使用mkdirall会创建多层级目录
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return ""
	}
	return folderPath + "/"
}

func CheckExt(filename string) bool {
	ext := filepath.Ext(filename)
	if ext == ".mp4" {
		return true
	}
	return false
}

func GetVideoCover(videoPath, snapshotPath string, frameNum int) error {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}

	err = imaging.Save(img, snapshotPath)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}

	return nil
}
