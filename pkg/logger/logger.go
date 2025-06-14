package logger

import (
	"io"
	"os"
	"strings"

	formatters "github.com/fabienm/go-logrus-formatters"
	graylog "github.com/gemnasium/logrus-graylog-hook"
	"github.com/pmshoot/service-lib/pkg/lib"
	"github.com/sirupsen/logrus"
)

// TODO: добавить трэйсер ошибок
//
//	import "github.com/pkg/errors"
//
//	errors.WithStack(err))

type Logger struct {
	*logrus.Logger
}

const (
	// log levels
	TRACELEVEL = "trace"
	DEBUGLEVEL = "debug"
	INFOLEVEL  = "info"
	WARNLEVEL  = "warn"
	ERRORLEVEL = "error"

	// log format
	TEXTFORMAT = "text"
	JSONFORMAT = "json"
	GELFFORMAT = "gelf"

	//
	GRAYLOGADDRDEFAULT = "localhost:12201"

	// env vars
	GRAYLOGADDR = "HS_GRAYLOG_ADDR"
	LOGFORMAT   = "HS_LOG_FORMAT"
	LOGLEVEL    = "HS_LOG_LEVEL"
	LOGFILE     = "HS_LOG_FILE"
)

func New() (*Logger, error) {
	lgr := logrus.New()
	level := lib.GetEnv(LOGLEVEL, INFOLEVEL)

	switch strings.ToLower(level) {
	case TRACELEVEL:
		lgr.SetLevel(logrus.TraceLevel)
		lgr.SetReportCaller(true)
	case DEBUGLEVEL:
		lgr.SetLevel(logrus.DebugLevel)
	case INFOLEVEL:
		lgr.SetLevel(logrus.InfoLevel)
	case WARNLEVEL:
		lgr.SetLevel(logrus.WarnLevel)
	case ERRORLEVEL:
		lgr.SetLevel(logrus.ErrorLevel)
	default:
		lgr.SetLevel(logrus.InfoLevel)
	}

	formatter := lib.GetEnv(LOGFORMAT, TEXTFORMAT)
	switch strings.ToLower(formatter) {
	case JSONFORMAT:
		lgr.SetFormatter(&logrus.JSONFormatter{})
	case TEXTFORMAT:
		lgr.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	case GELFFORMAT:
		hostname := lib.HostnameOrLocalHost()
		logHostname := lib.GetEnv(GRAYLOGADDR, GRAYLOGADDRDEFAULT)
		glfFormatter := formatters.NewGelf(hostname)
		lgr.SetFormatter(glfFormatter)
		hook := graylog.NewGraylogHook(logHostname, map[string]any{})
		lgr.AddHook(hook)
	default:
		lgr.SetFormatter(&logrus.TextFormatter{})
	}

	lgr.SetOutput(os.Stdout)

	// set to log to file
	logFile := lib.GetEnv(LOGFILE, "")
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			lgr.Warnln("failed log file open, using default stderr")
		}

		multi := io.MultiWriter(os.Stdout, file)
		lgr.SetOutput(multi)
	}

	log := new(Logger)
	log.Logger = lgr
	return log, nil
}
