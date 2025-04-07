package process

import "strings"

func GetTaskCategory(taskListName string) string {
	taskListNameLower := strings.ToLower(taskListName)

	isCatchup := strings.Contains(taskListNameLower, "catchup")
	isCall := strings.Contains(taskListNameLower, "call")
	isMeeting := strings.Contains(taskListNameLower, "meeting")
	isTraining := strings.Contains(taskListNameLower, "training")
	isDataflow := strings.Contains(taskListNameLower, "dataflow")
	isImplementation := strings.Contains(taskListNameLower, "implementation")
	isEmail := strings.Contains(taskListNameLower, "email")
	isScreensets := strings.Contains(taskListNameLower, "screensets")
	isEmarsys := strings.Contains(taskListNameLower, "emarsys")
	isTemplating := strings.Contains(taskListNameLower, "templating")
	isPoC := strings.Contains(taskListNameLower, "poc")
	isDocumentation := strings.Contains(taskListNameLower, "documentation")
	isSlack := strings.Contains(taskListNameLower, "slack")
	isTeams := strings.Contains(taskListNameLower, "teams")
	isDemo := strings.Contains(taskListNameLower, "demoing")
	isChat := strings.Contains(taskListNameLower, "chat")
	isAdding := strings.Contains(taskListNameLower, "adding")
	isPassword := strings.Contains(taskListNameLower, "password")
	isConfluence := strings.Contains(taskListNameLower, "confluence")
	isRecaptcha := strings.Contains(taskListNameLower, "recaptcha")

	isLiteReg := strings.Contains(taskListNameLower, "lite reg")
	isFullReg := strings.Contains(taskListNameLower, "full reg")
	isScript := strings.Contains(taskListNameLower, "script")
	isForms := strings.Contains(taskListNameLower, "forms")
	isBug := strings.Contains(taskListNameLower, " bug ")
	isEndpoint := strings.Contains(taskListNameLower, "endpoint")
	isInvestigating := strings.Contains(taskListNameLower, "investigating ")
	isSSO := strings.Contains(taskListNameLower, " sso ")
	isFixing := strings.Contains(taskListNameLower, "fixing ")
	isTesting := strings.Contains(taskListNameLower, "testing")
	isImport := strings.Contains(taskListNameLower, "import")
	isUseCase := strings.Contains(taskListNameLower, "use case")
	isSpeaker := strings.Contains(taskListNameLower, "speak")
	isFix := strings.Contains(taskListNameLower, "fix")
	isUsers := strings.Contains(taskListNameLower, "users")
	isLaunch := strings.Contains(taskListNameLower, "launch")
	isChecks := strings.Contains(taskListNameLower, "checks")
	isChecking := strings.Contains(taskListNameLower, "checking")
	isTest := strings.Contains(taskListNameLower, "test")
	isOidc := strings.Contains(taskListNameLower, "oidc")
	isDoc := strings.Contains(taskListNameLower, "doc")
	isTicket := strings.Contains(taskListNameLower, "ticket")
	isConsent := strings.Contains(taskListNameLower, "consent")
	isWeekly := strings.Contains(taskListNameLower, "weekly")
	isMail := strings.Contains(taskListNameLower, "mail")
	isSchema := strings.Contains(taskListNameLower, "schema")
	isEnrollment := strings.Contains(taskListNameLower, "enrollment")
	isKickoff := strings.Contains(taskListNameLower, "kickoff")
	isAnswering := strings.Contains(taskListNameLower, "answering")
	isCaptcha := strings.Contains(taskListNameLower, "captcha")
	isNull := strings.Contains(taskListNameLower, "null")
	isRevert := strings.Contains(taskListNameLower, "revert")
	isUpdate := strings.Contains(taskListNameLower, "update")
	isImproving := strings.Contains(taskListNameLower, "improving")
	isPreparing := strings.Contains(taskListNameLower, "preparing")
	isAnswer := strings.Contains(taskListNameLower, "answer")
	isExplaining := strings.Contains(taskListNameLower, "explaining")
	isAnswered := strings.Contains(taskListNameLower, "answered")
	isRipper := strings.Contains(taskListNameLower, "ripper")
	isGenerate := strings.Contains(taskListNameLower, "generate")
	isCatchupII := strings.Contains(taskListNameLower, "catch up")
	isReport := strings.Contains(taskListNameLower, "report")
	isHolidays := strings.Contains(taskListNameLower, "holidays")
	isDiscussion := strings.Contains(taskListNameLower, "discussion")
	isConversation := strings.Contains(taskListNameLower, "conversation")
	isSupporting := strings.Contains(taskListNameLower, "supporting")
	isCss := strings.Contains(taskListNameLower, "css")
	isEvents := strings.Contains(taskListNameLower, "events")
	isCNAME := strings.Contains(taskListNameLower, "cname")
	isWebhooks := strings.Contains(taskListNameLower, "webhooks")
	isConver := strings.Contains(taskListNameLower, "conver")
	isProblem := strings.Contains(taskListNameLower, "problem")
	isDiscussing := strings.Contains(taskListNameLower, "discussing")
	isCerts := strings.Contains(taskListNameLower, "certs")
	isInvestigate := strings.Contains(taskListNameLower, "investigate")
	isGoLive := strings.Contains(taskListNameLower, "go-live")
	isBackend := strings.Contains(taskListNameLower, "backend")
	isNextSteps := strings.Contains(taskListNameLower, "next steps")
	isGithub := strings.Contains(taskListNameLower, "github")
	isData := strings.Contains(taskListNameLower, "data")
	isDataFlow := strings.Contains(taskListNameLower, "dataflow")
	isIncidence := strings.Contains(taskListNameLower, "incidence")
	isExtensions := strings.Contains(taskListNameLower, "extensions")
	isArrays := strings.Contains(taskListNameLower, "arrays")
	isCLP := strings.Contains(taskListNameLower, "clp")
	isStruct := strings.Contains(taskListNameLower, "struct")
	isStructure := strings.Contains(taskListNameLower, "structure")
	isLIVX := strings.Contains(taskListNameLower, "livx")
	isRegistrationCompletion := strings.Contains(taskListNameLower, "registration completion")
	isTypescript := strings.Contains(taskListNameLower, "typescript")
	isReact := strings.Contains(taskListNameLower, "react")
	isSaturday := strings.Contains(taskListNameLower, "saturday")
	isSunday := strings.Contains(taskListNameLower, "sunday")
	isInvoice := strings.Contains(taskListNameLower, "invoice")
	isNeopoly := strings.Contains(taskListNameLower, "neopoly")
	isDeletionProcess := strings.Contains(taskListNameLower, "deletion process")
	isDeletion := strings.Contains(taskListNameLower, "deletion")
	isUserFlows := strings.Contains(taskListNameLower, "user flows")
	isLogs := strings.Contains(taskListNameLower, "logs")
	isCreatingSystem := strings.Contains(taskListNameLower, "creating system")
	isResponsive := strings.Contains(taskListNameLower, "responsive")
	isFantasy := strings.Contains(taskListNameLower, "fantasy")
	isTasks := strings.Contains(taskListNameLower, "tasks")
	isLPC := strings.Contains(taskListNameLower, "lpc")
	isOIDC := strings.Contains(taskListNameLower, "oidc")
	isAPIKey := strings.Contains(taskListNameLower, "api key")
	isBOTF := strings.Contains(taskListNameLower, "botf")
	isEnvFiles := strings.Contains(taskListNameLower, "env files")
	isIssues := strings.Contains(taskListNameLower, "issues")
	isIssue := strings.Contains(taskListNameLower, "issue")
	isPreferences := strings.Contains(taskListNameLower, "preferences")
	isFrontal := strings.Contains(taskListNameLower, "frontal")
	isMonitoring := strings.Contains(taskListNameLower, "monitoring")
	isExport := strings.Contains(taskListNameLower, "export")
	isBlacklist := strings.Contains(taskListNameLower, "blacklist")
	isCDC := strings.Contains(taskListNameLower, "cdc")
	isGlances := strings.Contains(taskListNameLower, "glances")
	isBackfields := strings.Contains(taskListNameLower, "backfields")
	isBackfill := strings.Contains(taskListNameLower, "backfill")
	switch {
	case isCatchup || isMeeting || isCall || isExplaining || isSupporting || isSaturday || isSunday || isInvoice || isLIVX || isNeopoly || isCreatingSystem || isResponsive || isFantasy:
		return "Catchups / Meetings"
	case isTraining || isDataflow || isImplementation || isScreensets || isSpeaker || isCaptcha || isBackend || isGithub || isData || isDataFlow || isExtensions || isAPIKey || isBOTF || isEnvFiles || isIssues || isIssue || isPreferences || isFrontal ||
		isEmarsys || isTemplating || isPoC || isDemo || isAdding || isPassword || isRecaptcha ||
		isRegistrationCompletion || isTypescript || isReact || isArrays || isCLP || isStruct || isStructure ||
		isLiteReg || isFullReg || isScript || isForms || isBug || isEndpoint ||
		isInvestigating || isSSO || isFixing || isTesting || isImport || isUseCase || isFix || isUsers || isLaunch || isChecks || isChecking ||
		isTest || isOidc || isConfluence || isDocumentation || isTicket || isWeekly || isMail || isConsent || isSchema || isEnrollment || isKickoff || isAnswering ||
		isNull || isRevert || isUpdate || isImproving || isPreparing || isRipper || isGenerate || isCatchupII || isCss || isEvents || isProblem || isInvestigate || isGoLive || isLogs ||
		isDeletionProcess || isDeletion || isUserFlows || isLPC || isOIDC || isAPIKey || isBOTF || isEnvFiles || isPreferences || isFrontal || isGlances || isTasks || isMonitoring || isExport || isBlacklist || isCDC:
		return "Implementation / Configuration tasks"
	case isEmail || isDocumentation || isConfluence || isDoc || isAnswer || isReport || isCss || isCNAME || isWebhooks || isCerts || isBackfields || isBackfill ||
		isNextSteps || isIncidence:
		return "Emails / Documentation"
	case isSlack || isTeams || isChat || isWeekly || isMail || isConsent || isSchema || isEnrollment || isKickoff || isAnswering || isAnswered || isMeeting ||
		isExplaining || isHolidays || isDiscussion || isConversation || isConver || isDiscussing || isIssue || isIssues:
		return "Slack / Teams Conversations"
	default:
		return "Other"
	}
}
func GetIconForCategory(category string) string {
	switch category {
	case "Catchups / Meetings":
		return "📅"
	case "Implementation / Configuration tasks":
		return "💪"
	case "Emails / Documentation":
		return "📧"
	case "Slack / Teams Conversations":
		return "💬"
	default:
		return "❓"
	}
}
