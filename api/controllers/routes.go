package controllers

import "medium-clone-backend/api/middlewares"

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		//Users routes
		v1.POST("/users", middlewares.TokenAuthMiddleware(), s.CreateUser)
		v1.GET("/users", middlewares.TokenAuthMiddleware(), s.GetUsers)
		v1.GET("/users/:id", middlewares.TokenAuthMiddleware(), s.GetUser)
		v1.PUT("/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateUser)
		v1.PUT("/avatar/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateAvatar)
		v1.DELETE("/users/:id", middlewares.TokenAuthMiddleware(), s.DeleteUser)

		// Reset password:
		v1.POST("/password/forgot", s.ForgotPassword)
		v1.POST("/password/reset", s.ResetPassword)

		// Login Route
		v1.POST("/login", s.Login)

	}
}
