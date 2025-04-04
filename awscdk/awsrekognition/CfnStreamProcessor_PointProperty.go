package awsrekognition


// The X and Y coordinates of a point on an image or video frame.
//
// The X and Y values are ratios of the overall image size or video resolution. For example, if the input image is 700x200 and the values are X=0.5 and Y=0.25, then the point is at the (350,50) pixel coordinate on the image.
//
// An array of `Point` objects, `Polygon` , is returned by [DetectText](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DetectText) and by [DetectCustomLabels](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DetectCustomLabels) or used to define regions of interest in Amazon Rekognition Video operations such as `CreateStreamProcessor` . `Polygon` represents a fine-grained polygon around a detected item. For more information, see [Geometry](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_Geometry) .
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
	// The value of the X coordinate for a point on a `Polygon` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-point.html#cfn-rekognition-streamprocessor-point-x
	//
	X *float64 `field:"required" json:"x" yaml:"x"`
	// The value of the Y coordinate for a point on a `Polygon` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-point.html#cfn-rekognition-streamprocessor-point-y
	//
	Y *float64 `field:"required" json:"y" yaml:"y"`
}

