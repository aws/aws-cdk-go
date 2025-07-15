package awsbedrock


// Contains details about topics that the guardrail should identify and deny.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicPolicyConfigProperty := &TopicPolicyConfigProperty{
//   	TopicsConfig: []interface{}{
//   		&TopicConfigProperty{
//   			Definition: jsii.String("definition"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Examples: []*string{
//   				jsii.String("examples"),
//   			},
//   			InputAction: jsii.String("inputAction"),
//   			InputEnabled: jsii.Boolean(false),
//   			OutputAction: jsii.String("outputAction"),
//   			OutputEnabled: jsii.Boolean(false),
//   		},
//   	},
//
//   	// the properties below are optional
//   	TopicsTierConfig: &TopicsTierConfigProperty{
//   		TierName: jsii.String("tierName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicpolicyconfig.html
//
type CfnGuardrail_TopicPolicyConfigProperty struct {
	// A list of policies related to topics that the guardrail should deny.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicpolicyconfig.html#cfn-bedrock-guardrail-topicpolicyconfig-topicsconfig
	//
	TopicsConfig interface{} `field:"required" json:"topicsConfig" yaml:"topicsConfig"`
	// Guardrail tier config for topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicpolicyconfig.html#cfn-bedrock-guardrail-topicpolicyconfig-topicstierconfig
	//
	TopicsTierConfig interface{} `field:"optional" json:"topicsTierConfig" yaml:"topicsTierConfig"`
}

