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
//   cfnSegmentProps := &cfnSegmentProps{
//   	applicationId: jsii.String("applicationId"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dimensions: &segmentDimensionsProperty{
//   		attributes: attributes,
//   		behavior: &behaviorProperty{
//   			recency: &recencyProperty{
//   				duration: jsii.String("duration"),
//   				recencyType: jsii.String("recencyType"),
//   			},
//   		},
//   		demographic: &demographicProperty{
//   			appVersion: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			channel: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			deviceType: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			make: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			model: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			platform: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		location: &locationProperty{
//   			country: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			gpsPoint: &gPSPointProperty{
//   				coordinates: &coordinatesProperty{
//   					latitude: jsii.Number(123),
//   					longitude: jsii.Number(123),
//   				},
//   				rangeInKilometers: jsii.Number(123),
//   			},
//   		},
//   		metrics: metrics,
//   		userAttributes: userAttributes,
//   	},
//   	segmentGroups: &segmentGroupsProperty{
//   		groups: []interface{}{
//   			&groupsProperty{
//   				dimensions: []interface{}{
//   					&segmentDimensionsProperty{
//   						attributes: attributes,
//   						behavior: &behaviorProperty{
//   							recency: &recencyProperty{
//   								duration: jsii.String("duration"),
//   								recencyType: jsii.String("recencyType"),
//   							},
//   						},
//   						demographic: &demographicProperty{
//   							appVersion: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							channel: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							deviceType: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							make: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							model: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							platform: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   						},
//   						location: &locationProperty{
//   							country: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							gpsPoint: &gPSPointProperty{
//   								coordinates: &coordinatesProperty{
//   									latitude: jsii.Number(123),
//   									longitude: jsii.Number(123),
//   								},
//   								rangeInKilometers: jsii.Number(123),
//   							},
//   						},
//   						metrics: metrics,
//   						userAttributes: userAttributes,
//   					},
//   				},
//   				sourceSegments: []interface{}{
//   					&sourceSegmentsProperty{
//   						id: jsii.String("id"),
//
//   						// the properties below are optional
//   						version: jsii.Number(123),
//   					},
//   				},
//   				sourceType: jsii.String("sourceType"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		include: jsii.String("include"),
//   	},
//   	tags: tags,
//   }
//
type CfnSegmentProps struct {
	// The unique identifier for the Amazon Pinpoint application that the segment is associated with.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The name of the segment.
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

