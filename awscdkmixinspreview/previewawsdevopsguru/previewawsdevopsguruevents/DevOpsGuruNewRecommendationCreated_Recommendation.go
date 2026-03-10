package previewawsdevopsguruevents


// Type definition for Recommendation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recommendation := &Recommendation{
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	Link: []*string{
//   		jsii.String("link"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Reason: []*string{
//   		jsii.String("reason"),
//   	},
//   	RelatedAnomalies: []RelatedAnomaly{
//   		&RelatedAnomaly{
//   			AssociatedResourceArns: []*string{
//   				jsii.String("associatedResourceArns"),
//   			},
//   			Resources: []AnomalyResource{
//   				&AnomalyResource{
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Type: []*string{
//   						jsii.String("type"),
//   					},
//   				},
//   			},
//   			SourceDetails: []RelatedAnomalySourceDetail{
//   				&RelatedAnomalySourceDetail{
//   					CloudWatchMetrics: []CloudWatchDimension{
//   						&CloudWatchDimension{
//   							MetricName: []*string{
//   								jsii.String("metricName"),
//   							},
//   							Namespace: []*string{
//   								jsii.String("namespace"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RelatedEvents: []RelatedEvent{
//   		&RelatedEvent{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Resources: []EventResource{
//   				&EventResource{
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Type: []*string{
//   						jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewRecommendationCreated_Recommendation struct {
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// link property.
	//
	// Specify an array of string values to match this event if the actual value of link is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Link *[]*string `field:"optional" json:"link" yaml:"link"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// reason property.
	//
	// Specify an array of string values to match this event if the actual value of reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
	// relatedAnomalies property.
	//
	// Specify an array of string values to match this event if the actual value of relatedAnomalies is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RelatedAnomalies *[]*DevOpsGuruNewRecommendationCreated_RelatedAnomaly `field:"optional" json:"relatedAnomalies" yaml:"relatedAnomalies"`
	// relatedEvents property.
	//
	// Specify an array of string values to match this event if the actual value of relatedEvents is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RelatedEvents *[]*DevOpsGuruNewRecommendationCreated_RelatedEvent `field:"optional" json:"relatedEvents" yaml:"relatedEvents"`
}

