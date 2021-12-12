package utils

import (
	"strings"
)

func getSearchKeyArray(key string) []string {
	words := strings.Split(key, " ")

	// length := len(key)
	// for i := 0; i < length-3; i++ {
	// 	words = append(words, strings.Trim(key[i:i+3], " "))
	// }
	return words
}

// func removeSign(province string) string {
// 	var Regexp_A = `à|á|ạ|ã|ả|ă|ắ|ằ|ẳ|ẵ|ặ|â|ấ|ầ|ẩ|ẫ|ậ`
// 	var Regexp_E = `è|ẻ|ẽ|é|ẹ|ê|ề|ể|ễ|ế|ệ`
// 	var Regexp_I = `ì|ỉ|ĩ|í|ị`
// 	var Regexp_U = `ù|ủ|ũ|ú|ụ|ư|ừ|ử|ữ|ứ|ự`
// 	var Regexp_Y = `ỳ|ỷ|ỹ|ý|ỵ`
// 	var Regexp_O = `ò|ỏ|õ|ó|ọ|ô|ồ|ổ|ỗ|ố|ộ|ơ|ờ|ở|ỡ|ớ|ợ`
// 	var Regexp_D = `Đ|đ`
// 	reg_a := regexp.MustCompile(Regexp_A)
// 	reg_e := regexp.MustCompile(Regexp_E)
// 	reg_i := regexp.MustCompile(Regexp_I)
// 	reg_o := regexp.MustCompile(Regexp_O)
// 	reg_u := regexp.MustCompile(Regexp_U)
// 	reg_y := regexp.MustCompile(Regexp_Y)
// 	reg_d := regexp.MustCompile(Regexp_D)
// 	province = reg_a.ReplaceAllLiteralString(province, "a")
// 	province = reg_e.ReplaceAllLiteralString(province, "e")
// 	province = reg_i.ReplaceAllLiteralString(province, "i")
// 	province = reg_o.ReplaceAllLiteralString(province, "o")
// 	province = reg_u.ReplaceAllLiteralString(province, "u")
// 	province = reg_y.ReplaceAllLiteralString(province, "y")
// 	province = reg_d.ReplaceAllLiteralString(province, "d")

// 	// regexp remove charaters in ()
// 	var RegexpPara = `\(.*\)`
// 	reg_para := regexp.MustCompile(RegexpPara)
// 	province = reg_para.ReplaceAllLiteralString(province, "")

// 	return province
// }
