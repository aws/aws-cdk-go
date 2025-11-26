package previewawspinpointmixins


// Specifies the GPS coordinates of a location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   coordinatesProperty := &CoordinatesProperty{
//   	Latitude: jsii.Number(123),
//   	Longitude: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-coordinates.html
//
type CfnSegmentPropsMixin_CoordinatesProperty struct {
	// The latitude coordinate of the location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-coordinates.html#cfn-pinpoint-segment-coordinates-latitude
	//
	Latitude *float64 `field:"optional" json:"latitude" yaml:"latitude"`
	// The longitude coordinate of the location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-coordinates.html#cfn-pinpoint-segment-coordinates-longitude
	//
	Longitude *float64 `field:"optional" json:"longitude" yaml:"longitude"`
}

