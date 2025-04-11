package prettylog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"sync"
)
// коды для назначения цвета
// цветов создано больше чем использовано, на случай если понадобится их поменять
const (
	reset = "\033[0m"

	black        = 30
	red          = 31
	green        = 32
	yellow       = 33
	blue         = 34
	magenta      = 35
	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightGreen   = 92
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	lightCyan    = 96
	white        = 97
)


// функция окрашивания строки, использована вместо пакета color 
func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}


// стуктура в которую вложены:
// slog.handler для отработки логики ручки 
// ссылка на bytes.bufer для отлова вывода из ручки
// mutex для избежания проблем с буфером
type Handler struct{
	h  slog.Handler
	b *bytes.Buffer
	m *sync.Mutex
}


func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}
//функция отвечает за правильную обработку поступивших атрибутов
//на h.m.Lock() происходит блокировка байтого буфера
//на if err := h.h.Handle обработка ошибки при обращении к внутренней ручке
//на var attrs map[string]any задаётся строковая мапа, в которую позже запишут обработанные атрибуты
//на json.Unmarshal(h.b.Bytes(), &attrs) анмаршал атрибутов из буфера в вышезаданную мапу
//после обработка ошибки анмаршала результатов работы внутренней ручки
func (h *Handler) computeAttrs( 
	ctx context.Context,
	r slog.Record,
) (map[string]any, error) {
	h.m.Lock() 					
	defer func() {
		h.b.Reset()				
		h.m.Unlock()
	}()
	if err := h.h.Handle(ctx, r); err != nil { 
		return nil, fmt.Errorf("error when calling inner handler's Handle: %w", err)
	}

	var attrs map[string]any  	
	err := json.Unmarshal(h.b.Bytes(), &attrs) 
	if err != nil { 						   
		return nil, fmt.Errorf("error when unmarshaling inner handler's Handle result: %w", err)
	}
	return attrs, nil
}

//в каком формате будет выводится время в логах
const timeFormat = "[15:04:05.000]" 

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {

	level := r.Level.String() + ":"
	//окрашивание индикации уровня логов через switch
	switch r.Level {
	case slog.LevelDebug:
		level = colorize(darkGray, level)
	case slog.LevelInfo:
		level = colorize(cyan, level)
	case slog.LevelWarn:
		level = colorize(lightYellow, level)
	case slog.LevelError:
		level = colorize(lightRed, level)
	}
	


	attrs, err := h.computeAttrs(ctx, r)
	if err != nil {
		return err
	}
	//здесь обрабатываются полученные атрибуты и оформляются для вывода
	bytes, err := json.MarshalIndent(attrs, "", "  ")
	if err != nil {
		return fmt.Errorf("error when marshaling attrs: %w", err)
	}
	


	//здесь происходит вывод логов, окрашенных в указаные цвета
	fmt.Println(
		colorize(lightGray, r.Time.Format(timeFormat)),
		level,
		colorize(white, r.Message),
		colorize(darkGray, string(bytes)),
	)
	

	return nil
}



func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{h: h.h.WithAttrs(attrs), b: h.b, m: h.m}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{h: h.h.WithGroup(name), b: h.b, m: h.m}
}



//убираем встрроенные выводы времени, уровня и сообщения внутреннего slog.Handler,
//так как мы уже их выводим сами
func suppressDefaults( 						  
	next func([]string, slog.Attr) slog.Attr, 
) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey ||
			a.Key == slog.LevelKey ||
			a.Key == slog.MessageKey {
			return slog.Attr{}
		}
		if next == nil {
			return a
		}
		return next(groups, a)
	}
}



//функция для создание ручки логгера
func NewHandler(opts *slog.HandlerOptions) *Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	b := &bytes.Buffer{}
	return &Handler{
		b: b,
		h: slog.NewJSONHandler(b, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: suppressDefaults(opts.ReplaceAttr),
		}),
		m: &sync.Mutex{},
	}
}

