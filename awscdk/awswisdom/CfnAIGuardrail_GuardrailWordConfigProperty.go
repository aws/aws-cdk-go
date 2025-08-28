package awswisdom


// A custom word config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardrailWordConfigProperty := &GuardrailWordConfigProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailwordconfig.html
//
type CfnAIGuardrail_GuardrailWordConfigProperty struct {
	// The custom word text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailwordconfig.html#cfn-wisdom-aiguardrail-guardrailwordconfig-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
}

