package mixinsawswisdom


// A configuration for grounding filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   guardrailContextualGroundingFilterConfigProperty := &GuardrailContextualGroundingFilterConfigProperty{
//   	Threshold: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailcontextualgroundingfilterconfig.html
//
type CfnAIGuardrailPropsMixin_GuardrailContextualGroundingFilterConfigProperty struct {
	// The threshold for this filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailcontextualgroundingfilterconfig.html#cfn-wisdom-aiguardrail-guardrailcontextualgroundingfilterconfig-threshold
	//
	// Default: - 0.
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// The type of this filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailcontextualgroundingfilterconfig.html#cfn-wisdom-aiguardrail-guardrailcontextualgroundingfilterconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

