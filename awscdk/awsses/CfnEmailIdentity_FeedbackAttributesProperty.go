package awsses


// Used to enable or disable feedback forwarding for an identity.
//
// This setting determines what happens when an identity is used to send an email that results in a bounce or complaint event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   feedbackAttributesProperty := &FeedbackAttributesProperty{
//   	EmailForwardingEnabled: jsii.Boolean(false),
//   }
//
type CfnEmailIdentity_FeedbackAttributesProperty struct {
	// Sets the feedback forwarding configuration for the identity.
	//
	// If the value is `true` , you receive email notifications when bounce or complaint events occur. These notifications are sent to the address that you specified in the `Return-Path` header of the original email.
	//
	// You're required to have a method of tracking bounces and complaints. If you haven't set up another mechanism for receiving bounce or complaint notifications (for example, by setting up an event destination), you receive an email notification when these events occur (even if this setting is disabled).
	EmailForwardingEnabled interface{} `field:"optional" json:"emailForwardingEnabled" yaml:"emailForwardingEnabled"`
}

