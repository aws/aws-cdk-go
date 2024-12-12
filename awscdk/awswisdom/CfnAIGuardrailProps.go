package awswisdom


// Properties for defining a `CfnAIGuardrail`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAIGuardrailProps := &CfnAIGuardrailProps{
//   	AssistantId: jsii.String("assistantId"),
//   	BlockedInputMessaging: jsii.String("blockedInputMessaging"),
//   	BlockedOutputsMessaging: jsii.String("blockedOutputsMessaging"),
//
//   	// the properties below are optional
//   	ContentPolicyConfig: &AIGuardrailContentPolicyConfigProperty{
//   		FiltersConfig: []interface{}{
//   			&GuardrailContentFilterConfigProperty{
//   				InputStrength: jsii.String("inputStrength"),
//   				OutputStrength: jsii.String("outputStrength"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ContextualGroundingPolicyConfig: &AIGuardrailContextualGroundingPolicyConfigProperty{
//   		FiltersConfig: []interface{}{
//   			&GuardrailContextualGroundingFilterConfigProperty{
//   				Threshold: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	SensitiveInformationPolicyConfig: &AIGuardrailSensitiveInformationPolicyConfigProperty{
//   		PiiEntitiesConfig: []interface{}{
//   			&GuardrailPiiEntityConfigProperty{
//   				Action: jsii.String("action"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		RegexesConfig: []interface{}{
//   			&GuardrailRegexConfigProperty{
//   				Action: jsii.String("action"),
//   				Name: jsii.String("name"),
//   				Pattern: jsii.String("pattern"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TopicPolicyConfig: &AIGuardrailTopicPolicyConfigProperty{
//   		TopicsConfig: []interface{}{
//   			&GuardrailTopicConfigProperty{
//   				Definition: jsii.String("definition"),
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Examples: []*string{
//   					jsii.String("examples"),
//   				},
//   			},
//   		},
//   	},
//   	WordPolicyConfig: &AIGuardrailWordPolicyConfigProperty{
//   		ManagedWordListsConfig: []interface{}{
//   			&GuardrailManagedWordsConfigProperty{
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		WordsConfig: []interface{}{
//   			&GuardrailWordConfigProperty{
//   				Text: jsii.String("text"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html
//
type CfnAIGuardrailProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-assistantid
	//
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// Messaging for when violations are detected in text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-blockedinputmessaging
	//
	BlockedInputMessaging *string `field:"required" json:"blockedInputMessaging" yaml:"blockedInputMessaging"`
	// Messaging for when violations are detected in text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-blockedoutputsmessaging
	//
	BlockedOutputsMessaging *string `field:"required" json:"blockedOutputsMessaging" yaml:"blockedOutputsMessaging"`
	// Content policy config for a guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-contentpolicyconfig
	//
	ContentPolicyConfig interface{} `field:"optional" json:"contentPolicyConfig" yaml:"contentPolicyConfig"`
	// Contextual grounding policy config for a guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-contextualgroundingpolicyconfig
	//
	ContextualGroundingPolicyConfig interface{} `field:"optional" json:"contextualGroundingPolicyConfig" yaml:"contextualGroundingPolicyConfig"`
	// Description of the guardrail or its version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Sensitive information policy config for a guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-sensitiveinformationpolicyconfig
	//
	SensitiveInformationPolicyConfig interface{} `field:"optional" json:"sensitiveInformationPolicyConfig" yaml:"sensitiveInformationPolicyConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Topic policy config for a guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-topicpolicyconfig
	//
	TopicPolicyConfig interface{} `field:"optional" json:"topicPolicyConfig" yaml:"topicPolicyConfig"`
	// Word policy config for a guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrail.html#cfn-wisdom-aiguardrail-wordpolicyconfig
	//
	WordPolicyConfig interface{} `field:"optional" json:"wordPolicyConfig" yaml:"wordPolicyConfig"`
}

