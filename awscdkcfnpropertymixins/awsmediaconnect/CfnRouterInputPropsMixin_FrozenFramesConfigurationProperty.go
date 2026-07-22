package awsmediaconnect


// Detects frozen video frames in the router input's source content and reports them through a CloudWatch metric, an EventBridge event, and a router input message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   frozenFramesConfigurationProperty := &FrozenFramesConfigurationProperty{
//   	State: jsii.String("state"),
//   	ThresholdSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-frozenframesconfiguration.html
//
type CfnRouterInputPropsMixin_FrozenFramesConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-frozenframesconfiguration.html#cfn-mediaconnect-routerinput-frozenframesconfiguration-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The number of consecutive seconds of a frozen frame that MediaConnect must detect before it reports an issue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-frozenframesconfiguration.html#cfn-mediaconnect-routerinput-frozenframesconfiguration-thresholdseconds
	//
	ThresholdSeconds *float64 `field:"optional" json:"thresholdSeconds" yaml:"thresholdSeconds"`
}

