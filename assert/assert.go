package assert

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"bytes"
	"time"
)

type tester interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fail()
	FailNow()
	Failed() bool

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

func isEmpty(expr interface{}) bool {
	if expr == nil {
		return true
	}
	switch v := expr.(type) {
	case bool:
		return !v
	case int:
		return 0 == v
	case int8:
		return 0 == v
	case int16:
		return 0 == v
	case int32:
		return 0 == v
	case int64:
		return 0 == v
	case uint:
		return 0 == v
	case uint8:
		return 0 == v
	case uint16:
		return 0 == v
	case uint32:
		return 0 == v
	case uint64:
		return 0 == v
	case string:
		return len(v) == 0
	case float32:
		return 0 == v
	case float64:
		return 0 == v
	case time.Time:
		return v.IsZero()
	case *time.Time:
		return v.IsZero()
	}
	if isNil(expr) {
		return true
	}
	v := reflect.ValueOf(expr)
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan:
		return 0 == v.Len()
	}
	return false
}

func isNil(expr interface{}) bool {
	if nil == expr {
		return true
	}
	v := reflect.ValueOf(expr)
	k := v.Kind()
	return (k == reflect.Chan ||
		k == reflect.Func ||
		k == reflect.Interface ||
		k == reflect.Map ||
		k == reflect.Ptr ||
		k == reflect.Slice) &&
		v.IsNil()
}

func isEqual(v1, v2 interface{}) bool {
	if reflect.DeepEqual(v1, v2) {
		return true
	}
	vv1 := reflect.ValueOf(v1)
	vv2 := reflect.ValueOf(v2)
	if !vv1.IsValid() || !vv2.IsValid() {
		return false
	}
	if vv1 == vv2 {
		return true
	}
	vv1Type := vv1.Type()
	vv2Type := vv2.Type()
	switch vv1Type.Kind() {
	case reflect.Struct, reflect.Ptr, reflect.Func, reflect.Interface:
		return false
	case reflect.Slice, reflect.Array:
		// vv2.Kind()与vv1的不相同
		if vv2.Kind() != reflect.Slice && vv2.Kind() != reflect.Array {
			if vv2Type.ConvertibleTo(vv1Type) {
				return isEqual(vv1.Interface(), vv2.Convert(vv1Type).Interface())
			}
			return false
		}
		if vv1.Len() != vv2.Len() {
			return false
		}

		for i := 0; i < vv1.Len(); i++ {
			if !isEqual(vv1.Index(i).Interface(), vv2.Index(i).Interface()) {
				return false
			}
		}
		return true // for中所有的值比较都相等，返回true
	case reflect.Map:
		if vv2.Kind() != reflect.Map {
			return false
		}
		if vv1.IsNil() != vv2.IsNil() {
			return false
		}
		if vv1.Len() != vv2.Len() {
			return false
		}
		if vv1.Pointer() == vv2.Pointer() {
			return true
		}
		if vv2Type.Key().Kind() != vv1Type.Key().Kind() {
			return false
		}

		for _, index := range vv1.MapKeys() {
			vv2Index := vv2.MapIndex(index)
			if !vv2Index.IsValid() {
				return false
			}

			if !isEqual(vv1.MapIndex(index).Interface(), vv2Index.Interface()) {
				return false
			}
		}
		return true // for中所有的值比较都相等，返回true
	case reflect.String:
		if vv2.Kind() == reflect.String {
			return vv1.String() == vv2.String()
		}
		if vv2Type.ConvertibleTo(vv1Type) { // 考虑v1是string，v2是[]byte的情况
			return isEqual(vv1.Interface(), vv2.Convert(vv1Type).Interface())
		}

		return false
	}

	if vv1Type.ConvertibleTo(vv2Type) {
		return vv2.Interface() == vv1.Convert(vv2Type).Interface()
	} else if vv2Type.ConvertibleTo(vv1Type) {
		return vv1.Interface() == vv2.Convert(vv1Type).Interface()
	}

	return false
}

func hPanic(fn func()) (h bool, msg interface{}) {
	defer func() {
		if msg = recover(); msg != nil {
			h = true
		}
	}()
	fn()

	return
}

