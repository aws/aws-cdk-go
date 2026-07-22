package awsmediaconnect


// Detects black frames in the router input's source content and reports them through a CloudWatch metric, an EventBridge event, and a router input message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blackFramesConfigurationProperty := &BlackFramesConfigurationProperty{
//   	State: jsii.String("state"),
//   	ThresholdSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-blackframesconfiguration.html
//
type CfnRouterInput_BlackFramesConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-blackframesconfiguration.html#cfn-mediaconnect-routerinput-blackframesconfiguration-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
	// The number of consecutive seconds of black frames that MediaConnect must detect before it reports an issue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-blackframesconfiguration.html#cfn-mediaconnect-routerinput-blackframesconfiguration-thresholdseconds
	//
	ThresholdSeconds *float64 `field:"required" json:"thresholdSeconds" yaml:"thresholdSeconds"`
}

