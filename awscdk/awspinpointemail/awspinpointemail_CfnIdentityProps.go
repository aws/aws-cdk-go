package awspinpointemail


// Properties for defining a `CfnIdentity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentityProps := &cfnIdentityProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dkimSigningEnabled: jsii.Boolean(false),
//   	feedbackForwardingEnabled: jsii.Boolean(false),
//   	mailFromAttributes: &mailFromAttributesProperty{
//   		behaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   		mailFromDomain: jsii.String("mailFromDomain"),
//   	},
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIdentityProps struct {
	// The address or domain of the identity, such as *sender@example.com* or *example.co.uk* .
	Name *string `field:"required" json:"name" yaml:"name"`
	// For domain identities, this attribute is used to enable or disable DomainKeys Identified Mail (DKIM) signing for the domain.
	//
	// If the value is `true` , then the messages that you send from the domain are signed using both the DKIM keys for your domain, as well as the keys for the `amazonses.com` domain. If the value is `false` , then the messages that you send are only signed using the DKIM keys for the `amazonses.com` domain.
	DkimSigningEnabled interface{} `field:"optional" json:"dkimSigningEnabled" yaml:"dkimSigningEnabled"`
	// Used to enable or disable feedback forwarding for an identity.
	//
	// This setting determines what happens when an identity is used to send an email that results in a bounce or complaint event.
	//
	// When you enable feedback forwarding, Amazon Pinpoint sends you email notifications when bounce or complaint events occur. Amazon Pinpoint sends this notification to the address that you specified in the Return-Path header of the original email.
	//
	// When you disable feedback forwarding, Amazon Pinpoint sends notifications through other mechanisms, such as by notifying an Amazon SNS topic. You're required to have a method of tracking bounces and complaints. If you haven't set up another mechanism for receiving bounce or complaint notifications, Amazon Pinpoint sends an email notification when these events occur (even if this setting is disabled).
	FeedbackForwardingEnabled interface{} `field:"optional" json:"feedbackForwardingEnabled" yaml:"feedbackForwardingEnabled"`
	// Used to enable or disable the custom Mail-From domain configuration for an email identity.
	MailFromAttributes interface{} `field:"optional" json:"mailFromAttributes" yaml:"mailFromAttributes"`
	// An object that defines the tags (keys and values) that you want to associate with the email identity.
	Tags *[]*CfnIdentity_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

