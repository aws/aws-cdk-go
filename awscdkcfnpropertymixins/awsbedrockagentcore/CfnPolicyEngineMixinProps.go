package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPolicyEnginePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPolicyEngineMixinProps := &CfnPolicyEngineMixinProps{
//   	Description: jsii.String("description"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policyengine.html
//
type CfnPolicyEngineMixinProps struct {
	// A human-readable description of the policy engine's purpose and scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policyengine.html#cfn-bedrockagentcore-policyengine-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the KMS key used to encrypt the policy engine data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policyengine.html#cfn-bedrockagentcore-policyengine-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The customer-assigned immutable name for the policy engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policyengine.html#cfn-bedrockagentcore-policyengine-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of tags to assign to the policy engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policyengine.html#cfn-bedrockagentcore-policyengine-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

