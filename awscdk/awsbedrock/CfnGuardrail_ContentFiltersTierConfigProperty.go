package awsbedrock


// Guardrail tier config for content policy.
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
	// Tier name for tier configuration in content filters policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterstierconfig.html#cfn-bedrock-guardrail-contentfilterstierconfig-tiername
	//
	TierName *string `field:"required" json:"tierName" yaml:"tierName"`
}

