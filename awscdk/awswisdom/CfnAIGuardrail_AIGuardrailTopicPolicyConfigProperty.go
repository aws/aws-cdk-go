package awswisdom


// Topic policy configuration for a guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIGuardrailTopicPolicyConfigProperty := &AIGuardrailTopicPolicyConfigProperty{
//   	TopicsConfig: []interface{}{
//   		&GuardrailTopicConfigProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailtopicpolicyconfig.html
//
type CfnAIGuardrail_AIGuardrailTopicPolicyConfigProperty struct {
	// List of topic configs in topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailtopicpolicyconfig.html#cfn-wisdom-aiguardrail-aiguardrailtopicpolicyconfig-topicsconfig
	//
	TopicsConfig interface{} `field:"required" json:"topicsConfig" yaml:"topicsConfig"`
}

