package awsbedrock


// Properties for defining a `CfnGuardrailVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGuardrailVersionProps := &CfnGuardrailVersionProps{
//   	GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrailversion.html
//
type CfnGuardrailVersionProps struct {
	// The unique identifier of the guardrail.
	//
	// This can be an ID or the ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrailversion.html#cfn-bedrock-guardrailversion-guardrailidentifier
	//
	GuardrailIdentifier *string `field:"required" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// A description of the guardrail version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrailversion.html#cfn-bedrock-guardrailversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

