package main

import (
	"fmt"
	"log"
)

func (s *Service) Serve() {
	listenerPort := fmt.Sprintf(":%s", "8080")

	// testingGroup := s.httpServer.Group("question")

	if err := s.httpServer.Start(listenerPort); err != nil {
		log.Fatal(err)
	}
}
