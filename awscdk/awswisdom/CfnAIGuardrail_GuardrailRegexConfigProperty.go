package awswisdom


// A regex configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardrailRegexConfigProperty := &GuardrailRegexConfigProperty{
//   	Action: jsii.String("action"),
//   	Name: jsii.String("name"),
//   	Pattern: jsii.String("pattern"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html
//
type CfnAIGuardrail_GuardrailRegexConfigProperty struct {
	// The action of the guardrail regex configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html#cfn-wisdom-aiguardrail-guardrailregexconfig-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// A regex configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html#cfn-wisdom-aiguardrail-guardrailregexconfig-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The regex pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html#cfn-wisdom-aiguardrail-guardrailregexconfig-pattern
	//
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The regex description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html#cfn-wisdom-aiguardrail-guardrailregexconfig-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

