package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成或透传 RequestId
		requestID := c.GetHeader("X-Request-Id")
		if requestID == "" {
			requestID = generateRequestID()
		}
		// 将 RequestId 写入响应头，便于链路追踪
		c.Writer.Header().Set("X-Request-Id", requestID)

		// 开始时间
		start := time.Now()
		// 请求路径
		path := c.Request.URL.Path
		// 请求方法
		method := c.Request.Method
		// 查询参数（原样字符串）
		query := c.Request.URL.RawQuery

		// 处理请求
		c.Next()

		// 结束时间
		latency := time.Since(start)
		// 状态码
		statusCode := c.Writer.Status()

		// 结构化日志（JSON 行）
		entry := map[string]interface{}{
			"timestamp":   time.Now().Format(time.RFC3339Nano),
			"method":      method,
			"path":        path,
			"query":       query,
			"ip":          c.ClientIP(),
			"status":      statusCode,
			"latency_ms":  float64(latency.Microseconds()) / 1000.0,
			"request_id":  requestID,
			"user_agent":  c.Request.UserAgent(),
			"content_len": c.Request.ContentLength,
		}
		data, err := json.Marshal(entry)
		if err != nil {
			// 兜底：编码失败时输出简单文本日志到控制台
			log.Printf("[LOG_ENCODING_ERROR] %s %s %s %d %v rid=%s",
				method, path, c.ClientIP(), statusCode, latency, requestID,
			)
			return
		}

		// 写入到 控制台 + 文件（当天日志文件）
		writer, werr := getDailyLogWriter()
		if werr != nil {
			// 兜底：文件不可用时仍然输出控制台
			log.Printf("%s", string(data))
			return
		}
		// 同时写入 stdout 与文件
		multi := io.MultiWriter(os.Stdout, writer)
		_, _ = multi.Write(append(data, '\n'))
	}
}

// ------- 以下为文件写入与工具函数（线程安全，按天切换文件）-------

var (
	logMutex       sync.Mutex
	currentLogDate string
	currentLogFile *os.File
)

// getDailyLogWriter 获取当天日志文件的 io.Writer，自动创建目录与文件，并在日期变化时切换文件
func getDailyLogWriter() (io.Writer, error) {
	logMutex.Lock()
	defer logMutex.Unlock()

	today := time.Now().Format("2006-01-02")
	if currentLogFile != nil && currentLogDate == today {
		return currentLogFile, nil
	}

	// 关闭旧文件
	if currentLogFile != nil {
		_ = currentLogFile.Close()
		currentLogFile = nil
	}

	// 确保目录存在
	logDir := "./logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	// 文件名示例：./logs/2025-11-11.log
	filename := filepath.Join(logDir, today+".log")
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	currentLogDate = today
	currentLogFile = f
	return currentLogFile, nil
}

// generateRequestID 生成随机 RequestId（16 字节十六进制字符串）
func generateRequestID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		// 兜底：若随机读取失败，退化为时间戳
		return hex.EncodeToString([]byte(time.Now().Format("20060102150405.000000000")))
	}
	return hex.EncodeToString(b[:])
}
