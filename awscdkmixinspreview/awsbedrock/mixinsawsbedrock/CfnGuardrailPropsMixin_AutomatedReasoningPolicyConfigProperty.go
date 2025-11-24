package mixinsawsbedrock


// Configuration settings for integrating Automated Reasoning policies with Amazon Bedrock Guardrails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   automatedReasoningPolicyConfigProperty := &AutomatedReasoningPolicyConfigProperty{
//   	ConfidenceThreshold: jsii.Number(123),
//   	Policies: []*string{
//   		jsii.String("policies"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-automatedreasoningpolicyconfig.html
//
type CfnGuardrailPropsMixin_AutomatedReasoningPolicyConfigProperty struct {
	// The minimum confidence level required for Automated Reasoning policy violations to trigger guardrail actions.
	//
	// Values range from 0.0 to 1.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-automatedreasoningpolicyconfig.html#cfn-bedrock-guardrail-automatedreasoningpolicyconfig-confidencethreshold
	//
	ConfidenceThreshold *float64 `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// The list of Automated Reasoning policy ARNs that should be applied as part of this guardrail configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-automatedreasoningpolicyconfig.html#cfn-bedrock-guardrail-automatedreasoningpolicyconfig-policies
	//
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

