package awsbedrockagentcore


// Rate configuration for a metric (requests or tokens).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   rateConfigProperty := &RateConfigProperty{
//   	Period: jsii.String("period"),
//   	Rate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayratelimit-rateconfig.html
//
type CfnGatewayRateLimitPropsMixin_RateConfigProperty struct {
	// Time period for rate limiting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayratelimit-rateconfig.html#cfn-bedrockagentcore-gatewayratelimit-rateconfig-period
	//
	Period *string `field:"optional" json:"period" yaml:"period"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayratelimit-rateconfig.html#cfn-bedrockagentcore-gatewayratelimit-rateconfig-rate
	//
	Rate *float64 `field:"optional" json:"rate" yaml:"rate"`
}

