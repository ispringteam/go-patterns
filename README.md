# go-patterns

Библиотека на Go для поддержки микросервисов.

# Политика релизов

После каждого изменения кода должна выпускаться новая версия (новый тег в git).

Не следует вносить правки и выпускать версию без согласования с авторами библиотеки.

Как выпустить версию:

* Посмотрите список версий, используя `git fetch && git tag`
* Выберите имя новой версии согласно принципам [semver.org](https://semver.org/lang/ru/)
* Используйте `git tag :имя` для создания нового тега
* Используйте `git push origin :имя` для отправки тега на Github

# Как указать версию для go mod

При использовании go modules надо убедиться, что в `go.mod` прописан актуальный тег. Пример:

```go
module myservice

go 1.12

require (
	go.ispring.lan/go v1.10.0
)
```

# Модули

## errors

Модуль `"github.com/ispringtech/go-patterns/infrastructure/errors"` предоставляет функции, дополняющие пакет [github.com/pkg/errors](https://github.com/pkg/errors).

## log

Модуль `"github.com/ispringtech/go-patterns/infrastructure/log"` объявляет интерфейс `log.Logger`. Детали показаны в статье [Бережная обработка ошибок в микросервисах](https://habr.com/ru/post/459130/).

## jsonlog

Модуль `"github.com/ispringtech/go-patterns/infrastructure/jsonlog"` реализует интерфейс `log.Logger`. Детали показаны в статье [Бережная обработка ошибок в микросервисах](https://habr.com/ru/post/459130/).

Пример:

```go
import (
	"github.com/ispringtech/go-patterns/infrastructure/jsonlog"
)

func main() {
	logger := jsonlog.NewLogger(&jsonlog.Config{
			Level:   jsonlog.InfoLevel,
			AppName: "mymicroservice",
			PrettyPrint: true,
	})
	// ...
}
```
