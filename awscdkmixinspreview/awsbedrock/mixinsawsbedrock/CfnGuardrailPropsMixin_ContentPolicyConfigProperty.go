package mixinsawsbedrock


// Contains details about how to handle harmful content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contentPolicyConfigProperty := &ContentPolicyConfigProperty{
//   	ContentFiltersTierConfig: &ContentFiltersTierConfigProperty{
//   		TierName: jsii.String("tierName"),
//   	},
//   	FiltersConfig: []interface{}{
//   		&ContentFilterConfigProperty{
//   			InputAction: jsii.String("inputAction"),
//   			InputEnabled: jsii.Boolean(false),
//   			InputModalities: []*string{
//   				jsii.String("inputModalities"),
//   			},
//   			InputStrength: jsii.String("inputStrength"),
//   			OutputAction: jsii.String("outputAction"),
//   			OutputEnabled: jsii.Boolean(false),
//   			OutputModalities: []*string{
//   				jsii.String("outputModalities"),
//   			},
//   			OutputStrength: jsii.String("outputStrength"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentpolicyconfig.html
//
type CfnGuardrailPropsMixin_ContentPolicyConfigProperty struct {
	// The tier that your guardrail uses for content filters.
	//
	// Consider using a tier that balances performance, accuracy, and compatibility with your existing generative AI workflows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentpolicyconfig.html#cfn-bedrock-guardrail-contentpolicyconfig-contentfilterstierconfig
	//
	ContentFiltersTierConfig interface{} `field:"optional" json:"contentFiltersTierConfig" yaml:"contentFiltersTierConfig"`
	// Contains the type of the content filter and how strongly it should apply to prompts and model responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentpolicyconfig.html#cfn-bedrock-guardrail-contentpolicyconfig-filtersconfig
	//
	FiltersConfig interface{} `field:"optional" json:"filtersConfig" yaml:"filtersConfig"`
}

