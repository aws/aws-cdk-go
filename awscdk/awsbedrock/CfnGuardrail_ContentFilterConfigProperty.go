package awsbedrock


// Content filter config in content policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentFilterConfigProperty := &ContentFilterConfigProperty{
//   	InputStrength: jsii.String("inputStrength"),
//   	OutputStrength: jsii.String("outputStrength"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html
//
type CfnGuardrail_ContentFilterConfigProperty struct {
	// Strength for filters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-inputstrength
	//
	InputStrength *string `field:"required" json:"inputStrength" yaml:"inputStrength"`
	// Strength for filters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-outputstrength
	//
	OutputStrength *string `field:"required" json:"outputStrength" yaml:"outputStrength"`
	// Type of filter in content policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

