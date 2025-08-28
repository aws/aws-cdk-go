package awsses


// The action to take if the required MX record for the MAIL FROM domain isn't found.
type MailFromBehaviorOnMxFailure string

const (
	// The mail is sent using amazonses.com as the MAIL FROM domain.
	MailFromBehaviorOnMxFailure_USE_DEFAULT_VALUE MailFromBehaviorOnMxFailure = "USE_DEFAULT_VALUE"
	// The Amazon SES API v2 returns a `MailFromDomainNotVerified` error and doesn't attempt to deliver the email.
	MailFromBehaviorOnMxFailure_REJECT_MESSAGE MailFromBehaviorOnMxFailure = "REJECT_MESSAGE"
)

