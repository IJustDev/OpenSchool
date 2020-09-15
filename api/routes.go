package main

func (s *server) routes() {
	s.router.HandleFunc("/login", s.handleAuthLogin())
}
