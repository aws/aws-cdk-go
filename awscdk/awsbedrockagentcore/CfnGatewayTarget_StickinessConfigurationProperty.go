package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stickinessConfigurationProperty := &StickinessConfigurationProperty{
//   	Identifier: jsii.String("identifier"),
//
//   	// the properties below are optional
//   	Timeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-stickinessconfiguration.html
//
type CfnGatewayTarget_StickinessConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-stickinessconfiguration.html#cfn-bedrockagentcore-gatewaytarget-stickinessconfiguration-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-stickinessconfiguration.html#cfn-bedrockagentcore-gatewaytarget-stickinessconfiguration-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

