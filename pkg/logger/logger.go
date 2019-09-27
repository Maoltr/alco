package logger

import (
	"github.com/Maoltr/alco/pkg/config"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

type log struct {
	*logrus.Logger
}

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Hook() echo.MiddlewareFunc
}

func New(cfg config.Logger) Logger {
	logger := logrus.New()
	logger.SetFormatter(&Formatter{
		TimestampFormat: cfg.TimestampFormat,
		FieldsOrder:     cfg.FieldsOrder,
		HideKeys:        cfg.HideKeys,
		NoColors:        cfg.NoColors,
		NoFieldsColors:  cfg.NoFieldsColors,
		ShowFullLevel:   cfg.ShowFullLevel,
	})

	if level, err := logrus.ParseLevel(cfg.LogLevel); err == nil {
		logger.SetLevel(level)
	}

	return log{logger}
}

func (l log) logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path
	if p == "" {
		p = "/"
	}

	bytesIn := req.Header.Get(echo.HeaderContentLength)
	if bytesIn == "" {
		bytesIn = "0"
	}

	if res.Status == http.StatusOK {
		l.WithFields(map[string]interface{}{
			"remote_ip":  c.RealIP(),
			"method":     req.Method,
			"path":       p,
			"referer":    req.Referer(),
			"user_agent": req.UserAgent(),
			"status":     res.Status,
			"latency":    strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		}).Info("Handled request")
	}

	if res.Status != http.StatusOK {
		l.WithFields(map[string]interface{}{
			"remote_ip":     c.RealIP(),
			"host":          req.Host,
			"uri":           req.RequestURI,
			"method":        req.Method,
			"path":          p,
			"referer":       req.Referer(),
			"user_agent":    req.UserAgent(),
			"status":        res.Status,
			"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
			"latency_human": stop.Sub(start).String(),
			"bytes_in":      bytesIn,
			"bytes_out":     strconv.FormatInt(res.Size, 10),
		}).Warn("Error request")
	}

	return nil
}

func (l log) logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return l.logrusMiddlewareHandler(c, next)
	}
}

func (l log) Hook() echo.MiddlewareFunc {
	return l.logger
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
