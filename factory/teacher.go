package factory

import "log"

type teacher struct {
}

func (t *teacher) DoMorning() error {
	log.Println("TEACHER:立即起床")
	return nil
}
