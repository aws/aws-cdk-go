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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentity.html
//
type CfnEmailIdentityProps struct {
	// The email address or domain to verify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentity.html#cfn-ses-emailidentity-emailidentity
	//
	EmailIdentity *string `field:"required" json:"emailIdentity" yaml:"emailIdentity"`
	// Used to associate a configuration set with an email identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentity.html#cfn-ses-emailidentity-configurationsetattributes
	//
	ConfigurationSetAttributes interface{} `field:"optional" json:"configurationSetAttributes" yaml:"configurationSetAttributes"`
	// An object that contains information about the DKIM attributes for the identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentity.html#cfn-ses-emailidentity-dkimattributes
	//
	DkimAttributes interface{} `field:"optional" json:"dkimAttributes" yaml:"dkimAttributes"`
	// If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for [Easy DKIM](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html) .
	//
	// You can only specify this object if the email identity is a domain, as opposed to an address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentity.html#cfn-ses-emailidentity-dkimsigningattributes
	//
	DkimSigningAttributes interface{} `field:"optional" json:"dkimSigningAttributes" yaml:"dkimSigningAttributes"`
	// Used to enable or disable feedback forwarding for an identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentity.html#cfn-ses-emailidentity-feedbackattributes
	//
	FeedbackAttributes interface{} `field:"optional" json:"feedbackAttributes" yaml:"feedbackAttributes"`
	// Used to enable or disable the custom Mail-From domain configuration for an email identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentity.html#cfn-ses-emailidentity-mailfromattributes
	//
	MailFromAttributes interface{} `field:"optional" json:"mailFromAttributes" yaml:"mailFromAttributes"`
}

