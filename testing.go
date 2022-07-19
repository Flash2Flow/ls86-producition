package main

import (
	"log"
)

func (s *Server) checkAll() (bool, error) {

	rest, err2 := ls86.Server.checkREST()
	if rest == true {

		acp, err3 := ls86.Server.checkACP()
		if acp == true {
			ls86.Logs.Server.AllOk()
			return true, nil
		} else {
			log.Println(err3)
			return false, err2
		}

	} else {
		log.Println(err2)
		return false, err2
	}

}

func (s *Server) checkREST() (bool, error) {
	return true, nil
}

func (s *Server) checkACP() (bool, error) {
	return true, nil
}
