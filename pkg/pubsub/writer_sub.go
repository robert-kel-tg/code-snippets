package pubsub

import (
	"errors"
	"io"
)

type writerSubscriber struct {
	id     int
	Writer io.Writer
}

func (ws *writerSubscriber) Notify(msg interface{}) error {
	return errors.New("not implemented yet")
}

func (ws *writerSubscriber) Close() {}

func NewWriterSubscriber(id int, Writer io.Writer) Subscriber {
	return &writerSubscriber{
		id:     id,
		Writer: Writer,
	}
}
