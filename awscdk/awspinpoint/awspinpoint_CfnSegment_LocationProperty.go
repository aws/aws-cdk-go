package awspinpoint


// Specifies location-based criteria, such as region or GPS coordinates, for the segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &locationProperty{
//   	country: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	gpsPoint: &gPSPointProperty{
//   		coordinates: &coordinatesProperty{
//   			latitude: jsii.Number(123),
//   			longitude: jsii.Number(123),
//   		},
//   		rangeInKilometers: jsii.Number(123),
//   	},
//   }
//
type CfnSegment_LocationProperty struct {
	// The country or region code, in ISO 3166-1 alpha-2 format, for the segment.
	Country interface{} `field:"optional" json:"country" yaml:"country"`
	// The GPS point dimension for the segment.
	GpsPoint interface{} `field:"optional" json:"gpsPoint" yaml:"gpsPoint"`
}

