package errlib

import (
    "log"
    "os"
)

// ErrorHandler интерфейс для обработки ошибок
type ErrorHandler interface {
    Handle(err error) *ErrorWrapper
    WithLogger(logger *log.Logger) ErrorHandler
    WithHook(hook func(*ErrorWrapper)) ErrorHandler
}

type errorHandler struct {
    logger *log.Logger
    hook   func(*ErrorWrapper)
}

// NewErrorHandler создает новый ErrorHandler
func NewErrorHandler() ErrorHandler {
    return &errorHandler{
        logger: log.New(os.Stderr, "", log.LstdFlags),
    }
}

// Handle обрабатывает ошибку
func (h *errorHandler) Handle(err error) *ErrorWrapper {
    if err == nil {
        return nil
    }

    var ew *ErrorWrapper
    if e, ok := err.(*ErrorWrapper); ok {
        ew = e
    } else {
        ew = New(err)
    }

    if h.hook != nil {
        h.hook(ew)
    }

    h.log(ew)
    return ew
}

// WithLogger устанавливает кастмный логгер
func (h *errorHandler) WithLogger(logger *log.Logger) ErrorHandler {
    h.logger = logger
    return h
}

// WithHook устанавливает функцию-хук
func (h *errorHandler) WithHook(hook func (*ErrorWrapper)) ErrorHandler {
    h.hook = hook
    return h
}

func (h *errorHandler) log(e *ErrorWrapper) {
    prefix := "[" + e.Severity.String() + "] "
    h.logger.SetPrefix(prefix)
    h.logger.Println(e.Error())

    if e.Severity == Critical {
        panic(e)
    }
}
