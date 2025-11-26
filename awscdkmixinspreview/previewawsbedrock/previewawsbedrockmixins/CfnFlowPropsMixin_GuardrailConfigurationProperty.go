package previewawsbedrockmixins


// Configuration information for a guardrail that you use with the [Converse](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_Converse.html) operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   guardrailConfigurationProperty := &GuardrailConfigurationProperty{
//   	GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   	GuardrailVersion: jsii.String("guardrailVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-guardrailconfiguration.html
//
type CfnFlowPropsMixin_GuardrailConfigurationProperty struct {
	// The identifier for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-guardrailconfiguration.html#cfn-bedrock-flow-guardrailconfiguration-guardrailidentifier
	//
	GuardrailIdentifier *string `field:"optional" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// The version of the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-guardrailconfiguration.html#cfn-bedrock-flow-guardrailconfiguration-guardrailversion
	//
	GuardrailVersion *string `field:"optional" json:"guardrailVersion" yaml:"guardrailVersion"`
}

