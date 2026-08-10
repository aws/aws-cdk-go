package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerPrefixProperty := &ProviderPrefixProperty{
//   	Separator: jsii.String("separator"),
//   	Strip: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-providerprefix.html
//
type CfnGatewayTarget_ProviderPrefixProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-providerprefix.html#cfn-bedrockagentcore-gatewaytarget-providerprefix-separator
	//
	// Default: - "."
	//
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-providerprefix.html#cfn-bedrockagentcore-gatewaytarget-providerprefix-strip
	//
	// Default: - false.
	//
	Strip interface{} `field:"optional" json:"strip" yaml:"strip"`
}

