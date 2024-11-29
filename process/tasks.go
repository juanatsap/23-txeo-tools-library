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

	switch {
	case isCatchup || isMeeting || isCall || isExplaining:
		return "Catchups / Meetings"
	case isTraining || isDataflow || isImplementation || isScreensets || isSpeaker || isCaptcha ||
		isEmarsys || isTemplating || isPoC || isDemo || isAdding || isPassword || isRecaptcha ||
		isLiteReg || isFullReg || isScript || isForms || isBug || isEndpoint ||
		isInvestigating || isSSO || isFixing || isTesting || isImport || isUseCase || isFix || isUsers || isLaunch || isChecks || isChecking ||
		isTest || isOidc || isDocumentation || isTicket || isWeekly || isMail || isConsent || isSchema || isEnrollment || isKickoff || isAnswering ||
		isNull || isRevert || isUpdate || isImproving || isPreparing || isRipper || isGenerate:
		return "Implementation / Configuration tasks"
	case isEmail || isDocumentation || isConfluence || isDoc || isAnswer:
		return "Emails / Documentation"
	case isSlack || isTeams || isChat || isWeekly || isMail || isConsent || isSchema || isEnrollment || isKickoff || isAnswering || isAnswered || isExplaining:
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
