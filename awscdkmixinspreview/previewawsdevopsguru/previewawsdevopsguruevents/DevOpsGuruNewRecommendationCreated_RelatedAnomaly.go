package previewawsdevopsguruevents


// Type definition for RelatedAnomaly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   relatedAnomaly := &RelatedAnomaly{
//   	AssociatedResourceArns: []*string{
//   		jsii.String("associatedResourceArns"),
//   	},
//   	Resources: []AnomalyResource{
//   		&AnomalyResource{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   	SourceDetails: []RelatedAnomalySourceDetail{
//   		&RelatedAnomalySourceDetail{
//   			CloudWatchMetrics: []CloudWatchDimension{
//   				&CloudWatchDimension{
//   					MetricName: []*string{
//   						jsii.String("metricName"),
//   					},
//   					Namespace: []*string{
//   						jsii.String("namespace"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewRecommendationCreated_RelatedAnomaly struct {
	// associatedResourceArns property.
	//
	// Specify an array of string values to match this event if the actual value of associatedResourceArns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AssociatedResourceArns *[]*string `field:"optional" json:"associatedResourceArns" yaml:"associatedResourceArns"`
	// resources property.
	//
	// Specify an array of string values to match this event if the actual value of resources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Resources *[]*DevOpsGuruNewRecommendationCreated_AnomalyResource `field:"optional" json:"resources" yaml:"resources"`
	// sourceDetails property.
	//
	// Specify an array of string values to match this event if the actual value of sourceDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceDetails *[]*DevOpsGuruNewRecommendationCreated_RelatedAnomalySourceDetail `field:"optional" json:"sourceDetails" yaml:"sourceDetails"`
}

