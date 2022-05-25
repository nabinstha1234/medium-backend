package controllers

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		//Users routes
		v1.POST("/users", s.CreateUser)
		v1.GET("/users", s.GetUsers)
		v1.GET("/users/:id", s.GetUser)
		v1.PUT("/users/:id", s.UpdateUser)
		v1.PUT("/avatar/users/:id", s.UpdateAvatar)
		v1.DELETE("/users/:id", s.DeleteUser)

		// Reset password:
		v1.POST("/password/forgot", s.ForgotPassword)
		v1.POST("/password/reset", s.ResetPassword)

		// Login Route
		v1.POST("/login", s.Login)

	}
}
