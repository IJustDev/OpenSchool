package main

func (s *server) routes() {
	userRepository := LiveUserRepository{}
	s.router.HandleFunc("/login", s.handleAuthLogin(userRepository)).Methods("POST")
	s.router.HandleFunc("/register", s.handleAuthLogin(userRepository)).Methods("POST")
	s.router.HandleFunc("/students/{id}", s.handleStudentDetails(userRepository)).Methods("GET")
}
