package chainOfResponsibility

import (
	"fmt"
	"time"
)

type Logger struct {
	next Handler
}

func (l *Logger) SetNext(n Handler) {
	l.next = n
}

func (l *Logger) Execute(ctx IContext) error {
	var err error
	start := time.Now()
	if l.next != nil {
		err = l.next.Execute(ctx)
	}
	fmt.Printf("本次处理时间:%+v\n", time.Since(start))
	return err
}
