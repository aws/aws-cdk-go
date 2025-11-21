package awsmediaconnect


// The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ristRouterInputConfigurationProperty := &RistRouterInputConfigurationProperty{
//   	Port: jsii.Number(123),
//   	RecoveryLatencyMilliseconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-ristrouterinputconfiguration.html
//
type CfnRouterInput_RistRouterInputConfigurationProperty struct {
	// The port number used for the RIST protocol in the router input configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-ristrouterinputconfiguration.html#cfn-mediaconnect-routerinput-ristrouterinputconfiguration-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The recovery latency in milliseconds for the RIST protocol in the router input configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-ristrouterinputconfiguration.html#cfn-mediaconnect-routerinput-ristrouterinputconfiguration-recoverylatencymilliseconds
	//
	RecoveryLatencyMilliseconds *float64 `field:"required" json:"recoveryLatencyMilliseconds" yaml:"recoveryLatencyMilliseconds"`
}

