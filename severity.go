package errlib

// String возвращает строковое представление Severity
func (s Severity) String() string {
    switch s {
    case Debug:
        return "DEBUG"
    case Info:
        return "INFO"
    case Error:
        return "ERROR"
    case Critical:
        return "CRITICAL"
    default:
        return "UNKNOWN"
    }
}
