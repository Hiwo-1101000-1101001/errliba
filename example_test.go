package errlib

import (
    "context"
    "errors"
    "fmt"
    "log"
    "os"
    "testing"
)

func TestErrorHandling(t *testing.T) {
    // Создаём новый ErrorHandler
    handler := NewErrorHandler().WithLogger(log.New(os.Stdout, "", log.LstdFlags)).WithHook(func(e *ErrorWrapper) {
        // Кастомный хук, например, отправка в мониторинг
        fmt.Printf("Hook: Отправка ошибки в мониторинг: %v\n", e)
    })

    // Пример использования
    err := doSomething()
    if err != nil {
        handler.Handle(err).WithMessage("Ошибка при выполнении doSomething").WithSeverity(Critical)
    }
}

func doSomething() error {
    // Возвращаем простую ошибку
    return errors.New("исходная ошибка")
}

func TestContext(t *testing.T) {
    ctx := context.Background()
    err := errors.New("ошибка для контекста")
    ctx = WithError(ctx, err)

    // Передаём контекст в другую функцию
    processContext(ctx)
}

func processContext(ctx context.Context) {
    if err := ErrorFromContext(ctx); err != nil {
        fmt.Printf("Получена ошибка из контекста: %v\n", err)
    }
}

