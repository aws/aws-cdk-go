package awspinpoint


// An array that defines the set of segment criteria to evaluate when handling segment groups for the segment.
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
//   groupsProperty := &GroupsProperty{
//   	Dimensions: []interface{}{
//   		&SegmentDimensionsProperty{
//   			Attributes: attributes,
//   			Behavior: &BehaviorProperty{
//   				Recency: &RecencyProperty{
//   					Duration: jsii.String("duration"),
//   					RecencyType: jsii.String("recencyType"),
//   				},
//   			},
//   			Demographic: &DemographicProperty{
//   				AppVersion: &SetDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Channel: &SetDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				DeviceType: &SetDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Make: &SetDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Model: &SetDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Platform: &SetDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			Location: &LocationProperty{
//   				Country: &SetDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				GpsPoint: &GPSPointProperty{
//   					Coordinates: &CoordinatesProperty{
//   						Latitude: jsii.Number(123),
//   						Longitude: jsii.Number(123),
//   					},
//   					RangeInKilometers: jsii.Number(123),
//   				},
//   			},
//   			Metrics: metrics,
//   			UserAttributes: userAttributes,
//   		},
//   	},
//   	SourceSegments: []interface{}{
//   		&SourceSegmentsProperty{
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			Version: jsii.Number(123),
//   		},
//   	},
//   	SourceType: jsii.String("sourceType"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-groups.html
//
type CfnSegment_GroupsProperty struct {
	// An array that defines the dimensions to include or exclude from the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-groups.html#cfn-pinpoint-segment-groups-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The base segment to build the segment on.
	//
	// A base segment, also called a *source segment* , defines the initial population of endpoints for a segment. When you add dimensions to the segment, Amazon Pinpoint filters the base segment by using the dimensions that you specify.
	//
	// You can specify more than one dimensional segment or only one imported segment. If you specify an imported segment, the segment size estimate that displays on the Amazon Pinpoint console indicates the size of the imported segment without any filters applied to it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-groups.html#cfn-pinpoint-segment-groups-sourcesegments
	//
	SourceSegments interface{} `field:"optional" json:"sourceSegments" yaml:"sourceSegments"`
	// Specifies how to handle multiple base segments for the segment.
	//
	// For example, if you specify three base segments for the segment, whether the resulting segment is based on all, any, or none of the base segments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-groups.html#cfn-pinpoint-segment-groups-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Specifies how to handle multiple dimensions for the segment.
	//
	// For example, if you specify three dimensions for the segment, whether the resulting segment includes endpoints that match all, any, or none of the dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-groups.html#cfn-pinpoint-segment-groups-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

