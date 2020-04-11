package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	_ "reflect"
	"runtime"
	"strings"
)

// 自定义结构体
type WLog struct {
}

/**
为MyStruct添加Debug方法
*/
func (m *WLog) Debug(args ...interface{}) {
	Debugf("", args...)
}

func (m *WLog) Debugf(formating string, args ...interface{}) {
	Debugf(formating, args...)
}

func (m *WLog) Error(args ...interface{}) {
	Errorf("", args...)
}

func (m *WLog) Errorf(formating string, args ...interface{}) {
	Errorf(formating, args...)
}

/**
封装日志方法
*/
func Debug(args ...interface{}) {
	LOG("DEBUG", "%+v", args...)
}
func Debugf(formating string, args ...interface{}) {
	LOG("DEBUG", formating, args...)
}

func Error(args ...interface{}) {
	LOG("ERROR", "%+v", args...)
}

func Errorf(formating string, args ...interface{}) {
	LOG("ERROR", formating, args...)
}

/**
获取文件名,函数,行号
*/
func LOG(level string, formating string, args ...interface{}) {
	filename, line, funcname := "???", 0, "???"
	// 获取调用的名称与行号
	pc, filename, line, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo

		filename = filepath.Base(filename) // /full/path/basename.go => basename.go
	}

	log.Printf("%s:%d:%s: %s: %s\n", filename, line, funcname, level, fmt.Sprintf(formating, args...))
}

/**
对象转为格式化json
*/
func ToJson(v interface{}) string {
	var result = ""
	//  将字符串格式化
	buf, _ := json.Marshal(v)
	var jsonBuf bytes.Buffer
	_ = json.Indent(&jsonBuf, buf, "", "\t")
	result = jsonBuf.String()

	//fmt.Printf("haha:%s \n", result)
	//fmt.Println("formated:", result)

	return result
}

/**
 * IO流toJSON
 * Context.Request.Body
 */
func ReaderToJSON(reader *io.ReadCloser) string {
	var result = ""
	var buf []byte
	buf, err := ioutil.ReadAll(*reader) // 读取Request.Body的数据流
	if err == nil {
		*reader = ioutil.NopCloser(bytes.NewBuffer(buf)) // 返回数据流到Request.Body中
		result = string(buf)

		///*  将字符串格式化
		var jsonBuf bytes.Buffer
		_ = json.Indent(&jsonBuf, []byte(result), "", "\t")
		result = jsonBuf.String()
		//*/

		//result = ToJson(result)
	} else {
		Error(err)
	}
	return result
}

/**
获取当前程序路径
*/
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func init() {
	/*
		// 定义一个文件
		fileName := "wlog.log"
		logFile, err := os.Create(fileName)

		defer logFile.Close()

		if err != nil {
			log.Fatalln("open file error !")
		}
		// 创建一个日志对象
		debugLog := log.New(logFile, "[Debug]", log.LstdFlags)

		debugLog.Println("A debug message here")
		//配置一个日志格式的前缀
		debugLog.SetPrefix("[Info]")
		debugLog.Println("A Info Message here ")
		//配置log的Flag参数
		debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
		debugLog.Println("A different prefix")
	*/
}
