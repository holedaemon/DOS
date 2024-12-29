// Package api implements a client for the Discord REST API.
package api

import (
	"fmt"

	"github.com/holedaemon/DOS/meta"
)

const (
	baseURL = "https://discord.com/api"
	version = "10"

	contentTypeJSON = "application/json"
	contentTypeForm = "application/x-www-form-urlencoded"
	contentTypeData = "multipart/form-data"
)

// UserAgentFormat is the format of User-Agent headers expected by Discord.
// Use [fmt.Sprintf] to format accordingly.
const UserAgentFormat = "DiscordBot (%s, v%s)"

var (
	repoURL = fmt.Sprintf("https://github.com/%s", meta.GitSlug)

	userAgent = fmt.Sprintf(UserAgentFormat, repoURL, meta.Version)
)
