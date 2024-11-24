package repository

import "real-estate-backend/model"

type AgentRepository interface {
	Save(agent model.Agent)
	Update(agent model.Agent)
	Delete(agentId int)
	FindByUserId(userId int) (*model.Agent, error)
	FindByAgentId(agentId int) *model.Agent
}
