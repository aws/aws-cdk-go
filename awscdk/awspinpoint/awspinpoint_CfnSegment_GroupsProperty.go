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
//   groupsProperty := &groupsProperty{
//   	dimensions: []interface{}{
//   		&segmentDimensionsProperty{
//   			attributes: attributes,
//   			behavior: &behaviorProperty{
//   				recency: &recencyProperty{
//   					duration: jsii.String("duration"),
//   					recencyType: jsii.String("recencyType"),
//   				},
//   			},
//   			demographic: &demographicProperty{
//   				appVersion: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				channel: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				deviceType: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				make: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				model: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				platform: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			location: &locationProperty{
//   				country: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				gpsPoint: &gPSPointProperty{
//   					coordinates: &coordinatesProperty{
//   						latitude: jsii.Number(123),
//   						longitude: jsii.Number(123),
//   					},
//   					rangeInKilometers: jsii.Number(123),
//   				},
//   			},
//   			metrics: metrics,
//   			userAttributes: userAttributes,
//   		},
//   	},
//   	sourceSegments: []interface{}{
//   		&sourceSegmentsProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			version: jsii.Number(123),
//   		},
//   	},
//   	sourceType: jsii.String("sourceType"),
//   	type: jsii.String("type"),
//   }
//
type CfnSegment_GroupsProperty struct {
	// An array that defines the dimensions to include or exclude from the segment.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The base segment to build the segment on.
	//
	// A base segment, also called a *source segment* , defines the initial population of endpoints for a segment. When you add dimensions to the segment, Amazon Pinpoint filters the base segment by using the dimensions that you specify.
	//
	// You can specify more than one dimensional segment or only one imported segment. If you specify an imported segment, the segment size estimate that displays on the Amazon Pinpoint console indicates the size of the imported segment without any filters applied to it.
	SourceSegments interface{} `field:"optional" json:"sourceSegments" yaml:"sourceSegments"`
	// Specifies how to handle multiple base segments for the segment.
	//
	// For example, if you specify three base segments for the segment, whether the resulting segment is based on all, any, or none of the base segments.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Specifies how to handle multiple dimensions for the segment.
	//
	// For example, if you specify three dimensions for the segment, whether the resulting segment includes endpoints that match all, any, or none of the dimensions.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

