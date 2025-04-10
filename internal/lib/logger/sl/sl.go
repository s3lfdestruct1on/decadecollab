package sl

import (
	"decadecollab/internal/lib/logger/prettylog"
	"log/slog"
	//"errors"
)
//функции *Opts принимают не только строку/ошибку, но и атрибуты

func Debug(msg string) {
	logger := slog.New(prettylog.NewHandler(nil))

	logger.Debug(msg)
}


func DebugOpts(msg string, opts *slog.HandlerOptions) {
	logger := slog.New(prettylog.NewHandler(opts))

	logger.Debug(msg)
}


func Info(msg string) {
	logger := slog.New(prettylog.NewHandler(nil))

	logger.Info(msg)
}

func InfoOpts(msg string, opts *slog.HandlerOptions) {
	logger := slog.New(prettylog.NewHandler(opts))

	logger.Info(msg)
}

func Err(err error){
	logger := slog.New(prettylog.NewHandler(nil))

	logger.Error("We have some problems", "error", err)
}

func ErrOpts(err error, opts *slog.HandlerOptions){
	logger := slog.New(prettylog.NewHandler(opts))

	logger.Error("We have some problems", "error", err)
}
//эта функция тест обработки ошибки с атрибутами
//на ReplaceAttr происходит перезапись стандартных аттрибутов нашими
//если a.Key пустой, то не делаем ничего 
/*
func TestErrOpts(err error){     
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr{  
			if a.Key == "nothing"{								
				return slog.Attr{}
			}
			return a
		},
	}

	logger := slog.New(prettylog.NewHandler(opts)) //вместо nil указываются аттрибуты

	logger.Error("We have some problems", "error", err)
*/