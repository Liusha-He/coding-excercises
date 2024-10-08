package service

import "github.com/gin-gonic/gin"

func (server *Server) UsersRouter(r *gin.RouterGroup) {
	r.POST("/users", server.createUser)
	r.POST("/users/login", server.login)
}

func (server *Server) AccountsRouters(r *gin.RouterGroup) {
	r.POST("/accounts", server.createAccount)
	r.GET("/accounts/:id", server.getAccount)
	r.GET("/accounts", server.listAccounts)
}

func (server *Server) TransfersRouters(r *gin.RouterGroup) {
	r.POST("/transfers", server.createTransfer)
}
