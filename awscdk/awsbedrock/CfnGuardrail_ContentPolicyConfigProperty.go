package awsbedrock


// Content policy config for a guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentPolicyConfigProperty := &ContentPolicyConfigProperty{
//   	FiltersConfig: []interface{}{
//   		&ContentFilterConfigProperty{
//   			InputStrength: jsii.String("inputStrength"),
//   			OutputStrength: jsii.String("outputStrength"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentpolicyconfig.html
//
type CfnGuardrail_ContentPolicyConfigProperty struct {
	// List of content filter configs in content policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentpolicyconfig.html#cfn-bedrock-guardrail-contentpolicyconfig-filtersconfig
	//
	FiltersConfig interface{} `field:"required" json:"filtersConfig" yaml:"filtersConfig"`
}

