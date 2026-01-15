package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
)

// Properties for an email identity.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var user User
//
//
//   identity := ses.NewEmailIdentity(this, jsii.String("Identity"), &EmailIdentityProps{
//   	Identity: ses.Identity_Domain(jsii.String("cdk.dev")),
//   })
//
//   identity.grantSendEmail(user)
//
type EmailIdentityProps struct {
	// The email address or domain to verify.
	Identity Identity `field:"required" json:"identity" yaml:"identity"`
	// The configuration set to associate with the email identity.
	// Default: - do not use a specific configuration set.
	//
	ConfigurationSet interfacesawsses.IConfigurationSetRef `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// The type of DKIM identity to use.
	// Default: - Easy DKIM with a key length of 2048-bit.
	//
	DkimIdentity DkimIdentity `field:"optional" json:"dkimIdentity" yaml:"dkimIdentity"`
	// Whether the messages that are sent from the identity are signed using DKIM.
	// Default: true.
	//
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
	// Default: true.
	//
	FeedbackForwarding *bool `field:"optional" json:"feedbackForwarding" yaml:"feedbackForwarding"`
	// The action to take if the required MX record for the MAIL FROM domain isn't found when you send an email.
	// Default: MailFromBehaviorOnMxFailure.USE_DEFAULT_VALUE
	//
	MailFromBehaviorOnMxFailure MailFromBehaviorOnMxFailure `field:"optional" json:"mailFromBehaviorOnMxFailure" yaml:"mailFromBehaviorOnMxFailure"`
	// The custom MAIL FROM domain that you want the verified identity to use.
	//
	// The MAIL FROM domain
	// must meet the following criteria:
	//   - It has to be a subdomain of the verified identity
	//   - It can't be used to receive email
	//   - It can't be used in a "From" address if the MAIL FROM domain is a destination for feedback
	// forwarding emails.
	// Default: - use amazonses.com
	//
	MailFromDomain *string `field:"optional" json:"mailFromDomain" yaml:"mailFromDomain"`
}

