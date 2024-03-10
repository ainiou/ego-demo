package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gotomicro/ego/core/elog"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultSeparator = ";"
	DefaultConnector = ":"
)

// ParseGorm 解析gorm的结构，获取map
// excludeColumns 指定需要跳过的列
func ParseGorm(obj interface{}, excludeColumns ...string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	excludeMap := map[string]bool{}
	for _, t := range excludeColumns {
		excludeMap[t] = true
	}
	result := map[string]interface{}{}
	for i := 0; i < t.NumField(); i++ {
		fType := t.Field(i)
		tagValue, ok := fType.Tag.Lookup("gorm")
		if !ok {
			continue
		}
		tagMap := ParseStrToMap(tagValue, ";", ":")
		column := tagMap["column"]
		if len(column) <= 0 || excludeMap[column] {
			// 不存在column，过滤
			continue
		}
		result[column] = v.Field(i).Interface()
	}
	return result
}

func ParseStrToMap(source string, separator, connector string) map[string]string {
	if len(separator) <= 0 {
		separator = DefaultSeparator
	}
	if len(connector) <= 0 {
		connector = DefaultConnector
	}
	result := map[string]string{}
	for _, item := range strings.Split(source, separator) {
		if len(item) <= 0 {
			continue
		}

		splits := strings.Split(item, connector)
		key, value := splits[0], ""
		if len(splits) >= 2 {
			value = splits[1]
		}
		result[key] = value
	}
	return result
}

// HTTPGetJSON 发起HTTP GET请求并返回响应内容和错误
func HTTPGetJSON(url string, data interface{}) error {
	// 发起GET请求
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 检查响应状态码
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response：statusCode%d status%s", response.StatusCode, response.Status)
	}

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("json umarshal err:%v", err)
	}

	return nil
}

// md5 加密
func MD5(rawMsg string) string {
	data := []byte(rawMsg)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str1)
}

func ParseToInt64(input string) (int64, error) {
	input = strings.TrimSpace(input)
	userId, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		elog.Warn("ParseToInt32 ParseInt Error", elog.Any("input", input), elog.FieldErr(err))
		// 尝试用科学计数法
		decimalNum, err := decimal.NewFromString(input)
		if err != nil {
			elog.Error("ParseToInt32 decimal.NewFromString error", elog.Any("input", input), elog.FieldErr(err))
			return 0, err
		}
		userId = decimalNum.IntPart()
	}

	return userId, nil
}

// HashPassword 使用bcrypt加密密码
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func SubPhoneName(phone string) string {
	if len(phone) <= 4 {
		return phone
	}
	return "用户" + phone[len(phone)-4:]
}

//
//// CopyTo 根据属性名称拷贝，只做一级拷贝
//func CopyTo(source, dest interface{}) {
//	if source == nil || dest == nil {
//		return
//	}
//	st, sv := reflect.TypeOf(source), reflect.ValueOf(source)
//	dt, dv := reflect.TypeOf(dest), reflect.ValueOf(dest)
//	if dt.Kind() != reflect.Pointer {
//		elog.Panic("CopyTo dest 仅支持指针类型", elog.Any("source", source), elog.Any("dest", dest))
//		return
//	}
//	dt, dv = dt.Elem(), dv.Elem()
//
//	for i := 0; i < dv.NumField(); i++ {
//		field := dt.Field(i)
//		if !field.IsExported() {
//			// 属性不可见时，不进行赋值
//			continue
//		}
//		fName := field.Name         // 当前属性的名称
//		sField, ok := st.FieldByName(fName)
//		if !ok {
//			// 源对象中
//		}
//		df := dv.Field(i)           // 当前目标value
//		sf := sv.FieldByName(fName) // 当前源value
//		if sf.IsValid() && {
//			df.Set(sf) // sf合法，则设置到df中去
//		}
//	}
//	return
//}

// GroupBy 将slice进行分组
func GroupBy[K comparable, V any](slice []V, mapFunc func(V) K) map[K][]V {
	result := map[K][]V{}
	for _, item := range slice {
		k := mapFunc(item)
		result[k] = append(result[k], item)
	}
	return result
}

// Mapping V转换成其他类型
func Mapping[S any, T any](slice []S, mapFunc func(S) T) []T {
	var result []T
	for _, item := range slice {
		t := mapFunc(item)
		result = append(result, t)
	}
	return result
}

func Slice[S any](slice []S, num int) [][]S {
	var result [][]S
	for i := 0; i < len(slice); i += num {
		end := i + num
		if end >= len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

// DurationToZero 到明天0点的时长
func DurationToZero() time.Duration {
	now := time.Now()
	year, month, day := now.AddDate(0, 0, 1).Date() // 明天
	tomorrow := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return tomorrow.Sub(now)
}
