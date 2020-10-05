module github.com/mattermost/mattermost-plugin-starter-template

go 1.12

require (
	github.com/mattermost/mattermost-server/v5 v5.26.2
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
)

replace github.com/mattermost/mattermost-server/v5 v5.26.2 => github.com/shieldsjared/mattermost-server/v5 v5.0.0-20201004175034-613d90476472
