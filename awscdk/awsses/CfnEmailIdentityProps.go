package awsses


// Properties for defining a `CfnEmailIdentity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEmailIdentityProps := &CfnEmailIdentityProps{
//   	EmailIdentity: jsii.String("emailIdentity"),
//
//   	// the properties below are optional
//   	ConfigurationSetAttributes: &ConfigurationSetAttributesProperty{
//   		ConfigurationSetName: jsii.String("configurationSetName"),
//   	},
//   	DkimAttributes: &DkimAttributesProperty{
//   		SigningEnabled: jsii.Boolean(false),
//   	},
//   	DkimSigningAttributes: &DkimSigningAttributesProperty{
//   		DomainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   		DomainSigningSelector: jsii.String("domainSigningSelector"),
//   		NextSigningKeyLength: jsii.String("nextSigningKeyLength"),
//   	},
//   	FeedbackAttributes: &FeedbackAttributesProperty{
//   		EmailForwardingEnabled: jsii.Boolean(false),
//   	},
//   	MailFromAttributes: &MailFromAttributesProperty{
//   		BehaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   		MailFromDomain: jsii.String("mailFromDomain"),
//   	},
//   }
//
type CfnEmailIdentityProps struct {
	// The email address or domain to verify.
	EmailIdentity *string `field:"required" json:"emailIdentity" yaml:"emailIdentity"`
	// Used to associate a configuration set with an email identity.
	ConfigurationSetAttributes interface{} `field:"optional" json:"configurationSetAttributes" yaml:"configurationSetAttributes"`
	// An object that contains information about the DKIM attributes for the identity.
	DkimAttributes interface{} `field:"optional" json:"dkimAttributes" yaml:"dkimAttributes"`
	// If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for [Easy DKIM](https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html) .
	DkimSigningAttributes interface{} `field:"optional" json:"dkimSigningAttributes" yaml:"dkimSigningAttributes"`
	// Used to enable or disable feedback forwarding for an identity.
	FeedbackAttributes interface{} `field:"optional" json:"feedbackAttributes" yaml:"feedbackAttributes"`
	// Used to enable or disable the custom Mail-From domain configuration for an email identity.
	MailFromAttributes interface{} `field:"optional" json:"mailFromAttributes" yaml:"mailFromAttributes"`
}

