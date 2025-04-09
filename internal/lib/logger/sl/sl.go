package sl

import (
	"decadecollab/internal/lib/logger/prettylog"
	"log/slog"
	//"errors"
)

func Debug(msg string) {
	logger := slog.New(prettylog.NewHandler(nil))

	logger.Debug(msg)
}


func Info(msg string) {
	logger := slog.New(prettylog.NewHandler(nil))

	logger.Info(msg)
}

func Err(err error){
	logger := slog.New(prettylog.NewHandler(nil))

	logger.Error("We have some problems", "error", err)
}

func ErrWithOpts(err error){     //логгирование ошибок с атрибутами
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr{  //здесь происходит перезапись стандартных аттрибутов нашими
			if a.Key == "nothing"{								//если ничего нет, то и перезаписывать ничего не надо	
				return slog.Attr{}
			}
			return a
		},
	}

	logger := slog.New(prettylog.NewHandler(opts)) //вместо nil указываются аттрибуты

	logger.Error("We have some problems", "error", err)
}