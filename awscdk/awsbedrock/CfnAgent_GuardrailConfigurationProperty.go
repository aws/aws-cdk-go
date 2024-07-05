package awsbedrock


// Configuration information for a guardrail that you use with the `Converse` action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardrailConfigurationProperty := &GuardrailConfigurationProperty{
//   	GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   	GuardrailVersion: jsii.String("guardrailVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-guardrailconfiguration.html
//
type CfnAgent_GuardrailConfigurationProperty struct {
	// The identifier for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-guardrailconfiguration.html#cfn-bedrock-agent-guardrailconfiguration-guardrailidentifier
	//
	GuardrailIdentifier *string `field:"optional" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// The version of the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-guardrailconfiguration.html#cfn-bedrock-agent-guardrailconfiguration-guardrailversion
	//
	GuardrailVersion *string `field:"optional" json:"guardrailVersion" yaml:"guardrailVersion"`
}

