package log

import (
	"io/ioutil"
	stlog "log" //将标注库的log引入并起一个别名
	"net/http"
	"os"
)

//日志服务的逻辑：每次程序发生行为都需要打印出一条日志内容

//定义一个日志对象，这里使用的是标准库里的log对象，它本身实现了很多方法，这里直接用即可
//它也是一个全局变量
var log *stlog.Logger

//定义一个全局的路径变量,本质上它是一个string，这里为了以后更明显的使用,后续作为string需要显式转换
type logPath string

//为变量实现一个写方法
func (lp logPath) Write(data []byte) (int, error) {

	//若没有文件就创建，并且只写或者附加
	file, err := os.OpenFile(string(lp), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		//如果打开失败，返回错误
		return 0, err
	}
	//凡是打开数据库/文件操作，切记使用defer关闭
	defer file.Close()

	//写入数据
	return file.Write(data)
}

//运行日志（实际上是实例化了一条日志）

func Run(destination string) {
	//将文件写入到路径，并且添加上前缀和长日期和长时间
	log = stlog.New(logPath(destination), "[go] ", stlog.LstdFlags)
}

//新建一个Handler来处理POST请求

func RegisterHandler() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		//做个判断，如果是POST请求，读取内容
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			//判错/判空
			if err != nil || len(msg) == 0 {
				//返回状态码
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			//调用写函数将消息写入到指定路径
			write(string(msg))
		default:
			//不是POST请求则返回方法不允许状态码
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
