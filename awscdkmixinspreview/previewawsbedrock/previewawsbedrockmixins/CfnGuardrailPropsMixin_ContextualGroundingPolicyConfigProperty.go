package previewawsbedrockmixins


// The policy configuration details for the guardrails contextual grounding policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contextualGroundingPolicyConfigProperty := &ContextualGroundingPolicyConfigProperty{
//   	FiltersConfig: []interface{}{
//   		&ContextualGroundingFilterConfigProperty{
//   			Action: jsii.String("action"),
//   			Enabled: jsii.Boolean(false),
//   			Threshold: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingpolicyconfig.html
//
type CfnGuardrailPropsMixin_ContextualGroundingPolicyConfigProperty struct {
	// List of contextual grounding filter configs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingpolicyconfig.html#cfn-bedrock-guardrail-contextualgroundingpolicyconfig-filtersconfig
	//
	FiltersConfig interface{} `field:"optional" json:"filtersConfig" yaml:"filtersConfig"`
}

