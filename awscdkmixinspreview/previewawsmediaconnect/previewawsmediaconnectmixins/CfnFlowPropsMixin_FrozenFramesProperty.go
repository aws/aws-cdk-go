package previewawsmediaconnectmixins


// Configures settings for the `FrozenFrames` metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   frozenFramesProperty := &FrozenFramesProperty{
//   	State: jsii.String("state"),
//   	ThresholdSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-frozenframes.html
//
type CfnFlowPropsMixin_FrozenFramesProperty struct {
	// Indicates whether the `FrozenFrames` metric is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-frozenframes.html#cfn-mediaconnect-flow-frozenframes-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// Specifies the number of consecutive seconds of a static image that triggers an event or alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-frozenframes.html#cfn-mediaconnect-flow-frozenframes-thresholdseconds
	//
	ThresholdSeconds *float64 `field:"optional" json:"thresholdSeconds" yaml:"thresholdSeconds"`
}

