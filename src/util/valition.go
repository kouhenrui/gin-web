package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) UInt64() (uint64, error) {
	v, err := strconv.Atoi(s.String())
	return uint64(v), err
}

func (s StrTo) MustUInt32() uint32 {
	curInt32, _ := strconv.Atoi(s.String())
	if curInt32 < 0 {
		return 0
	}
	v, _ := s.UInt32()
	return v
}

func (s StrTo) MustUInt64() uint64 {
	//fmt.Println(s)
	curInt64, _ := strconv.ParseInt(s.String(), 10, 64)
	if curInt64 < 0 {
		return 0
	}
	v, _ := s.UInt64()
	//fmt.Println(v)

	return v
}
func ValidateToStruct(c *gin.Context, a *struct{}) error {
	if err := c.ShouldBindJSON(&a); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		return errs
	}
	return nil
}

//func String(translate map[string][]string) string {
//	return translate
//}
