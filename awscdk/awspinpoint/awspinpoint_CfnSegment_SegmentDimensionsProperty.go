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
//   segmentDimensionsProperty := &segmentDimensionsProperty{
//   	attributes: attributes,
//   	behavior: &behaviorProperty{
//   		recency: &recencyProperty{
//   			duration: jsii.String("duration"),
//   			recencyType: jsii.String("recencyType"),
//   		},
//   	},
//   	demographic: &demographicProperty{
//   		appVersion: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		channel: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		deviceType: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		make: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		model: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		platform: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	location: &locationProperty{
//   		country: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		gpsPoint: &gPSPointProperty{
//   			coordinates: &coordinatesProperty{
//   				latitude: jsii.Number(123),
//   				longitude: jsii.Number(123),
//   			},
//   			rangeInKilometers: jsii.Number(123),
//   		},
//   	},
//   	metrics: metrics,
//   	userAttributes: userAttributes,
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

