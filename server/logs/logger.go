package logs

import (
	"fmt"
	"github.com/892294101/jxutils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 自义定日志结构
type myFormatter struct {
	AppID string
}

func (s *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05.00")
	var reason interface{}
	if v, ok := entry.Data["err"]; ok {
		reason = v
	} else {
		reason = nil
	}
	var msg string
	fName := entry.Caller.Function[strings.LastIndex(entry.Caller.Function, ".")+1:]
	if reason == nil {
		msg = fmt.Sprintf("%s %s %s %v [M] [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String())[:1], fName, entry.Caller.Line, s.groupid, entry.Message)
	} else {
		msg = fmt.Sprintf("%s %s %s %d [M] [%s] %s %v\n", timestamp, strings.ToUpper(entry.Level.String())[:1], fName, entry.Caller.Line, s.groupid, entry.Message, reason)
	}
	return []byte(msg), nil
}

func (s *myFormatter) SetGroupId(groupid string) {
	s.AppID = groupid
}

func InitDDSlog(groupid string) (*logrus.Logger, error) {
	ddslog := logrus.New()

	dir, err := jxutils.GetProgramHome()
	if err != nil {
		return nil, err
	}

	log := filepath.Join(dir, "logs", "jxlife.log")
	logfile, err := os.OpenFile(log, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	//writers := []io.Writer{logfile, os.Stdout}
	writers := []io.Writer{logfile}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Log file open failed: %s", err)
		os.Exit(1)
	} else {
		ddslog.SetOutput(fileAndStdoutWriter)
	}
	ddslog.SetLevel(logrus.DebugLevel)
	ddslog.SetReportCaller(true)

	delog := new(myFormatter)
	delog.SetGroupId(groupid)
	ddslog.SetFormatter(delog)
	//ddslog.Infof("Initialize log file: %s", log)
	return ddslog, nil
}

var AppLog *logrus.Logger
var WebLog *logrus.Logger

func Setup() {
	initAppLog()
	initWebLog()
}

func initAppLog() {
	logFileName := "app.log"
	AppLog = initLog(logFileName)
}

func initWebLog() {
	logFileName := "web.log"
	WebLog = initLog(logFileName)
}

func initLog(logFileName string) *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logPath := "logs/"
	logName := logPath + logFileName
	var f *os.File
	var err error
	//判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(logName); os.IsNotExist(err) {
		f, err = os.Create(logName)
	} else {
		f, err = os.OpenFile(logName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	if err != nil {
		fmt.Println("open log file failed")
	}

	log.Out = f
	log.Level = logrus.InfoLevel
	return log
}

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		WebLog.WithFields(logrus.Fields{
			"http_status": statusCode,
			"total_time":  latencyTime,
			"ip":          clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}).Info("access")
	}
}
