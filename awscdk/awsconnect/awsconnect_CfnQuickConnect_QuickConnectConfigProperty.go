package awsconnect


// Contains configuration settings for a quick connect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quickConnectConfigProperty := &quickConnectConfigProperty{
//   	quickConnectType: jsii.String("quickConnectType"),
//
//   	// the properties below are optional
//   	phoneConfig: &phoneNumberQuickConnectConfigProperty{
//   		phoneNumber: jsii.String("phoneNumber"),
//   	},
//   	queueConfig: &queueQuickConnectConfigProperty{
//   		contactFlowArn: jsii.String("contactFlowArn"),
//   		queueArn: jsii.String("queueArn"),
//   	},
//   	userConfig: &userQuickConnectConfigProperty{
//   		contactFlowArn: jsii.String("contactFlowArn"),
//   		userArn: jsii.String("userArn"),
//   	},
//   }
//
type CfnQuickConnect_QuickConnectConfigProperty struct {
	// The type of quick connect.
	//
	// In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).
	QuickConnectType *string `field:"required" json:"quickConnectType" yaml:"quickConnectType"`
	// The phone configuration.
	//
	// This is required only if QuickConnectType is PHONE_NUMBER.
	PhoneConfig interface{} `field:"optional" json:"phoneConfig" yaml:"phoneConfig"`
	// The queue configuration.
	//
	// This is required only if QuickConnectType is QUEUE.
	QueueConfig interface{} `field:"optional" json:"queueConfig" yaml:"queueConfig"`
	// The user configuration.
	//
	// This is required only if QuickConnectType is USER.
	UserConfig interface{} `field:"optional" json:"userConfig" yaml:"userConfig"`
}

