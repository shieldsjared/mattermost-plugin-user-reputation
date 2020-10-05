package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) ReactionHasBeenAdded(c *plugin.Context, reaction *model.Reaction) {
	post, appErr := p.API.GetPost(reaction.PostId)
	if appErr != nil {
		p.API.LogError(
			"Falied to fetch post to get author (in order to increment reputation).",
			"post_id", reaction.PostId)
	}
	if err := p.IncrementUserReputation(post.UserId, reaction.EmojiName); err != nil {
		p.API.LogError(
			"Encountered an error while incrementing user reputation", 
			"error", err, 
			"user_id", post.UserId,
			"emoji_name", reaction.EmojiName)
	}
}

func (p *Plugin) ReactionHasBeenRemoved(c *plugin.Context, reaction *model.Reaction) {
	post, appErr := p.API.GetPost(reaction.PostId)
	if appErr != nil {
		p.API.LogError(
			"Falied to fetch post to get author (in order to decrement reputation).",
			"post_id", reaction.PostId)
	}
	if err := p.DecrementUserReputation(post.UserId, reaction.EmojiName); err != nil {
		if(err != nil) {
			p.API.LogError(
				"Encountered an error while decrementing user reputation", 
				"error", err, 
				"user_id", post.UserId,
				"emoji_name", reaction.EmojiName)
		}
	}
}