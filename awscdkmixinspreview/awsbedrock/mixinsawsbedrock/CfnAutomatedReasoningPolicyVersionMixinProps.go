package mixinsawsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAutomatedReasoningPolicyVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAutomatedReasoningPolicyVersionMixinProps := &CfnAutomatedReasoningPolicyVersionMixinProps{
//   	LastUpdatedDefinitionHash: jsii.String("lastUpdatedDefinitionHash"),
//   	PolicyArn: jsii.String("policyArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-automatedreasoningpolicyversion.html
//
type CfnAutomatedReasoningPolicyVersionMixinProps struct {
	// The hash of the policy definition that was last updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-automatedreasoningpolicyversion.html#cfn-bedrock-automatedreasoningpolicyversion-lastupdateddefinitionhash
	//
	LastUpdatedDefinitionHash *string `field:"optional" json:"lastUpdatedDefinitionHash" yaml:"lastUpdatedDefinitionHash"`
	// The Amazon Resource Name (ARN) of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-automatedreasoningpolicyversion.html#cfn-bedrock-automatedreasoningpolicyversion-policyarn
	//
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
	// The tags associated with the Automated Reasoning policy version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-automatedreasoningpolicyversion.html#cfn-bedrock-automatedreasoningpolicyversion-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

