package repository

import (
	"errors"
	"real-estate-backend/helper"
	"real-estate-backend/model"

	"gorm.io/gorm"
)

type AgentRepositoryImpl struct {
	Db *gorm.DB
}

func NewAgentReposiotryImpl(Db *gorm.DB) AgentRepository {
	return &AgentRepositoryImpl{Db: Db}
}

// Delete implements AgentRepository.
func (a *AgentRepositoryImpl) Delete(agentId int) {
	var agent model.Agent
	result := a.Db.Where("id = ?", agentId).Delete(&agent)
	helper.PanicIfError(result.Error)
}

// Save implements AgentRepository.
func (a *AgentRepositoryImpl) Save(agent model.Agent) {
	result := a.Db.Create(&agent)
	helper.PanicIfError(result.Error)
}

// Update implements AgentRepository.
func (a *AgentRepositoryImpl) Update(agent model.Agent) {
	result := a.Db.Model(&agent).Where("id = ?", agent.Id).Updates(agent)
	helper.PanicIfError(result.Error)
}

// FindByUserId implements AgentRepository.
func (a *AgentRepositoryImpl) FindByUserId(userId int) (*model.Agent, error) {
	var agent model.Agent
	err := a.Db.Where("user_id = ?", userId).First(&agent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &agent, nil
}

// FindByAgentId implements AgentRepository.
func (a *AgentRepositoryImpl) FindByAgentId(agentId int) *model.Agent {
	var agent model.Agent
	err := a.Db.Where("Id = ?", agentId).First(&agent).Error

	helper.PanicIfError(err)
	return &agent
}
