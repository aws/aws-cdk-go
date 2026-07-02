package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   stickinessConfigurationProperty := &StickinessConfigurationProperty{
//   	Identifier: jsii.String("identifier"),
//   	Timeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-stickinessconfiguration.html
//
type CfnGatewayTargetPropsMixin_StickinessConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-stickinessconfiguration.html#cfn-bedrockagentcore-gatewaytarget-stickinessconfiguration-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-stickinessconfiguration.html#cfn-bedrockagentcore-gatewaytarget-stickinessconfiguration-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

