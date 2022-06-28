package test

import (
	"github.com/dxvgef/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestPosts(t *testing.T) {
	engine := gin.Default()
	var a = 0
	engine.POST("/speed", func(context *gin.Context) {
		context.JSON(http.StatusOK, a)
	})
	engine.PUT("/speed", func(context *gin.Context) {
		a1 := context.DefaultQuery("a1", "1000")
		parseInt, _ := strconv.ParseInt(a1, 10, 64)
		a = int(parseInt)
		context.JSON(http.StatusOK, a)
	})
	engine.Run()
}

func TestCopyFile(t *testing.T) {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		// 传输demo.mp4文件，限速每秒100KKB
		if err := limiter.ServeFile(resp, req, "./demo.mp4", 100*1024); err != nil {
			resp.WriteHeader(500)
			resp.Write([]byte(err.Error()))
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err.Error())
		return
	}
}

// 服务端文件传输限速
func ServeFile(resp http.ResponseWriter, req *http.Request, filePath string, speed float64) error {
	// 当前连接的限速500KB/s
	speedLimiter := NewSpeedLimiter(speed)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	http.ServeContent(
		resp,
		req,
		filePath,
		fileInfo.ModTime(),
		NewReadSeeker(file, speedLimiter),
	)
	return nil
}

type Reader struct {
	reader  io.Reader
	limiter *rate.Limiter
}

func (reader *Reader) Read(buf []byte) (int, error) {
	n, err := reader.reader.Read(buf)
	if n <= 0 {
		return n, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = reader.limiter.WaitN(ctx, n)
	return n, err
}

type ReadSeeker struct {
	io.ReadSeeker
	reader io.Reader
}

func (rs ReadSeeker) Read(p []byte) (int, error) {
	return rs.reader.Read(p)
}

func NewSpeedLimiter(speed float64) *rate.Limiter {
	return rate.NewLimiter(rate.Limit(speed), int(speed))
}

func NewReader(reader io.Reader, limiter *rate.Limiter) io.Reader {
	return &Reader{
		reader:  reader,
		limiter: limiter,
	}
}

func NewReadSeeker(readSeeker io.ReadSeeker, limiter *rate.Limiter) io.ReadSeeker {
	return ReadSeeker{
		readSeeker,
		NewReader(readSeeker, limiter),
	}
}
