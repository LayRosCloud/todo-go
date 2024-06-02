package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const HeaderAuthorization = "Authorization"

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(HeaderAuthorization)

	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "Authorization header is clear")
		return
	}

	els := strings.Split(header, " ")
	if len(els) != 2 || els[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "Authorization header has bad form")
		return
	}
	token := els[1]

	userId, err := h.services.Authorization.ParseToken(token)

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "token has bad form")
		return
	}

	c.Set("userId", userId)
}