func isContains(container, item interface{}) bool {
	if container == nil { // nil不包含任何东西
		return false
	}
	cv := reflect.ValueOf(container)
	iv := reflect.ValueOf(item)
	if cv.Kind() == reflect.Ptr {
		cv = cv.Elem()
	}
	if iv.Kind() == reflect.Ptr {
		iv = iv.Elem()
	}
	if isEqual(container, item) {
		return true
	}
	switch c := cv.Interface().(type) {
	case string:
		switch i := iv.Interface().(type) {
		case string:
			return strings.Contains(c, i)
		case []byte:
			return strings.Contains(c, string(i))
		case []rune:
			return strings.Contains(c, string(i))
		case byte:
			return bytes.IndexByte([]byte(c), i) != -1
		case rune:
			return bytes.IndexRune([]byte(c), i) != -1
		}
	case []byte:
		switch i := iv.Interface().(type) {
		case string:
			return bytes.Contains(c, []byte(i))
		case []byte:
			return bytes.Contains(c, i)
		case []rune:
			return strings.Contains(string(c), string(i))
		case byte:
			return bytes.IndexByte(c, i) != -1
		case rune:
			return bytes.IndexRune(c, i) != -1
		}
	case []rune:
		switch i := iv.Interface().(type) {
		case string:
			return strings.Contains(string(c), string(i))
		case []byte:
			return strings.Contains(string(c), string(i))
		case []rune:
			return strings.Contains(string(c), string(i))
		case byte:
			return strings.IndexByte(string(c), i) != -1
		case rune:
			return strings.IndexRune(string(c), i) != -1
		}
	}
	if (cv.Kind() == reflect.Slice) || (cv.Kind() == reflect.Array) {
		if !cv.IsValid() || cv.Len() == 0 { // 空的，就不算包含另一个，即使另一个也是空值。
			return false
		}
		if !iv.IsValid() {
			return false
		}
		for i := 0; i < cv.Len(); i++ {
			if isEqual(cv.Index(i).Interface(), iv.Interface()) {
				return true
			}
		}
		if (iv.Kind() != reflect.Slice) || (iv.Len() == 0) {
			return false
		}
		if iv.Len() > cv.Len() {
			return false
		}
		ivIndex := 0
		for i := 0; i < cv.Len(); i++ {
			if isEqual(cv.Index(i).Interface(), iv.Index(ivIndex).Interface()) {
				if (ivIndex == 0) && (i+iv.Len() > cv.Len()) {
					return false
				}
				ivIndex++
				if ivIndex == iv.Len() { // 已经遍历完iv
					return true
				}
			} else if ivIndex > 0 {
				return false
			}
		}
		return false
	}
	if cv.Kind() == reflect.Map {
		if cv.Len() == 0 {
			return false
		}
		if (iv.Kind() != reflect.Map) || (iv.Len() == 0) {
			return false
		}
		if iv.Len() > cv.Len() {
			return false
		}
		for _, key := range iv.MapKeys() {
			cvItem := iv.MapIndex(key)
			if !cvItem.IsValid() { // container中不包含该值。
				return false
			}
			if !isEqual(cvItem.Interface(), iv.MapIndex(key).Interface()) {
				return false
			}
		}
		return true
	}
	return false
}


func getCallerInfo() string {
	for i := 0; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		bename := path.Base(file)
		l := len(bename)
		if l < 8 || (bename[l-8:l] != "_test.go") {
			continue
		}
		funcName := runtime.FuncForPC(pc).Name()
		index := strings.LastIndex(funcName, ".Test")
		if -1 == index {
			continue
		}
		funcName = funcName[index+1:]
		if strings.IndexByte(funcName, '.') > -1 {
			continue
		}

		return funcName + "(" + bename + ":" + strconv.Itoa(line) + ")"
	}

	return "<无法获取调用者信息>"
}

func formatMessage(msg1 []interface{}, msg2 []interface{}) string {
	if len(msg1) == 0 {
		msg1 = msg2
	}

	if len(msg1) == 0 {
		return "<未提供任何错误信息>"
	}

	format := ""
	switch v := msg1[0].(type) {
	case []byte:
		format = string(v)
	case []rune:
		format = string(v)
	case string:
		format = v
	case fmt.Stringer:
		format = v.String()
	default:
		return "<无法正确转换错误提示信息>"
	}

	return fmt.Sprintf(format, msg1[1:]...)
}

