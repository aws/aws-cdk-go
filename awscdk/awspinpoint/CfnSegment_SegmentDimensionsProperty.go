package awspinpoint


// Specifies the dimension settings for a segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//   var userAttributes interface{}
//
//   segmentDimensionsProperty := &SegmentDimensionsProperty{
//   	Attributes: attributes,
//   	Behavior: &BehaviorProperty{
//   		Recency: &RecencyProperty{
//   			Duration: jsii.String("duration"),
//   			RecencyType: jsii.String("recencyType"),
//   		},
//   	},
//   	Demographic: &DemographicProperty{
//   		AppVersion: &SetDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Channel: &SetDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		DeviceType: &SetDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Make: &SetDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Model: &SetDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Platform: &SetDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	Location: &LocationProperty{
//   		Country: &SetDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		GpsPoint: &GPSPointProperty{
//   			Coordinates: &CoordinatesProperty{
//   				Latitude: jsii.Number(123),
//   				Longitude: jsii.Number(123),
//   			},
//   			RangeInKilometers: jsii.Number(123),
//   		},
//   	},
//   	Metrics: metrics,
//   	UserAttributes: userAttributes,
//   }
//
type CfnSegment_SegmentDimensionsProperty struct {
	// One or more custom attributes to use as criteria for the segment.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The behavior-based criteria, such as how recently users have used your app, for the segment.
	Behavior interface{} `field:"optional" json:"behavior" yaml:"behavior"`
	// The demographic-based criteria, such as device platform, for the segment.
	Demographic interface{} `field:"optional" json:"demographic" yaml:"demographic"`
	// The location-based criteria, such as region or GPS coordinates, for the segment.
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// One or more custom metrics to use as criteria for the segment.
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// One or more custom user attributes to use as criteria for the segment.
	UserAttributes interface{} `field:"optional" json:"userAttributes" yaml:"userAttributes"`
}

