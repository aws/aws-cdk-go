package awsbedrock


// The filter configuration details for the guardrails contextual grounding filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contextualGroundingFilterConfigProperty := &ContextualGroundingFilterConfigProperty{
//   	Threshold: jsii.Number(123),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Action: jsii.String("action"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html
//
type CfnGuardrail_ContextualGroundingFilterConfigProperty struct {
	// The threshold details for the guardrails contextual grounding filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-threshold
	//
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The filter details for the guardrails contextual grounding filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

