package mixinsawspinpointemail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIdentityPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdentityMixinProps := &CfnIdentityMixinProps{
//   	DkimSigningEnabled: jsii.Boolean(false),
//   	FeedbackForwardingEnabled: jsii.Boolean(false),
//   	MailFromAttributes: &MailFromAttributesProperty{
//   		BehaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   		MailFromDomain: jsii.String("mailFromDomain"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html
//
type CfnIdentityMixinProps struct {
	// For domain identities, this attribute is used to enable or disable DomainKeys Identified Mail (DKIM) signing for the domain.
	//
	// If the value is `true` , then the messages that you send from the domain are signed using both the DKIM keys for your domain, as well as the keys for the `amazonses.com` domain. If the value is `false` , then the messages that you send are only signed using the DKIM keys for the `amazonses.com` domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-dkimsigningenabled
	//
	DkimSigningEnabled interface{} `field:"optional" json:"dkimSigningEnabled" yaml:"dkimSigningEnabled"`
	// Used to enable or disable feedback forwarding for an identity.
	//
	// This setting determines what happens when an identity is used to send an email that results in a bounce or complaint event.
	//
	// When you enable feedback forwarding, Amazon Pinpoint sends you email notifications when bounce or complaint events occur. Amazon Pinpoint sends this notification to the address that you specified in the Return-Path header of the original email.
	//
	// When you disable feedback forwarding, Amazon Pinpoint sends notifications through other mechanisms, such as by notifying an Amazon SNS topic. You're required to have a method of tracking bounces and complaints. If you haven't set up another mechanism for receiving bounce or complaint notifications, Amazon Pinpoint sends an email notification when these events occur (even if this setting is disabled).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-feedbackforwardingenabled
	//
	FeedbackForwardingEnabled interface{} `field:"optional" json:"feedbackForwardingEnabled" yaml:"feedbackForwardingEnabled"`
	// Used to enable or disable the custom Mail-From domain configuration for an email identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-mailfromattributes
	//
	MailFromAttributes interface{} `field:"optional" json:"mailFromAttributes" yaml:"mailFromAttributes"`
	// The address or domain of the identity, such as *sender@example.com* or *example.co.uk* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An object that defines the tags (keys and values) that you want to associate with the email identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

