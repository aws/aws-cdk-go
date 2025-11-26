package previewawswisdommixins


// A regex configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   guardrailRegexConfigProperty := &GuardrailRegexConfigProperty{
//   	Action: jsii.String("action"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Pattern: jsii.String("pattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html
//
type CfnAIGuardrailPropsMixin_GuardrailRegexConfigProperty struct {
	// The action of the guardrail regex configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html#cfn-wisdom-aiguardrail-guardrailregexconfig-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The regex description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html#cfn-wisdom-aiguardrail-guardrailregexconfig-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A regex configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html#cfn-wisdom-aiguardrail-guardrailregexconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The regex pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html#cfn-wisdom-aiguardrail-guardrailregexconfig-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

