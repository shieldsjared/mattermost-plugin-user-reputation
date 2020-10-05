package main

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (p *Plugin) DecrementUserReputation(userId string, emojiName string) error {
	p.API.LogDebug("Decrementing user reputations", "user_id", userId, "emojiName", emojiName)
	reputations, err := p.GetUserReputations(userId)
	if err != nil {
		return err
	}
	if reputationScore, exists := reputations[emojiName]; exists {
		if reputationScore == 1 {
			delete(reputations,emojiName)
		} else {
			reputations[emojiName] = reputationScore - 1
		}
	}
	if err = p.SaveUserReputations(userId, &reputations); err != nil {
		return err
	}
	return nil
}

func (p *Plugin) GetUserReputations(userId string) (map[string]int, error) {
	p.API.LogDebug("Retrieving user reputations", "user_id", userId)
	reputations := map[string]int{}
	raw, appErr := p.API.KVGet(userId)
	if appErr != nil {
		return nil, errors.WithMessage(appErr, "failed to fetch user reputations, using KVGet")
	}
	if raw == nil {
		return reputations, nil
	}
	if err := json.Unmarshal(raw, &reputations); err != nil {
		return nil, err
	}
	p.API.LogDebug("Retrieved user reputations", "user_id", userId, "reputations", reputations)
	return reputations, nil
}

func (p *Plugin) IncrementUserReputation(userId string, emojiName string) error {
	p.API.LogDebug("Incrementing user reputations", "user_id", userId, "emojiName", emojiName)
	reputations, err := p.GetUserReputations(userId)
	if err != nil {
		return err
	}
	reputationScore, exists := reputations[emojiName]
	if(!exists) {
		reputations[emojiName] = 1
	} else {
		reputations[emojiName] = reputationScore + 1
	}
	if err = p.SaveUserReputations(userId, &reputations); err != nil {
		return err
	}
	return nil
}

func (p *Plugin) SaveUserReputations(userId string, reputations *map[string]int) error {
	p.API.LogDebug("Saving user reputations", "user_id", userId, "reputations", reputations)
	json, err := json.Marshal(reputations)
	if err != nil {
		return err
	}
	appErr := p.API.KVSet(userId, json)
	if appErr != nil {
		return errors.WithMessage(appErr, "failed to save user reputations, using KVSet")
	}
	return nil
}
