package awssagemaker


// The configuration for prefix-aware routing on a SageMaker real-time inference endpoint.
//
// Specify PrefixLength and ConcurrencyThreshold to control routing behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixAwareRoutingConfigProperty := &PrefixAwareRoutingConfigProperty{
//   	ConcurrencyThreshold: jsii.Number(123),
//   	PrefixLength: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-prefixawareroutingconfig.html
//
type CfnEndpointConfig_PrefixAwareRoutingConfigProperty struct {
	// The maximum number of in-flight requests on the target instance before the endpoint routes to another instance.
	//
	// Required when RoutingStrategy is PREFIX_AWARE. Valid values are 1 to 1024.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-prefixawareroutingconfig.html#cfn-sagemaker-endpointconfig-prefixawareroutingconfig-concurrencythreshold
	//
	ConcurrencyThreshold *float64 `field:"optional" json:"concurrencyThreshold" yaml:"concurrencyThreshold"`
	// The maximum length of the prefix used for routing decisions.
	//
	// Required when RoutingStrategy is PREFIX_AWARE. Valid values are 1024 to 65536.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-prefixawareroutingconfig.html#cfn-sagemaker-endpointconfig-prefixawareroutingconfig-prefixlength
	//
	PrefixLength *float64 `field:"optional" json:"prefixLength" yaml:"prefixLength"`
}

