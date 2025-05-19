package awspinpoint


// Specifies the GPS coordinates of the endpoint location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gPSPointProperty := &GPSPointProperty{
//   	Coordinates: &CoordinatesProperty{
//   		Latitude: jsii.Number(123),
//   		Longitude: jsii.Number(123),
//   	},
//   	RangeInKilometers: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-gpspoint.html
//
type CfnSegment_GPSPointProperty struct {
	// The GPS coordinates to measure distance from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-gpspoint.html#cfn-pinpoint-segment-gpspoint-coordinates
	//
	Coordinates interface{} `field:"required" json:"coordinates" yaml:"coordinates"`
	// The range, in kilometers, from the GPS coordinates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-gpspoint.html#cfn-pinpoint-segment-gpspoint-rangeinkilometers
	//
	RangeInKilometers *float64 `field:"required" json:"rangeInKilometers" yaml:"rangeInKilometers"`
}

