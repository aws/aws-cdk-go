package awsconnect


// Contains configuration settings for a quick connect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quickConnectConfigProperty := &QuickConnectConfigProperty{
//   	QuickConnectType: jsii.String("quickConnectType"),
//
//   	// the properties below are optional
//   	PhoneConfig: &PhoneNumberQuickConnectConfigProperty{
//   		PhoneNumber: jsii.String("phoneNumber"),
//   	},
//   	QueueConfig: &QueueQuickConnectConfigProperty{
//   		ContactFlowArn: jsii.String("contactFlowArn"),
//   		QueueArn: jsii.String("queueArn"),
//   	},
//   	UserConfig: &UserQuickConnectConfigProperty{
//   		ContactFlowArn: jsii.String("contactFlowArn"),
//   		UserArn: jsii.String("userArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html
//
type CfnQuickConnect_QuickConnectConfigProperty struct {
	// The type of quick connect.
	//
	// In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-quickconnecttype
	//
	QuickConnectType *string `field:"required" json:"quickConnectType" yaml:"quickConnectType"`
	// The phone configuration.
	//
	// This is required only if QuickConnectType is PHONE_NUMBER.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-phoneconfig
	//
	PhoneConfig interface{} `field:"optional" json:"phoneConfig" yaml:"phoneConfig"`
	// The queue configuration.
	//
	// This is required only if QuickConnectType is QUEUE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-queueconfig
	//
	QueueConfig interface{} `field:"optional" json:"queueConfig" yaml:"queueConfig"`
	// The user configuration.
	//
	// This is required only if QuickConnectType is USER.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-userconfig
	//
	UserConfig interface{} `field:"optional" json:"userConfig" yaml:"userConfig"`
}

