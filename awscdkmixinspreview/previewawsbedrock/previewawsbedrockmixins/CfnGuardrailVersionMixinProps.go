package previewawsbedrockmixins


// Properties for CfnGuardrailVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGuardrailVersionMixinProps := &CfnGuardrailVersionMixinProps{
//   	Description: jsii.String("description"),
//   	GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrailversion.html
//
type CfnGuardrailVersionMixinProps struct {
	// A description of the guardrail version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrailversion.html#cfn-bedrock-guardrailversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the guardrail.
	//
	// This can be an ID or the ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrailversion.html#cfn-bedrock-guardrailversion-guardrailidentifier
	//
	GuardrailIdentifier *string `field:"optional" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
}

