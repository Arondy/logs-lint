# logs-lint

Статический анализатор Go-кода для проверки сообщений в логах. Работает как standalone-анализатор или плагин для `golangci-lint`.

[Файл задания](selectel_backend_golang_testovoe.pdf) для ознакомления.

## Применяемые правила
- Сообщение лога должно начинаться со строчной латинской буквы
- Разрешены только латинские буквы, цифры и пробелы
- Среди аргументов не должно быть переменных с чувствительными названиями, такими как `password`, `key` или `token`

## Требования

- Go 1.25+
- Для использования как плагина - `golangci-lint` v2.12.2+

## Установка

### Через go install

```bash
go install github.com/Arondy/logs-lint/cmd/analyzer@latest
```

### Через go get (как зависимость)

```bash
go get github.com/Arondy/logs-lint
```

### Как плагин golangci-lint

`golangci-lint` поддерживает кастомные плагины через `.custom-gcl.yml`.

Сборка кастомного бинарника:

```bash
golangci-lint custom -v
```

После сборки подключите плагин в `.golangci.yml`.

## Использование

### Установленный бинарник

```bash
analyzer file.go
```

### Локально из исходников

```bash
go run ./cmd/analyzer file.go
```

### Собранный вручную

```bash
go build -o analyzer ./cmd/analyzer
analyzer file.go
```

### Кастомный golangci-lint

```bash
./build/logs-gcl run file.go
```

## Пример с slog

```go
// incorrect-slog.go
slog.Error("Starting server")
slog.Info("server started! 🚀")
slog.Info("password " + password + " updated")
```

```
$ go run ./cmd/analyzer incorrect-slog.go
incorrect-slog.go:4:13: message should start with lowercase letter
incorrect-slog.go:5:12: message contains prohibited characters
incorrect-slog.go:6:25: message contains sensitive variable
```

> [!NOTE]
>
> Для анализа `go.uber.org/zap` пакет должен быть в `go.mod` проекта!

## Тестирование

```bash
go test ./analyzer
```

- Юнит тесты для проверки каждого правила.
- Модульные тесты для прогона анализатора на тестовых файлах через `analysistest`.

Тестовые файлы лежат в `testdata/src/logs/`, заглушка `zap` для тестирования - в `testdata/src/go.uber.org/zap`.

## Возможные доработки

- Обработка `CallExpr` - проверка метода + пакета, затем аргументы через `checkExpr()`
- Поддержка функций `zap` вида `Logf`, `Logln`, `Logw`
- Расширение списка чувствительных переменных, регистронезависимые проверки
- Конфигурация анализатора для настройки разрешённых символов, правил
