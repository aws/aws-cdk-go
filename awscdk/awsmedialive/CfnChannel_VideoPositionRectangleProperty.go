package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoPositionRectangleProperty := &VideoPositionRectangleProperty{
//   	Height: jsii.Number(123),
//   	Width: jsii.Number(123),
//   	X: jsii.Number(123),
//   	Y: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videopositionrectangle.html
//
type CfnChannel_VideoPositionRectangleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videopositionrectangle.html#cfn-medialive-channel-videopositionrectangle-height
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videopositionrectangle.html#cfn-medialive-channel-videopositionrectangle-width
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videopositionrectangle.html#cfn-medialive-channel-videopositionrectangle-x
	//
	X *float64 `field:"optional" json:"x" yaml:"x"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videopositionrectangle.html#cfn-medialive-channel-videopositionrectangle-y
	//
	Y *float64 `field:"optional" json:"y" yaml:"y"`
}

