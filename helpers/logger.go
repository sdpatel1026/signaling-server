package helpers

import (
	"fmt"
	"log"
	"signaling-server/configs"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

const (
	KEY_LOGFILE_PATH        string = "LOGFILE_PATH"
	KEY_LOGFILE_ENCODING    string = "LOGFILE_ENCODING"
	KEY_LOGFILE_MESSAGE_KEY string = "LOGFILE_MESSAGE_KEY"
	KEY_LOGFILE_TIME_KEY    string = "LOGFILE_TIME_KEY"
	KEY_LOGFILE_LEVEL_KEY   string = "LOGFILE_LEVEL"
	KEY_LOGFILE_CALLER_KEY  string = "LOGFILE_CALLER_KEY"
	DEFAULT_LOGFILE_NAME    string = "logfile.log"
)

func init() {
	var err error
	Logger, err = GetStructuredLogger("")
	if err != nil {
		log.Fatalf("error in creating default logger:%s", err.Error())
	}
}
func GetStructuredLogger(fileName string) (*zap.SugaredLogger, error) {

	if strings.TrimSpace(fileName) == "" {
		fileName = DEFAULT_LOGFILE_NAME
	}

	logfileLocation := fmt.Sprintf("%s/%s", configs.GetEnvWithKey(KEY_LOGFILE_PATH, "."), fileName)
	var cfg zap.Config
	cfg.OutputPaths = []string{logfileLocation}
	cfg.Encoding = configs.GetEnvWithKey(KEY_LOGFILE_ENCODING, "json")
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.EncoderConfig = zapcore.EncoderConfig{MessageKey: configs.GetEnvWithKey(KEY_LOGFILE_MESSAGE_KEY, "message"),
		TimeKey:      configs.GetEnvWithKey(KEY_LOGFILE_TIME_KEY, "time"),
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		LevelKey:     configs.GetEnvWithKey(KEY_LOGFILE_LEVEL_KEY, "level"),
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		CallerKey:    configs.GetEnvWithKey(KEY_LOGFILE_CALLER_KEY, "callerKey"),
		EncodeCaller: zapcore.ShortCallerEncoder}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	sugaredLogger := logger.Sugar()
	defer sugaredLogger.Sync()
	return sugaredLogger, nil

}
