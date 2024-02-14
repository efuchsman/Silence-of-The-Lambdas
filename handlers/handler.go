package handlers

import (
	silence "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambdas"
)

type Handler struct {
	s silence.Client
}

func NewHandler(s silence.Client) *Handler {
	return &Handler{
		s: s,
	}
}
