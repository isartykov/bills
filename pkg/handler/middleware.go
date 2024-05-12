package handler

import (
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

// func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) {
// 	header := r.Header.Get(authorizationHeader)
// 	if header == "" {
// 		newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
// 		return
// 	}

// 	headerParts := strings.Split(header, " ")
// 	if len(headerParts) != 2 {
// 		newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
// 		return
// 	}

// 	_, err := h.services.Authorization.ParseToken(headerParts[1])
// 	if err != nil {
// 		newErrorResponse(w, http.StatusUnauthorized, err.Error())
// 		return
// 	}
// }

func (h *Handler) userIdentity(handler func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userID, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		handler(w, r, userID)
	}
}
