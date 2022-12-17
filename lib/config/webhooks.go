package config

import (
	"encoding/json"
	"os"

	"github.com/iVitaliya/logger-go"
)

type IBase struct {
	ID    string
	Token string
}

type IWebhooks struct {
	StartLogs       *IBase
	ShardLogs       *IBase
	ErrorLogs       *IBase
	DmLogs          *IBase
	VoiceLogs       *IBase
	ServerLogs      *IBase
	ServerLogs2     *IBase
	CommandLogs     *IBase
	ConsoleLogs     *IBase
	WarnLogs        *IBase
	VoiceErrorLogs  *IBase
	CreditLogs      *IBase
	EvalLogs        *IBase
	InteractionLogs *IBase
	BugReportLogs   *IBase
}

/*
Options For Fetching:

	StartLogs
	ShardLogs
	ErrorLogs
	DmLogs
	VoiceLogs
	ServerLogs
	ServerLogs2
	CommandLogs
	ConsoleLogs
	WarnLogs
	VoiceErrorLogs
	CreditLogs
	EvalLogs
	InteractionLogs
	BugReportLogs
*/
func fetch() IWebhooks {
	content, err := os.ReadFile("./webhooks.json")
	if err != nil {
		logger.Error(err.Error())
	}

	var payload IWebhooks
	err = json.Unmarshal(content, &payload)
	if err != nil {
		logger.Error("Error during (json).Unmarshal():", err.Error())
	}

	return payload
}

var Fetch IWebhooks = fetch()
