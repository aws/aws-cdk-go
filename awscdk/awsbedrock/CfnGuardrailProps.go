package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGuardrail`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGuardrailProps := &CfnGuardrailProps{
//   	BlockedInputMessaging: jsii.String("blockedInputMessaging"),
//   	BlockedOutputsMessaging: jsii.String("blockedOutputsMessaging"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ContentPolicyConfig: &ContentPolicyConfigProperty{
//   		FiltersConfig: []interface{}{
//   			&ContentFilterConfigProperty{
//   				InputStrength: jsii.String("inputStrength"),
//   				OutputStrength: jsii.String("outputStrength"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				InputModalities: []*string{
//   					jsii.String("inputModalities"),
//   				},
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				OutputModalities: []*string{
//   					jsii.String("outputModalities"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		ContentFiltersTierConfig: &ContentFiltersTierConfigProperty{
//   			TierName: jsii.String("tierName"),
//   		},
//   	},
//   	ContextualGroundingPolicyConfig: &ContextualGroundingPolicyConfigProperty{
//   		FiltersConfig: []interface{}{
//   			&ContextualGroundingFilterConfigProperty{
//   				Threshold: jsii.Number(123),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Action: jsii.String("action"),
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	CrossRegionConfig: &GuardrailCrossRegionConfigProperty{
//   		GuardrailProfileArn: jsii.String("guardrailProfileArn"),
//   	},
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	SensitiveInformationPolicyConfig: &SensitiveInformationPolicyConfigProperty{
//   		PiiEntitiesConfig: []interface{}{
//   			&PiiEntityConfigProperty{
//   				Action: jsii.String("action"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   			},
//   		},
//   		RegexesConfig: []interface{}{
//   			&RegexConfigProperty{
//   				Action: jsii.String("action"),
//   				Name: jsii.String("name"),
//   				Pattern: jsii.String("pattern"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TopicPolicyConfig: &TopicPolicyConfigProperty{
//   		TopicsConfig: []interface{}{
//   			&TopicConfigProperty{
//   				Definition: jsii.String("definition"),
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Examples: []*string{
//   					jsii.String("examples"),
//   				},
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   			},
//   		},
//
//   		// the properties below are optional
//   		TopicsTierConfig: &TopicsTierConfigProperty{
//   			TierName: jsii.String("tierName"),
//   		},
//   	},
//   	WordPolicyConfig: &WordPolicyConfigProperty{
//   		ManagedWordListsConfig: []interface{}{
//   			&ManagedWordsConfigProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   			},
//   		},
//   		WordsConfig: []interface{}{
//   			&WordConfigProperty{
//   				Text: jsii.String("text"),
//
//   				// the properties below are optional
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
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
	// The content filter policies to configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-contentpolicyconfig
	//
	ContentPolicyConfig interface{} `field:"optional" json:"contentPolicyConfig" yaml:"contentPolicyConfig"`
	// Contextual grounding policy config for a guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html#cfn-bedrock-guardrail-contextualgroundingpolicyconfig
	//
	ContextualGroundingPolicyConfig interface{} `field:"optional" json:"contextualGroundingPolicyConfig" yaml:"contextualGroundingPolicyConfig"`
	// The system-defined guardrail profile that you’re using with your guardrail.
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

