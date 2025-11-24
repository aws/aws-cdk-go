package mixinsawsbedrock


// The tier that your guardrail uses for denied topic filters.
//
// Consider using a tier that balances performance, accuracy, and compatibility with your existing generative AI workflows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicsTierConfigProperty := &TopicsTierConfigProperty{
//   	TierName: jsii.String("tierName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicstierconfig.html
//
type CfnGuardrailPropsMixin_TopicsTierConfigProperty struct {
	// The tier that your guardrail uses for denied topic filters. Valid values include:.
	//
	// - `CLASSIC` tier – Provides established guardrails functionality supporting English, French, and Spanish languages.
	// - `STANDARD` tier – Provides a more robust solution than the `CLASSIC` tier and has more comprehensive language support. This tier requires that your guardrail use [cross-Region inference](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-cross-region.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicstierconfig.html#cfn-bedrock-guardrail-topicstierconfig-tiername
	//
	TierName *string `field:"optional" json:"tierName" yaml:"tierName"`
}

