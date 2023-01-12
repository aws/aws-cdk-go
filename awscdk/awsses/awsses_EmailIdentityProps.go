package awsses


// Properties for an email identity.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myHostedZone iPublicHostedZone
//
//
//   identity := ses.NewEmailIdentity(stack, jsii.String("Identity"), &emailIdentityProps{
//   	identity: ses.identity.publicHostedZone(myHostedZone),
//   	mailFromDomain: jsii.String("mail.cdk.dev"),
//   })
//
type EmailIdentityProps struct {
	// The email address or domain to verify.
	Identity Identity `field:"required" json:"identity" yaml:"identity"`
	// The configuration set to associate with the email identity.
	ConfigurationSet IConfigurationSet `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// The type of DKIM identity to use.
	DkimIdentity DkimIdentity `field:"optional" json:"dkimIdentity" yaml:"dkimIdentity"`
	// Whether the messages that are sent from the identity are signed using DKIM.
	DkimSigning *bool `field:"optional" json:"dkimSigning" yaml:"dkimSigning"`
	// Whether to receive email notifications when bounce or complaint events occur.
	//
	// These notifications are sent to the address that you specified in the `Return-Path`
	// header of the original email.
	//
	// You're required to have a method of tracking bounces and complaints. If you haven't set
	// up another mechanism for receiving bounce or complaint notifications (for example, by
	// setting up an event destination), you receive an email notification when these events
	// occur (even if this setting is disabled).
	FeedbackForwarding *bool `field:"optional" json:"feedbackForwarding" yaml:"feedbackForwarding"`
	// The action to take if the required MX record for the MAIL FROM domain isn't found when you send an email.
	MailFromBehaviorOnMxFailure MailFromBehaviorOnMxFailure `field:"optional" json:"mailFromBehaviorOnMxFailure" yaml:"mailFromBehaviorOnMxFailure"`
	// The custom MAIL FROM domain that you want the verified identity to use.
	//
	// The MAIL FROM domain
	// must meet the following criteria:
	//    - It has to be a subdomain of the verified identity
	//    - It can't be used to receive email
	//    - It can't be used in a "From" address if the MAIL FROM domain is a destination for feedback
	// forwarding emails.
	MailFromDomain *string `field:"optional" json:"mailFromDomain" yaml:"mailFromDomain"`
}

