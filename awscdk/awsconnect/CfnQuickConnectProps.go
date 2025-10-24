package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnQuickConnect`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQuickConnectProps := &CfnQuickConnectProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	QuickConnectConfig: &QuickConnectConfigProperty{
//   		QuickConnectType: jsii.String("quickConnectType"),
//
//   		// the properties below are optional
//   		PhoneConfig: &PhoneNumberQuickConnectConfigProperty{
//   			PhoneNumber: jsii.String("phoneNumber"),
//   		},
//   		QueueConfig: &QueueQuickConnectConfigProperty{
//   			ContactFlowArn: jsii.String("contactFlowArn"),
//   			QueueArn: jsii.String("queueArn"),
//   		},
//   		UserConfig: &UserQuickConnectConfigProperty{
//   			ContactFlowArn: jsii.String("contactFlowArn"),
//   			UserArn: jsii.String("userArn"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html
//
type CfnQuickConnectProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the quick connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains information about the quick connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-quickconnectconfig
	//
	QuickConnectConfig interface{} `field:"required" json:"quickConnectConfig" yaml:"quickConnectConfig"`
	// The description of the quick connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "Tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

