import (
	"decadecollab/internal/lib/logger/prettylog"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CurrentTime(c echo.Context) error {
	prikol := slog.HandlerOptions{

		AddSource: true,
	}

	logger := slog.New(prettylog.NewHandler(&prikol))
	
	logger.Info("tam po CurrentTime postuchali")
  	return c.JSON(http.StatusOK, map[string]any{"current_time":time.Now().Format("2006-01-02 15:04:05")})
	
}
