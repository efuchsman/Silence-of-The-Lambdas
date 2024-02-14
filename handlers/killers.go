package handlers

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func (h *Handler) GetKiller(req events.APIGatewayProxyRequest, tableName string, fullName string) *events.APIGatewayProxyResponse {
	if strings.Contains(fullName, " ") {
		response := BadRequest400(events.APIGatewayProxyResponse{}, "Killer", "full_name")
		return &response
	}

	killer, err := h.s.ReturnKillerByFullName(fullName, tableName)
	if err != nil {
		log.Printf("Error getting killer: %v", err)
		response := InternalError500(events.APIGatewayProxyResponse{}, "Killer", err)
		return &response
	}

	if killer == nil {
		response := NotFound404(events.APIGatewayProxyResponse{}, "Killer")
		return &response
	}

	response := OK200(events.APIGatewayProxyResponse{}, killer)
	return &response
}
