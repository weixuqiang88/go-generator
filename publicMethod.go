package easy

import (
	"bytes"
	"regexp"
	"fmt"
	"os"
	"strconv"
	"time"
)

func CheckErr(e error) {
	if e != nil {
		panic("Error:" + e.Error())
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

func GetPathSeperator() string{
	if os.IsPathSeparator('\\'){
		return "\\"
	}else{
		return "/"
	}
}

func CreateDirectory(dirname string)(e error){
	fi,e:=os.Stat(dirname)
	if os.IsNotExist(e){
		e=os.MkdirAll(dirname,os.ModePerm)
	}else if !fi.IsDir(){
		e=os.MkdirAll(dirname,os.ModePerm)
	}
	return
}

func ParseIntToSecond(x int) (time.Duration, error) {
    fc := 0.001
    return time.ParseDuration(strconv.FormatFloat(float64(x)-fc, 'f', 3, 64) + "s")
}

func PrintArrayString(arr []string) {
	for _, a := range arr {
		fmt.Println(a)
	}
}

func GetNowString() string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func Iterator(str_arr []string) func() string {
	var rs_p string
	i := 0
	return func() string {
		for {
			if i < len(str_arr) {
				rs_p = str_arr[i]
				i += 1
				return rs_p
			} else {
				i = 1
				return str_arr[0]
			}
		}
	}
}

func TrimStringArray(arr []string) (rarr []string) {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "" {
			continue
		} else {
			rarr = arr[:i+1]
			break
		}
	}
	return
}

func IsDigitalString(s string) bool{
	p:=`\d+`
	reg:=regexp.MustCompile(p)
	return reg.MatchString(s)
}

func BufferAddString(sl ...string) string{
	var buf bytes.Buffer
	for _,s:=range sl{
		buf.WriteString(s)
	}
	return buf.String()
}