package previewawsbedrockmixins


// The regular expression to configure for the guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   regexConfigProperty := &RegexConfigProperty{
//   	Action: jsii.String("action"),
//   	Description: jsii.String("description"),
//   	InputAction: jsii.String("inputAction"),
//   	InputEnabled: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	OutputAction: jsii.String("outputAction"),
//   	OutputEnabled: jsii.Boolean(false),
//   	Pattern: jsii.String("pattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-regexconfig.html
//
type CfnGuardrailPropsMixin_RegexConfigProperty struct {
	// The guardrail action to configure when matching regular expression is detected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-regexconfig.html#cfn-bedrock-guardrail-regexconfig-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The description of the regular expression to configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-regexconfig.html#cfn-bedrock-guardrail-regexconfig-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the action to take when harmful content is detected in the input. Supported values include:.
	//
	// - `BLOCK` – Block the content and replace it with blocked messaging.
	// - `NONE` – Take no action but return detection information in the trace response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-regexconfig.html#cfn-bedrock-guardrail-regexconfig-inputaction
	//
	InputAction *string `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Specifies whether to enable guardrail evaluation on the input.
	//
	// When disabled, you aren't charged for the evaluation. The evaluation doesn't appear in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-regexconfig.html#cfn-bedrock-guardrail-regexconfig-inputenabled
	//
	InputEnabled interface{} `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The name of the regular expression to configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-regexconfig.html#cfn-bedrock-guardrail-regexconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the action to take when harmful content is detected in the output. Supported values include:.
	//
	// - `BLOCK` – Block the content and replace it with blocked messaging.
	// - `NONE` – Take no action but return detection information in the trace response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-regexconfig.html#cfn-bedrock-guardrail-regexconfig-outputaction
	//
	OutputAction *string `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Specifies whether to enable guardrail evaluation on the output.
	//
	// When disabled, you aren't charged for the evaluation. The evaluation doesn't appear in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-regexconfig.html#cfn-bedrock-guardrail-regexconfig-outputenabled
	//
	OutputEnabled interface{} `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
	// The regular expression pattern to configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-regexconfig.html#cfn-bedrock-guardrail-regexconfig-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

