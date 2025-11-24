package mixinsawsbedrock


// The filter configuration details for the guardrails contextual grounding filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contextualGroundingFilterConfigProperty := &ContextualGroundingFilterConfigProperty{
//   	Action: jsii.String("action"),
//   	Enabled: jsii.Boolean(false),
//   	Threshold: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html
//
type CfnGuardrailPropsMixin_ContextualGroundingFilterConfigProperty struct {
	// Specifies the action to take when content fails the contextual grounding evaluation. Supported values include:.
	//
	// - `BLOCK` – Block the content and replace it with blocked messaging.
	// - `NONE` – Take no action but return detection information in the trace response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Specifies whether to enable contextual grounding evaluation.
	//
	// When disabled, you aren't charged for the evaluation. The evaluation doesn't appear in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The threshold details for the guardrails contextual grounding filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// The filter details for the guardrails contextual grounding filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

