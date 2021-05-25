package factory

import "log"

type student struct {
}

func (s *student) DoMorning() error {

	log.Println("STUDENT:等等...,再等等...,再等等我嘛!")
	return nil
}
