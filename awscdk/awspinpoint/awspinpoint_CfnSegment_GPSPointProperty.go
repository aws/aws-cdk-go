package awspinpoint


// Specifies the GPS coordinates of the endpoint location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gPSPointProperty := &gPSPointProperty{
//   	coordinates: &coordinatesProperty{
//   		latitude: jsii.Number(123),
//   		longitude: jsii.Number(123),
//   	},
//   	rangeInKilometers: jsii.Number(123),
//   }
//
type CfnSegment_GPSPointProperty struct {
	// The GPS coordinates to measure distance from.
	Coordinates interface{} `field:"required" json:"coordinates" yaml:"coordinates"`
	// The range, in kilometers, from the GPS coordinates.
	RangeInKilometers *float64 `field:"required" json:"rangeInKilometers" yaml:"rangeInKilometers"`
}

