package awscognito


// Configuration for Cognito sending emails via Amazon SES.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	Email: cognito.UserPoolEmail_WithSES(&UserPoolSESOptions{
//   		FromEmail: jsii.String("noreply@myawesomeapp.com"),
//   		FromName: jsii.String("Awesome App"),
//   		ReplyTo: jsii.String("support@myawesomeapp.com"),
//   	}),
//   })
//
type UserPoolSESOptions struct {
	// The verified Amazon SES email address that Cognito should use to send emails.
	//
	// The email address used must be a verified email address
	// in Amazon SES and must be configured to allow Cognito to
	// send emails.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-email.html
	//
	FromEmail *string `field:"required" json:"fromEmail" yaml:"fromEmail"`
	// The name of a configuration set in Amazon SES that should be applied to emails sent via Cognito.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-configurationset
	//
	// Default: - no configuration set.
	//
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
	// An optional name that should be used as the sender's name along with the email.
	// Default: - no name.
	//
	FromName *string `field:"optional" json:"fromName" yaml:"fromName"`
	// The destination to which the receiver of the email should reply to.
	// Default: - same as the fromEmail.
	//
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
	// Required if the UserPool region is different than the SES region.
	//
	// If sending emails with a Amazon SES verified email address,
	// and the region that SES is configured is different than the
	// region in which the UserPool is deployed, you must specify that
	// region here.
	// Default: - The same region as the Cognito UserPool.
	//
	SesRegion *string `field:"optional" json:"sesRegion" yaml:"sesRegion"`
	// SES Verified custom domain to be used to verify the identity.
	// Default: - no domain.
	//
	SesVerifiedDomain *string `field:"optional" json:"sesVerifiedDomain" yaml:"sesVerifiedDomain"`
}

