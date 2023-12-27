package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheckControl(w *gin.Context) {
	w.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func Return404(w *gin.Context) {
	w.JSON(http.StatusNotFound, gin.H{"status": "not found"})
}
