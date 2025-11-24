package mixinsawspinpoint


// Specifies location-based criteria, such as region or GPS coordinates, for the segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   locationProperty := &LocationProperty{
//   	Country: &SetDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	GpsPoint: &GPSPointProperty{
//   		Coordinates: &CoordinatesProperty{
//   			Latitude: jsii.Number(123),
//   			Longitude: jsii.Number(123),
//   		},
//   		RangeInKilometers: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-location.html
//
type CfnSegmentPropsMixin_LocationProperty struct {
	// The country or region code, in ISO 3166-1 alpha-2 format, for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-location.html#cfn-pinpoint-segment-location-country
	//
	Country interface{} `field:"optional" json:"country" yaml:"country"`
	// The GPS point dimension for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-location.html#cfn-pinpoint-segment-location-gpspoint
	//
	GpsPoint interface{} `field:"optional" json:"gpsPoint" yaml:"gpsPoint"`
}

