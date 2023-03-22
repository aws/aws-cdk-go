package awscognito


// User pool configuration when administrators sign users up.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	userInvitation: &userInvitationConfig{
//   		emailSubject: jsii.String("Invite to join our awesome app!"),
//   		emailBody: jsii.String("Hello {username}, you have been invited to join our awesome app! Your temporary password is {####}"),
//   		smsMessage: jsii.String("Hello {username}, your temporary password for our awesome app is {####}"),
//   	},
//   })
//
// Experimental.
type UserInvitationConfig struct {
	// The template to the email body that is sent to the user when an administrator signs them up to the user pool.
	// Experimental.
	EmailBody *string `field:"optional" json:"emailBody" yaml:"emailBody"`
	// The template to the email subject that is sent to the user when an administrator signs them up to the user pool.
	// Experimental.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// The template to the SMS message that is sent to the user when an administrator signs them up to the user pool.
	// Experimental.
	SmsMessage *string `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

