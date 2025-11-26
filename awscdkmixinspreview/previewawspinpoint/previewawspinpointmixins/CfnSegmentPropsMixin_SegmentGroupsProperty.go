package previewawspinpointmixins


// Specifies the set of segment criteria to evaluate when handling segment groups for the segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var attributes interface{}
//   var metrics interface{}
//   var userAttributes interface{}
//
//   segmentGroupsProperty := &SegmentGroupsProperty{
//   	Groups: []interface{}{
//   		&GroupsProperty{
//   			Dimensions: []interface{}{
//   				&SegmentDimensionsProperty{
//   					Attributes: attributes,
//   					Behavior: &BehaviorProperty{
//   						Recency: &RecencyProperty{
//   							Duration: jsii.String("duration"),
//   							RecencyType: jsii.String("recencyType"),
//   						},
//   					},
//   					Demographic: &DemographicProperty{
//   						AppVersion: &SetDimensionProperty{
//   							DimensionType: jsii.String("dimensionType"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Channel: &SetDimensionProperty{
//   							DimensionType: jsii.String("dimensionType"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						DeviceType: &SetDimensionProperty{
//   							DimensionType: jsii.String("dimensionType"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Make: &SetDimensionProperty{
//   							DimensionType: jsii.String("dimensionType"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Model: &SetDimensionProperty{
//   							DimensionType: jsii.String("dimensionType"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Platform: &SetDimensionProperty{
//   							DimensionType: jsii.String("dimensionType"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   					Location: &LocationProperty{
//   						Country: &SetDimensionProperty{
//   							DimensionType: jsii.String("dimensionType"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						GpsPoint: &GPSPointProperty{
//   							Coordinates: &CoordinatesProperty{
//   								Latitude: jsii.Number(123),
//   								Longitude: jsii.Number(123),
//   							},
//   							RangeInKilometers: jsii.Number(123),
//   						},
//   					},
//   					Metrics: metrics,
//   					UserAttributes: userAttributes,
//   				},
//   			},
//   			SourceSegments: []interface{}{
//   				&SourceSegmentsProperty{
//   					Id: jsii.String("id"),
//   					Version: jsii.Number(123),
//   				},
//   			},
//   			SourceType: jsii.String("sourceType"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Include: jsii.String("include"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentgroups.html
//
type CfnSegmentPropsMixin_SegmentGroupsProperty struct {
	// Specifies the set of segment criteria to evaluate when handling segment groups for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentgroups.html#cfn-pinpoint-segment-segmentgroups-groups
	//
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// Specifies how to handle multiple segment groups for the segment.
	//
	// For example, if the segment includes three segment groups, whether the resulting segment includes endpoints that match all, any, or none of the segment groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentgroups.html#cfn-pinpoint-segment-segmentgroups-include
	//
	Include *string `field:"optional" json:"include" yaml:"include"`
}

