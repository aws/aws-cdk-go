package awscognito


// How will a user be able to recover their account?
//
// When a user forgets their password, they can have a code sent to their verified email or verified phone to recover their account.
// You can choose the preferred way to send codes below.
// We recommend not allowing phone to be used for both password resets and multi-factor authentication (MFA).
//
// Example:
//   cognito.NewUserPool(this, jsii.String("UserPool"), &userPoolProps{
//   	// ...
//   	accountRecovery: cognito.accountRecovery_EMAIL_ONLY,
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-recover-a-user-account.html
//
type AccountRecovery string

const (
	// Email if available, otherwise phone, but don’t allow a user to reset their password via phone if they are also using it for MFA.
	AccountRecovery_EMAIL_AND_PHONE_WITHOUT_MFA AccountRecovery = "EMAIL_AND_PHONE_WITHOUT_MFA"
	// Phone if available, otherwise email, but don’t allow a user to reset their password via phone if they are also using it for MFA.
	AccountRecovery_PHONE_WITHOUT_MFA_AND_EMAIL AccountRecovery = "PHONE_WITHOUT_MFA_AND_EMAIL"
	// Email only.
	AccountRecovery_EMAIL_ONLY AccountRecovery = "EMAIL_ONLY"
	// Phone only, but don’t allow a user to reset their password via phone if they are also using it for MFA.
	AccountRecovery_PHONE_ONLY_WITHOUT_MFA AccountRecovery = "PHONE_ONLY_WITHOUT_MFA"
	// (Not Recommended) Phone if available, otherwise email, and do allow a user to reset their password via phone if they are also using it for MFA.
	AccountRecovery_PHONE_AND_EMAIL AccountRecovery = "PHONE_AND_EMAIL"
	// None – users will have to contact an administrator to reset their passwords.
	AccountRecovery_NONE AccountRecovery = "NONE"
)

