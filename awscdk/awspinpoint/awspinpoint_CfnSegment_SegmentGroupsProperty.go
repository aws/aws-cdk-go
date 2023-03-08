package awspinpoint


// Specifies the set of segment criteria to evaluate when handling segment groups for the segment.
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
//   segmentGroupsProperty := &segmentGroupsProperty{
//   	groups: []interface{}{
//   		&groupsProperty{
//   			dimensions: []interface{}{
//   				&segmentDimensionsProperty{
//   					attributes: attributes,
//   					behavior: &behaviorProperty{
//   						recency: &recencyProperty{
//   							duration: jsii.String("duration"),
//   							recencyType: jsii.String("recencyType"),
//   						},
//   					},
//   					demographic: &demographicProperty{
//   						appVersion: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						channel: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						deviceType: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						make: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						model: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						platform: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   					location: &locationProperty{
//   						country: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						gpsPoint: &gPSPointProperty{
//   							coordinates: &coordinatesProperty{
//   								latitude: jsii.Number(123),
//   								longitude: jsii.Number(123),
//   							},
//   							rangeInKilometers: jsii.Number(123),
//   						},
//   					},
//   					metrics: metrics,
//   					userAttributes: userAttributes,
//   				},
//   			},
//   			sourceSegments: []interface{}{
//   				&sourceSegmentsProperty{
//   					id: jsii.String("id"),
//
//   					// the properties below are optional
//   					version: jsii.Number(123),
//   				},
//   			},
//   			sourceType: jsii.String("sourceType"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	include: jsii.String("include"),
//   }
//
type CfnSegment_SegmentGroupsProperty struct {
	// Specifies the set of segment criteria to evaluate when handling segment groups for the segment.
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// Specifies how to handle multiple segment groups for the segment.
	//
	// For example, if the segment includes three segment groups, whether the resulting segment includes endpoints that match all, any, or none of the segment groups.
	Include *string `field:"optional" json:"include" yaml:"include"`
}

