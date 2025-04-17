package awsbedrock


// Contains details about how to handle harmful content.
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
//
//   			// the properties below are optional
//   			InputModalities: []*string{
//   				jsii.String("inputModalities"),
//   			},
//   			OutputModalities: []*string{
//   				jsii.String("outputModalities"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentpolicyconfig.html
//
type CfnGuardrail_ContentPolicyConfigProperty struct {
	// Contains the type of the content filter and how strongly it should apply to prompts and model responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentpolicyconfig.html#cfn-bedrock-guardrail-contentpolicyconfig-filtersconfig
	//
	FiltersConfig interface{} `field:"required" json:"filtersConfig" yaml:"filtersConfig"`
}

