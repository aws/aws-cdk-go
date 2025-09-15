package awsbedrock


// The tier that your guardrail uses for content filters.
//
// Consider using a tier that balances performance, accuracy, and compatibility with your existing generative AI workflows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentFiltersTierConfigProperty := &ContentFiltersTierConfigProperty{
//   	TierName: jsii.String("tierName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterstierconfig.html
//
type CfnGuardrail_ContentFiltersTierConfigProperty struct {
	// The tier that your guardrail uses for content filters. Valid values include:.
	//
	// - `CLASSIC` tier – Provides established guardrails functionality supporting English, French, and Spanish languages.
	// - `STANDARD` tier – Provides a more robust solution than the `CLASSIC` tier and has more comprehensive language support. This tier requires that your guardrail use [cross-Region inference](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-cross-region.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterstierconfig.html#cfn-bedrock-guardrail-contentfilterstierconfig-tiername
	//
	TierName *string `field:"required" json:"tierName" yaml:"tierName"`
}

