package awsbedrock


// The policy configuration details for the guardrails contextual grounding policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contextualGroundingPolicyConfigProperty := &ContextualGroundingPolicyConfigProperty{
//   	FiltersConfig: []interface{}{
//   		&ContextualGroundingFilterConfigProperty{
//   			Threshold: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingpolicyconfig.html
//
type CfnGuardrail_ContextualGroundingPolicyConfigProperty struct {
	// List of contextual grounding filter configs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingpolicyconfig.html#cfn-bedrock-guardrail-contextualgroundingpolicyconfig-filtersconfig
	//
	FiltersConfig interface{} `field:"required" json:"filtersConfig" yaml:"filtersConfig"`
}

