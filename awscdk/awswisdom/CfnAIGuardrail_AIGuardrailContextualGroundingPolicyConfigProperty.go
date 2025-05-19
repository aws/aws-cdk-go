package awswisdom


// Contextual grounding policy config for a guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIGuardrailContextualGroundingPolicyConfigProperty := &AIGuardrailContextualGroundingPolicyConfigProperty{
//   	FiltersConfig: []interface{}{
//   		&GuardrailContextualGroundingFilterConfigProperty{
//   			Threshold: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailcontextualgroundingpolicyconfig.html
//
type CfnAIGuardrail_AIGuardrailContextualGroundingPolicyConfigProperty struct {
	// List of contextual grounding filter configs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailcontextualgroundingpolicyconfig.html#cfn-wisdom-aiguardrail-aiguardrailcontextualgroundingpolicyconfig-filtersconfig
	//
	FiltersConfig interface{} `field:"required" json:"filtersConfig" yaml:"filtersConfig"`
}