func sert(t tester, expr bool, msg1 []interface{}, msg2 []interface{}) {
	if !expr {
		t.Error(formatMessage(msg1, msg2) + "@" + getCallerInfo())
	}
}

func True(t tester, expr bool, args ...interface{}) {
	sert(t, expr, args, []interface{}{"True失败，实际值为[%T:%[1]v]", expr})
}

func False(t tester, expr bool, args ...interface{}) {
	sert(t, !expr, args, []interface{}{"False失败，实际值为[%T:%[1]v]", expr})
}

func Nil(t tester, expr interface{}, args ...interface{}) {
	sert(t, isNil(expr), args, []interface{}{"Nil失败，实际值为[%T:%[1]v]", expr})
}

func NotNil(t tester, expr interface{}, args ...interface{}) {
	sert(t, !isNil(expr), args, []interface{}{"NotNil失败，实际值为[%T:%[1]v]", expr})
}

func Equal(t tester, v1, v2 interface{}, args ...interface{}) {
	sert(t, isEqual(v1, v2), args, []interface{}{"Equal失败，实际值为v1=[%T:%[1]v];v2=[%T:%[2]v]", v1, v2})
}

func NotEqual(t tester, v1, v2 interface{}, args ...interface{}) {
	sert(t, !isEqual(v1, v2), args, []interface{}{"NotEqual失败，实际值为v1=[%T:%[1]v];v2=[%T:%[2]v]", v1, v2})
}

func Empty(t tester, expr interface{}, args ...interface{}) {
	sert(t, isEmpty(expr), args, []interface{}{"Empty失败，实际值为[%T:%[1]v]", expr})
}

func NotEmpty(t tester, expr interface{}, args ...interface{}) {
	sert(t, !isEmpty(expr), args, []interface{}{"NotEmpty失败，实际值为[%T:%[1]v]", expr})
}

func Error(t tester, expr interface{}, args ...interface{}) {
	if isNil(expr) { // 空值，必定没有错误
		sert(t, false, args, []interface{}{"Error失败，实际类型为[%T]", expr})
	} else {
		_, ok := expr.(error)
		sert(t, ok, args, []interface{}{"Error失败，实际类型为[%T]", expr})
	}
}

func ErrorString(t tester, expr interface{}, str string, args ...interface{}) {
	if err, ok := expr.(error); ok {
		index := strings.Index(err.Error(), str)
		sert(t, index >= 0, args, []interface{}{"Error失败，实际类型为[%T]", expr})
	}
}

func ErrorType(t tester, expr interface{}, typ error, args ...interface{}) {
	if _, ok := expr.(error); !ok {
		return
	}

	t1 := reflect.TypeOf(expr)
	t2 := reflect.TypeOf(typ)
	sert(t, t1 == t2, args, []interface{}{"ErrorType失败，v1[%v]为一个错误类型，但与v2[%v]的类型不相同", t1, t2})
}

func NotError(t tester, expr interface{}, args ...interface{}) {
	if isNil(expr) { // 空值必定没有错误
		sert(t, true, args, []interface{}{"NotError失败，实际类型为[%T]", expr})
	} else {
		err, ok := expr.(error)
		sert(t, !ok, args, []interface{}{"NotError失败，错误信息为[%v]", err})
	}
}

func FileExists(t tester, path string, args ...interface{}) {
	_, err := os.Stat(path)

	if err != nil && !os.IsExist(err) {
		sert(t, false, args, []interface{}{"FileExists发生以下错误：%v", err.Error()})
	}
}

func FileNotExists(t tester, path string, args ...interface{}) {
	_, err := os.Stat(path)
	sert(t, os.IsNotExist(err), args, []interface{}{"FileExists发生以下错误：%v", err.Error()})
}

func Panic(t tester, fn func(), args ...interface{}) {
	h, _ := hPanic(fn)
	sert(t, h, args, []interface{}{"并未发生panic"})
}

func PanicString(t tester, fn func(), str string, args ...interface{}) {
	if h, msg := hPanic(fn); h {
		index := strings.Index(fmt.Sprint(msg), str)
		sert(t, index >= 0, args, []interface{}{"并未发生panic"})
	}
}

