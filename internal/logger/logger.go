package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

// ANSIエスケープシーケンス（glgカラーの再現）
const (
	colorReset = "\033[0m"
	colorInfo  = "\033[36m" // シアン (glgのINFO)
	colorWarn  = "\033[33m" // イエロー (glgのWARN)
	colorError = "\033[31m" // レッド (glgのERR)
	colorDebug = "\033[35m" // マゼンタ (glg(DEBG))
)

type StyleHandler struct {
	w io.Writer
}

func New() *slog.Logger {
	handler := &StyleHandler{w: os.Stdout}
	return slog.New(handler)
}

func (h *StyleHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *StyleHandler) Handle(_ context.Context, r slog.Record) error {
	var levelColor, levelStr string

	switch r.Level {
	case slog.LevelDebug:
		levelColor = colorDebug
		levelStr = "[DEBG]"
	case slog.LevelInfo:
		levelColor = colorInfo
		levelStr = "[INFO]"
	case slog.LevelWarn:
		levelColor = colorWarn
		levelStr = "[WARN]"
	case slog.LevelError:
		levelColor = colorError
		levelStr = "[ERR ]"
	default:
		levelColor = colorReset
		levelStr = fmt.Sprintf("[%s]", r.Level.String())
	}

	timeStr := r.Time.Format("2006/01/02 15:04:05")

	fmt.Fprintf(h.w, "%s %s%s%s %s", levelColor, timeStr, levelStr, r.Message, colorReset)

	r.Attrs(func(a slog.Attr) bool {
		fmt.Fprintf(h.w, " %s=%v", a.Key, a.Value.Any())
		return true
	})

	fmt.Fprintln(h.w) // 改行
	return nil
}

func (h *StyleHandler) WithAttrs(_ []slog.Attr) slog.Handler { return h }
func (h *StyleHandler) WithGroup(_ string) slog.Handler      { return h }
