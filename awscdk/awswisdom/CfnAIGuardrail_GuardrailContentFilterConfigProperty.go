package awswisdom


// Content filter configuration in content policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardrailContentFilterConfigProperty := &GuardrailContentFilterConfigProperty{
//   	InputStrength: jsii.String("inputStrength"),
//   	OutputStrength: jsii.String("outputStrength"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailcontentfilterconfig.html
//
type CfnAIGuardrail_GuardrailContentFilterConfigProperty struct {
	// The strength of the input for the guardrail content filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailcontentfilterconfig.html#cfn-wisdom-aiguardrail-guardrailcontentfilterconfig-inputstrength
	//
	InputStrength *string `field:"required" json:"inputStrength" yaml:"inputStrength"`
	// The output strength of the guardrail content filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailcontentfilterconfig.html#cfn-wisdom-aiguardrail-guardrailcontentfilterconfig-outputstrength
	//
	OutputStrength *string `field:"required" json:"outputStrength" yaml:"outputStrength"`
	// The type of the guardrail content filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailcontentfilterconfig.html#cfn-wisdom-aiguardrail-guardrailcontentfilterconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

