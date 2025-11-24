package mixinsawsmediaconnect


// Configuration settings for a merge router input that combines two input sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mergeRouterInputConfigurationProperty := &MergeRouterInputConfigurationProperty{
//   	MergeRecoveryWindowMilliseconds: jsii.Number(123),
//   	NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   	ProtocolConfigurations: []interface{}{
//   		&MergeRouterInputProtocolConfigurationProperty{
//   			Rist: &RistRouterInputConfigurationProperty{
//   				Port: jsii.Number(123),
//   				RecoveryLatencyMilliseconds: jsii.Number(123),
//   			},
//   			Rtp: &RtpRouterInputConfigurationProperty{
//   				ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputconfiguration.html
//
type CfnRouterInputPropsMixin_MergeRouterInputConfigurationProperty struct {
	// The time window in milliseconds for merging the two input sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputconfiguration.html#cfn-mediaconnect-routerinput-mergerouterinputconfiguration-mergerecoverywindowmilliseconds
	//
	MergeRecoveryWindowMilliseconds *float64 `field:"optional" json:"mergeRecoveryWindowMilliseconds" yaml:"mergeRecoveryWindowMilliseconds"`
	// The ARN of the network interface to use for this merge router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputconfiguration.html#cfn-mediaconnect-routerinput-mergerouterinputconfiguration-networkinterfacearn
	//
	NetworkInterfaceArn *string `field:"optional" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// A list of exactly two protocol configurations for the merge input sources.
	//
	// Both must use the same protocol type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputconfiguration.html#cfn-mediaconnect-routerinput-mergerouterinputconfiguration-protocolconfigurations
	//
	ProtocolConfigurations interface{} `field:"optional" json:"protocolConfigurations" yaml:"protocolConfigurations"`
}

