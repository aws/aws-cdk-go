package awsbedrock


// Topic policy config for a guardrail.
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicpolicyconfig.html
//
type CfnGuardrail_TopicPolicyConfigProperty struct {
	// List of topic configs in topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicpolicyconfig.html#cfn-bedrock-guardrail-topicpolicyconfig-topicsconfig
	//
	TopicsConfig interface{} `field:"required" json:"topicsConfig" yaml:"topicsConfig"`
}

