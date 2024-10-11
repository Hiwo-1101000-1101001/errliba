package errlib

import (
    "fmt"
    "runtime"
)

// Уровень серьезности ошибки
type Severity int

const (
    Debug Severity = iota
    Info
    Warning
    Error
    Critical
)

// Оборачиваем стандартную ошибку и добавляем контекст
type ErrorWrapper struct {
    Err error
    Message string
    Severity Severity
    Cause error
    Stack []uintptr
}

// New Создает новый ErrorWrapper
func New(err error) *ErrorWrapper {
    if err == nil {
        return nil
    }

    return &ErrorWrapper{
        Err:      err,
        Severity: Error,
        Stack:    captureStack(),
    }
}

// Error возвращает строковое представление ошибки
func (e *ErrorWrapper) Error() string {
    if e.Message != "" {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }

    return e.Err.Error()
}

// Unwrap позволяет использовать errors.Unwrap
func (e *ErrorWrapper) Unwrap() error {
    return e.Cause
}

// WithMessage добавляет сообщение к ошибке
func (e *ErrorWrapper) WithMessage(msg string) *ErrorWrapper {
    e.Message = msg
    return e
}

// WithSeverity Устанавливает уровень серьезности
func (e *ErrorWrapper) WithSeverity(sev Severity) *ErrorWrapper {
    e.Severity = sev
    return e
}

// WithCause устанавливает причину ошибки
func (e *ErrorWrapper) WithCause(cause error) *ErrorWrapper {
    e.Cause = cause
    return e
}

// captureStack захватывает стек вызовов
func captureStack() []uintptr {
    const depth = 32
    pcs := make([]uintptr, depth)
    n := runtime.Callers(3, pcs)
    
    return pcs[:n]
}

// Format реализует интерфейс fmt.Formatter для форматированного вывода
func (e *ErrorWrapper) Format(s fmt.State, verb rune) {
    switch verb {
    case 'v':
        if s.Flag('+') {
            fmt.Fprintf(s, "%s\n", e.Error())
            for _, pc := range e.Stack {
                fn := runtime.FuncForPC(pc)
                file, line := fn.FileLine(pc)
                fmt.Fprintf(s, "\t%s\n\t\t%s:%d\n", fn.Name(), file, line)
            }
            if e.Cause != nil {
                fmt.Fprintf(s, "\nCaused by: %+v", e.Cause)
            }
            return
        }
        fallthrough
    case 's':
        fmt.Fprintf(s, e.Error())
    case 'q':
        fmt.Fprintf(s, "%q", e.Error())
    }
}
