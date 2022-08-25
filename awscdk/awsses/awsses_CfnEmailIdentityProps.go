package awsses


// Properties for defining a `CfnEmailIdentity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEmailIdentityProps := &cfnEmailIdentityProps{
//   	emailIdentity: jsii.String("emailIdentity"),
//
//   	// the properties below are optional
//   	configurationSetAttributes: &configurationSetAttributesProperty{
//   		configurationSetName: jsii.String("configurationSetName"),
//   	},
//   	dkimAttributes: &dkimAttributesProperty{
//   		signingEnabled: jsii.Boolean(false),
//   	},
//   	dkimSigningAttributes: &dkimSigningAttributesProperty{
//   		domainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   		domainSigningSelector: jsii.String("domainSigningSelector"),
//   		nextSigningKeyLength: jsii.String("nextSigningKeyLength"),
//   	},
//   	feedbackAttributes: &feedbackAttributesProperty{
//   		emailForwardingEnabled: jsii.Boolean(false),
//   	},
//   	mailFromAttributes: &mailFromAttributesProperty{
//   		behaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   		mailFromDomain: jsii.String("mailFromDomain"),
//   	},
//   }
//
type CfnEmailIdentityProps struct {
	// `AWS::SES::EmailIdentity.EmailIdentity`.
	EmailIdentity *string `field:"required" json:"emailIdentity" yaml:"emailIdentity"`
	// `AWS::SES::EmailIdentity.ConfigurationSetAttributes`.
	ConfigurationSetAttributes interface{} `field:"optional" json:"configurationSetAttributes" yaml:"configurationSetAttributes"`
	// `AWS::SES::EmailIdentity.DkimAttributes`.
	DkimAttributes interface{} `field:"optional" json:"dkimAttributes" yaml:"dkimAttributes"`
	// `AWS::SES::EmailIdentity.DkimSigningAttributes`.
	DkimSigningAttributes interface{} `field:"optional" json:"dkimSigningAttributes" yaml:"dkimSigningAttributes"`
	// `AWS::SES::EmailIdentity.FeedbackAttributes`.
	FeedbackAttributes interface{} `field:"optional" json:"feedbackAttributes" yaml:"feedbackAttributes"`
	// `AWS::SES::EmailIdentity.MailFromAttributes`.
	MailFromAttributes interface{} `field:"optional" json:"mailFromAttributes" yaml:"mailFromAttributes"`
}

