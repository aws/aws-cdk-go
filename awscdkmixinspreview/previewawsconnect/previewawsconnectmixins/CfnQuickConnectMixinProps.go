package previewawsconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnQuickConnectPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnQuickConnectMixinProps := &CfnQuickConnectMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	QuickConnectConfig: &QuickConnectConfigProperty{
//   		PhoneConfig: &PhoneNumberQuickConnectConfigProperty{
//   			PhoneNumber: jsii.String("phoneNumber"),
//   		},
//   		QueueConfig: &QueueQuickConnectConfigProperty{
//   			ContactFlowArn: jsii.String("contactFlowArn"),
//   			QueueArn: jsii.String("queueArn"),
//   		},
//   		QuickConnectType: jsii.String("quickConnectType"),
//   		UserConfig: &UserQuickConnectConfigProperty{
//   			ContactFlowArn: jsii.String("contactFlowArn"),
//   			UserArn: jsii.String("userArn"),
//   		},
//   	},
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
type CfnQuickConnectMixinProps struct {
	// The description of the quick connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The name of the quick connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Contains information about the quick connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-quickconnectconfig
	//
	QuickConnectConfig interface{} `field:"optional" json:"quickConnectConfig" yaml:"quickConnectConfig"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "Tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-quickconnect.html#cfn-connect-quickconnect-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

