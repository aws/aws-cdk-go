package awsrekognition


// Identifies the bounding box around the label, face, text, or personal protective equipment.
//
// The `left` (x-coordinate) and `top` (y-coordinate) are coordinates representing the top and left sides of the bounding box. Note that the upper-left corner of the image is the origin (0,0).
//
// The `top` and `left` values returned are ratios of the overall image size. For example, if the input image is 700x200 pixels, and the top-left coordinate of the bounding box is 350x50 pixels, the API returns a `left` value of 0.5 (350/700) and a `top` value of 0.25 (50/200).
//
// The `width` and `height` values represent the dimensions of the bounding box as a ratio of the overall image dimension. For example, if the input image is 700x200 pixels, and the bounding box width is 70 pixels, the width returned is 0.1. For more information, see [BoundingBox](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_BoundingBox) .
//
// > The bounding box coordinates can have negative values. For example, if Amazon Rekognition is able to detect a face that is at the image edge and is only partially visible, the service can return coordinates that are outside the image bounds and, depending on the image edge, you might get negative values or values greater than 1 for the `left` or `top` values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   boundingBoxProperty := &boundingBoxProperty{
//   	height: jsii.Number(123),
//   	left: jsii.Number(123),
//   	top: jsii.Number(123),
//   	width: jsii.Number(123),
//   }
//
type CfnStreamProcessor_BoundingBoxProperty struct {
	// Height of the bounding box as a ratio of the overall image height.
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// Left coordinate of the bounding box as a ratio of overall image width.
	Left *float64 `field:"required" json:"left" yaml:"left"`
	// Top coordinate of the bounding box as a ratio of overall image height.
	Top *float64 `field:"required" json:"top" yaml:"top"`
	// Width of the bounding box as a ratio of the overall image width.
	Width *float64 `field:"required" json:"width" yaml:"width"`
}

