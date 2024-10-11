package errlib

import "context"

type contextKey struct{}

// WithError добавляет ошибку в контекст
func WithError(ctx context.Context, err error) context.Context {
    return context.WithValue(ctx, contextKey{}, err)
}

// ErrorFromContext извлекает ошибку из контекста
func ErrorFromContext(ctx context.Context) error {
    if err, ok := ctx.Value(contextKey{}).(error); ok {
        return err
    }

    return nil
}
