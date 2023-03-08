package awspinpoint


// Properties for defining a `CfnSegment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//   var tags interface{}
//   var userAttributes interface{}
//
//   cfnSegmentProps := &CfnSegmentProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Dimensions: &SegmentDimensionsProperty{
//   		Attributes: attributes,
//   		Behavior: &BehaviorProperty{
//   			Recency: &RecencyProperty{
//   				Duration: jsii.String("duration"),
//   				RecencyType: jsii.String("recencyType"),
//   			},
//   		},
//   		Demographic: &DemographicProperty{
//   			AppVersion: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Channel: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			DeviceType: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Make: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Model: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Platform: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Location: &LocationProperty{
//   			Country: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			GpsPoint: &GPSPointProperty{
//   				Coordinates: &CoordinatesProperty{
//   					Latitude: jsii.Number(123),
//   					Longitude: jsii.Number(123),
//   				},
//   				RangeInKilometers: jsii.Number(123),
//   			},
//   		},
//   		Metrics: metrics,
//   		UserAttributes: userAttributes,
//   	},
//   	SegmentGroups: &SegmentGroupsProperty{
//   		Groups: []interface{}{
//   			&GroupsProperty{
//   				Dimensions: []interface{}{
//   					&SegmentDimensionsProperty{
//   						Attributes: attributes,
//   						Behavior: &BehaviorProperty{
//   							Recency: &RecencyProperty{
//   								Duration: jsii.String("duration"),
//   								RecencyType: jsii.String("recencyType"),
//   							},
//   						},
//   						Demographic: &DemographicProperty{
//   							AppVersion: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Channel: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							DeviceType: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Make: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Model: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							Platform: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   						},
//   						Location: &LocationProperty{
//   							Country: &SetDimensionProperty{
//   								DimensionType: jsii.String("dimensionType"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							GpsPoint: &GPSPointProperty{
//   								Coordinates: &CoordinatesProperty{
//   									Latitude: jsii.Number(123),
//   									Longitude: jsii.Number(123),
//   								},
//   								RangeInKilometers: jsii.Number(123),
//   							},
//   						},
//   						Metrics: metrics,
//   						UserAttributes: userAttributes,
//   					},
//   				},
//   				SourceSegments: []interface{}{
//   					&SourceSegmentsProperty{
//   						Id: jsii.String("id"),
//
//   						// the properties below are optional
//   						Version: jsii.Number(123),
//   					},
//   				},
//   				SourceType: jsii.String("sourceType"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Include: jsii.String("include"),
//   	},
//   	Tags: tags,
//   }
//
type CfnSegmentProps struct {
	// The unique identifier for the Amazon Pinpoint application that the segment is associated with.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The name of the segment.
	//
	// > A segment must have a name otherwise it will not appear in the Amazon Pinpoint console.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The criteria that define the dimensions for the segment.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The segment group to use and the dimensions to apply to the group's base segments in order to build the segment.
	//
	// A segment group can consist of zero or more base segments. Your request can include only one segment group.
	SegmentGroups interface{} `field:"optional" json:"segmentGroups" yaml:"segmentGroups"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

