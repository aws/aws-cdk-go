package awschime


// Specifies the type of message that triggers a bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   invokedByProperty := &InvokedByProperty{
//   	StandardMessages: jsii.String("standardMessages"),
//   	TargetedMessages: jsii.String("targetedMessages"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-invokedby.html
//
type CfnAppInstanceBot_InvokedByProperty struct {
	// Sets standard messages as the bot trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-invokedby.html#cfn-chime-appinstancebot-invokedby-standardmessages
	//
	StandardMessages *string `field:"required" json:"standardMessages" yaml:"standardMessages"`
	// Sets targeted messages as the bot trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstancebot-invokedby.html#cfn-chime-appinstancebot-invokedby-targetedmessages
	//
	TargetedMessages *string `field:"required" json:"targetedMessages" yaml:"targetedMessages"`
}

