package model

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// Auth model for
type Auth struct {
	gorm.Model
	Token      string
	Permission int32
}

func md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GenToken generate a random token
func (auth *Auth) GenToken() string {
	i := rand.Intn(777)
	t := time.Now().UnixNano()

	str := strconv.Itoa(i) + "-" + strconv.FormatInt(t, 10)

	return md5Encode(str)
}

// IsValidPassord literally
func (auth *Auth) IsValidPassord(pw string) bool {
	return md5Encode(pw) == "a14068983c35ab6018046087da23fd4c"
}
