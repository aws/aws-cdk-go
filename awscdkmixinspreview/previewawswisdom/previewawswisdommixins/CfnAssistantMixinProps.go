package previewawswisdommixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAssistantPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssistantMixinProps := &CfnAssistantMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistant.html
//
type CfnAssistantMixinProps struct {
	// The description of the assistant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistant.html#cfn-wisdom-assistant-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the assistant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistant.html#cfn-wisdom-assistant-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The configuration information for the customer managed key used for encryption.
	//
	// The customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom. To use Wisdom with chat, the key policy must also allow `kms:Decrypt` , `kms:GenerateDataKey*` , and `kms:DescribeKey` permissions to the `connect.amazonaws.com` service principal. For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistant.html#cfn-wisdom-assistant-serversideencryptionconfiguration
	//
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistant.html#cfn-wisdom-assistant-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of assistant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistant.html#cfn-wisdom-assistant-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

