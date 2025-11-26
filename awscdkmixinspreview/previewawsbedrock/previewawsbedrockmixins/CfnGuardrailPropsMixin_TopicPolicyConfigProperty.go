package previewawsbedrockmixins


// Contains details about topics that the guardrail should identify and deny.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicPolicyConfigProperty := &TopicPolicyConfigProperty{
//   	TopicsConfig: []interface{}{
//   		&TopicConfigProperty{
//   			Definition: jsii.String("definition"),
//   			Examples: []*string{
//   				jsii.String("examples"),
//   			},
//   			InputAction: jsii.String("inputAction"),
//   			InputEnabled: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			OutputAction: jsii.String("outputAction"),
//   			OutputEnabled: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	TopicsTierConfig: &TopicsTierConfigProperty{
//   		TierName: jsii.String("tierName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicpolicyconfig.html
//
type CfnGuardrailPropsMixin_TopicPolicyConfigProperty struct {
	// A list of policies related to topics that the guardrail should deny.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicpolicyconfig.html#cfn-bedrock-guardrail-topicpolicyconfig-topicsconfig
	//
	TopicsConfig interface{} `field:"optional" json:"topicsConfig" yaml:"topicsConfig"`
	// The tier that your guardrail uses for denied topic filters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicpolicyconfig.html#cfn-bedrock-guardrail-topicpolicyconfig-topicstierconfig
	//
	TopicsTierConfig interface{} `field:"optional" json:"topicsTierConfig" yaml:"topicsTierConfig"`
}

