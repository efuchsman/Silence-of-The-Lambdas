package handlers

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func (h *Handler) GetVictimsByKiller(req events.APIGatewayProxyRequest, tableName string, killerName string) *events.APIGatewayProxyResponse {
	if strings.Contains(killerName, " ") {
		response := BadRequest400(events.APIGatewayProxyResponse{}, "Killer", "full_name")
		return &response
	}

	victims, err := h.s.ReturnVictimsByKiller(killerName, tableName)
	if err != nil {
		log.Printf("Error getting victims: %v", err)
		response := InternalError500(events.APIGatewayProxyResponse{}, "Victims", err)
		return &response
	}

	if victims == nil {
		response := NotFound404(events.APIGatewayProxyResponse{}, "Victims")
		return &response
	}

	response := OK200(events.APIGatewayProxyResponse{}, victims)
	return &response
}
