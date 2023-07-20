package awsrekognition


// An (X, Y) cartesian coordinate denoting a point on the frame.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pointProperty := &PointProperty{
//   	X: jsii.Number(123),
//   	Y: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-point.html
//
type CfnStreamProcessor_PointProperty struct {
	// The X coordinate of the point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-point.html#cfn-rekognition-streamprocessor-point-x
	//
	X *float64 `field:"required" json:"x" yaml:"x"`
	// The Y coordinate of the point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-point.html#cfn-rekognition-streamprocessor-point-y
	//
	Y *float64 `field:"required" json:"y" yaml:"y"`
}

