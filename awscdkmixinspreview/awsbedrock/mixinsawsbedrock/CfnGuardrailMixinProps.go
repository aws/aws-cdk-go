package mixinsawsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnGuardrailPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGuardrailMixinProps := &CfnGuardrailMixinProps{
//   	AutomatedReasoningPolicyConfig: &AutomatedReasoningPolicyConfigProperty{
//   		ConfidenceThreshold: jsii.Number(123),
//   		Policies: []*string{
//   			jsii.String("policies"),
//   		},
//   	},
//   	BlockedInputMessaging: jsii.String("blockedInputMessaging"),
//   	BlockedOutputsMessaging: jsii.String("blockedOutputsMessaging"),
//   	ContentPolicyConfig: &ContentPolicyConfigProperty{
//   		ContentFiltersTierConfig: &ContentFiltersTierConfigProperty{
//   			TierName: jsii.String("tierName"),
//   		},
//   		FiltersConfig: []interface{}{
//   			&ContentFilterConfigProperty{
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				InputModalities: []*string{
//   					jsii.String("inputModalities"),
//   				},
//   				InputStrength: jsii.String("inputStrength"),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				OutputModalities: []*string{
//   					jsii.String("outputModalities"),
//   				},
//   				OutputStrength: jsii.String("outputStrength"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ContextualGroundingPolicyConfig: &ContextualGroundingPolicyConfigProperty{
//   		FiltersConfig: []interface{}{
//   			&ContextualGroundingFilterConfigProperty{
//   				Action: jsii.String("action"),
//   				Enabled: jsii.Boolean(false),
//   				Threshold: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	CrossRegionConfig: &GuardrailCrossRegionConfigProperty{
//   		GuardrailProfileArn: jsii.String("guardrailProfileArn"),
//   	},
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Name: jsii.String("name"),
//   	SensitiveInformationPolicyConfig: &SensitiveInformationPolicyConfigProperty{
//   		PiiEntitiesConfig: []interface{}{
//   			&PiiEntityConfigProperty{
//   				Action: jsii.String("action"),
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		RegexesConfig: []interface{}{
//   			&RegexConfigProperty{
//   				Action: jsii.String("action"),
//   				Description: jsii.String("description"),
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TopicPolicyConfig: &TopicPolicyConfigProperty{
//   		TopicsConfig: []interface{}{
//   			&TopicConfigProperty{
//   				Definition: jsii.String("definition"),
//   				Examples: []*string{
//   					jsii.String("examples"),
//   				},
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		TopicsTierConfig: &TopicsTierConfigProperty{
//   			TierName: jsii.String("tierName"),
//   		},
//   	},
//   	WordPolicyConfig: &WordPolicyConfigProperty{
//   		ManagedWordListsConfig: []interface{}{
//   			&ManagedWordsConfigProperty{
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		WordsConfig: []interface{}{
//   			&WordConfigProperty{
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Text: jsii.String("text"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html
//
type CfnGuardrailMixinProps struct {
	// Configuration settings for integrating Automated Reasoning policies with Amazon Bedrock Guardrails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-automatedreasoningpolicyconfig
	//
	AutomatedReasoningPolicyConfig interface{} `field:"optional" json:"automatedReasoningPolicyConfig" yaml:"automatedReasoningPolicyConfig"`
	// The message to return when the guardrail blocks a prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-blockedinputmessaging
	//
	BlockedInputMessaging *string `field:"optional" json:"blockedInputMessaging" yaml:"blockedInputMessaging"`
	// The message to return when the guardrail blocks a model response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-blockedoutputsmessaging
	//
	BlockedOutputsMessaging *string `field:"optional" json:"blockedOutputsMessaging" yaml:"blockedOutputsMessaging"`
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
	// The ARN of the AWS  key that you use to encrypt the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The name of the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
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

