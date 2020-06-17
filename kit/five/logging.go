package five

import (
	"time"

	"github.com/go-kit/kit/log"
	. "github.com/myesui/uuid"
	"github.com/myesui/uuid/kit"
)

var _ Service = &loggingMiddleware{}

type loggingMiddleware struct {
	kit.LoggingMiddleware
}

func (o loggingMiddleware) UUID(namespace Implementation, names ...interface{}) (id UUID) {
	defer func(begin time.Time) {
		o.Log(begin, "method", "uuid", "namespace", namespace, "names", names, "uuid", id, "error", "false")
	}(time.Now())

	id = o.Next().(Service).UUID(namespace, names...)
	return
}

func (o *loggingMiddleware) Add(service kit.UUIDKit) Service {
	return kit.AddMiddleware(service, o).(Service)
}

func (loggingMiddleware) String() string {
	return "uuid-five-logging"
}

func NewLoggingMiddleware(logger log.Logger) Service {
	return &loggingMiddleware{
		kit.LoggingMiddleware{
			Logger: logger,
		},
	}
}
