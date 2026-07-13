// Stub package for zap tests

package zap

type Field struct{}

type SugaredLogger struct{}

type Logger struct{}

func New() *Logger                                { return nil }
func NewExample() *Logger                         { return nil }
func L() *Logger                                  { return nil }
func (*Logger) Log(msg string, fields ...Field)   {}
func (*Logger) Debug(msg string, fields ...Field) {}
func (*Logger) Info(msg string, fields ...Field)  {}
func (*Logger) Warn(msg string, fields ...Field)  {}
func (*Logger) Error(msg string, fields ...Field) {}
func (*Logger) Fatal(msg string, fields ...Field) {}
func (*Logger) Sugar() *SugaredLogger             { return nil }
