package awswisdom


// Content policy config for a guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIGuardrailContentPolicyConfigProperty := &AIGuardrailContentPolicyConfigProperty{
//   	FiltersConfig: []interface{}{
//   		&GuardrailContentFilterConfigProperty{
//   			InputStrength: jsii.String("inputStrength"),
//   			OutputStrength: jsii.String("outputStrength"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailcontentpolicyconfig.html
//
type CfnAIGuardrail_AIGuardrailContentPolicyConfigProperty struct {
	// List of content filter configs in content policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailcontentpolicyconfig.html#cfn-wisdom-aiguardrail-aiguardrailcontentpolicyconfig-filtersconfig
	//
	FiltersConfig interface{} `field:"required" json:"filtersConfig" yaml:"filtersConfig"`
}