func PanicType(t tester, fn func(), typ interface{}, args ...interface{}) {
	h, msg := hPanic(fn)
	if !h {
		return
	}

	t1 := reflect.TypeOf(msg)
	t2 := reflect.TypeOf(typ)
	sert(t, t1 == t2, args, []interface{}{"PanicType失败，v1[%v]的类型与v2[%v]的类型不相同", t1, t2})

}

func NotPanic(t tester, fn func(), args ...interface{}) {
	h, msg := hPanic(fn)
	sert(t, !h, args, []interface{}{"发生了panic，其信息为[%v]", msg})
}

func Contains(t tester, container, item interface{}, args ...interface{}) {
	sert(t, isContains(container, item), args,
		[]interface{}{"container:[%v]并未包含item[%v]", container, item})
}

func NotContains(t tester, container, item interface{}, args ...interface{}) {
	sert(t, !isContains(container, item), args,
		[]interface{}{"container:[%v]包含item[%v]", container, item})
}

////////////////////////////////////////////////////////////////////////////////
// 接口部分
type sertion struct {
	t tester
}

func New(t tester) *sertion {
	return &sertion{t: t}
}

func (a *sertion) T() tester {
	return a.t
}

func (a *sertion) True(expr bool, msg ...interface{}) *sertion {
	True(a.t, expr, msg...)
	return a
}

func (a *sertion) False(expr bool, msg ...interface{}) *sertion {
	False(a.t, expr, msg...)
	return a
}

func (a *sertion) Nil(expr interface{}, msg ...interface{}) *sertion {
	Nil(a.t, expr, msg...)
	return a
}

func (a *sertion) NotNil(expr interface{}, msg ...interface{}) *sertion {
	NotNil(a.t, expr, msg...)
	return a
}

func (a *sertion) Equal(v1, v2 interface{}, msg ...interface{}) *sertion {
	Equal(a.t, v1, v2, msg...)
	return a
}

func (a *sertion) NotEqual(v1, v2 interface{}, msg ...interface{}) *sertion {
	NotEqual(a.t, v1, v2, msg...)
	return a
}

func (a *sertion) Empty(expr interface{}, msg ...interface{}) *sertion {
	Empty(a.t, expr, msg...)
	return a
}

func (a *sertion) NotEmpty(expr interface{}, msg ...interface{}) *sertion {
	NotEmpty(a.t, expr, msg...)
	return a
}

func (a *sertion) Error(expr interface{}, msg ...interface{}) *sertion {
	Error(a.t, expr, msg...)
	return a
}

func (a *sertion) ErrorString(expr interface{}, str string, msg ...interface{}) *sertion {
	ErrorString(a.t, expr, str, msg...)
	return a
}

func (a *sertion) ErrorType(expr interface{}, typ error, msg ...interface{}) *sertion {
	ErrorType(a.t, expr, typ, msg...)
	return a
}

func (a *sertion) NotError(expr interface{}, msg ...interface{}) *sertion {
	NotError(a.t, expr, msg...)
	return a
}

func (a *sertion) FileExists(path string, msg ...interface{}) *sertion {
	FileExists(a.t, path, msg...)
	return a
}

func (a *sertion) FileNotExists(path string, msg ...interface{}) *sertion {
	FileNotExists(a.t, path, msg...)
	return a
}

func (a *sertion) Panic(fn func(), msg ...interface{}) *sertion {
	Panic(a.t, fn, msg...)
	return a
}

func (a *sertion) PanicString(fn func(), str string, msg ...interface{}) *sertion {
	PanicString(a.t, fn, str, msg...)
	return a
}

func (a *sertion) PanicType(fn func(), typ interface{}, msg ...interface{}) *sertion {
	PanicType(a.t, fn, typ, msg...)
	return a
}

func (a *sertion) NotPanic(fn func(), msg ...interface{}) *sertion {
	NotPanic(a.t, fn, msg...)
	return a
}

func (a *sertion) Contains(container, item interface{}, msg ...interface{}) *sertion {
	Contains(a.t, container, item, msg...)
	return a
}

func (a *sertion) NotContains(container, item interface{}, msg ...interface{}) *sertion {
	NotContains(a.t, container, item, msg...)
	return a
}