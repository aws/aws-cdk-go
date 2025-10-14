package awsbedrock


// Configuration settings for integrating Automated Reasoning policies with Amazon Bedrock Guardrails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automatedReasoningPolicyConfigProperty := &AutomatedReasoningPolicyConfigProperty{
//   	Policies: []*string{
//   		jsii.String("policies"),
//   	},
//
//   	// the properties below are optional
//   	ConfidenceThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-automatedreasoningpolicyconfig.html
//
type CfnGuardrail_AutomatedReasoningPolicyConfigProperty struct {
	// The list of Automated Reasoning policy ARNs that should be applied as part of this guardrail configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-automatedreasoningpolicyconfig.html#cfn-bedrock-guardrail-automatedreasoningpolicyconfig-policies
	//
	Policies *[]*string `field:"required" json:"policies" yaml:"policies"`
	// The minimum confidence level required for Automated Reasoning policy violations to trigger guardrail actions.
	//
	// Values range from 0.0 to 1.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-automatedreasoningpolicyconfig.html#cfn-bedrock-guardrail-automatedreasoningpolicyconfig-confidencethreshold
	//
	ConfidenceThreshold *float64 `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
}

