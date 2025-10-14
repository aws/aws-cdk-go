package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGuardrail`.
//
// Example:
//   import bedrockl1 "github.com/aws/aws-cdk-go/awscdk"
//
//   // Import a guardrail created through the L1 CDK CfnGuardrail construct
//   l1guardrail := bedrockl1.NewCfnGuardrail(this, jsii.String("MyCfnGuardrail"), &CfnGuardrailProps{
//   	BlockedInputMessaging: jsii.String("blockedInputMessaging"),
//   	BlockedOutputsMessaging: jsii.String("blockedOutputsMessaging"),
//   	Name: jsii.String("namemycfnguardrails"),
//   	WordPolicyConfig: &WordPolicyConfigProperty{
//   		WordsConfig: []interface{}{
//   			&WordConfigProperty{
//   				Text: jsii.String("drugs"),
//   			},
//   		},
//   	},
//   })
//
//   importedGuardrail := bedrock.Guardrail_FromCfnGuardrail(l1guardrail)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html
//
type CfnGuardrailProps struct {
	// The message to return when the guardrail blocks a prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-blockedinputmessaging
	//
	BlockedInputMessaging *string `field:"required" json:"blockedInputMessaging" yaml:"blockedInputMessaging"`
	// The message to return when the guardrail blocks a model response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-blockedoutputsmessaging
	//
	BlockedOutputsMessaging *string `field:"required" json:"blockedOutputsMessaging" yaml:"blockedOutputsMessaging"`
	// The name of the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration settings for integrating Automated Reasoning policies with Amazon Bedrock Guardrails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-automatedreasoningpolicyconfig
	//
	AutomatedReasoningPolicyConfig interface{} `field:"optional" json:"automatedReasoningPolicyConfig" yaml:"automatedReasoningPolicyConfig"`
	// The content filter policies to configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-contentpolicyconfig
	//
	ContentPolicyConfig interface{} `field:"optional" json:"contentPolicyConfig" yaml:"contentPolicyConfig"`
	// Contextual grounding policy config for a guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-contextualgroundingpolicyconfig
	//
	ContextualGroundingPolicyConfig interface{} `field:"optional" json:"contextualGroundingPolicyConfig" yaml:"contextualGroundingPolicyConfig"`
	// The system-defined guardrail profile that you're using with your guardrail.
	//
	// Guardrail profiles define the destination AWS Regions where guardrail inference requests can be automatically routed. Using guardrail profiles helps maintain guardrail performance and reliability when demand increases.
	//
	// For more information, see the [Amazon Bedrock User Guide](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-cross-region.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-crossregionconfig
	//
	CrossRegionConfig interface{} `field:"optional" json:"crossRegionConfig" yaml:"crossRegionConfig"`
	// A description of the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the AWS KMS key that you use to encrypt the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The sensitive information policy to configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-sensitiveinformationpolicyconfig
	//
	SensitiveInformationPolicyConfig interface{} `field:"optional" json:"sensitiveInformationPolicyConfig" yaml:"sensitiveInformationPolicyConfig"`
	// The tags that you want to attach to the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The topic policies to configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-topicpolicyconfig
	//
	TopicPolicyConfig interface{} `field:"optional" json:"topicPolicyConfig" yaml:"topicPolicyConfig"`
	// The word policy you configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-wordpolicyconfig
	//
	WordPolicyConfig interface{} `field:"optional" json:"wordPolicyConfig" yaml:"wordPolicyConfig"`
}

