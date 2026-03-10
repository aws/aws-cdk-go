package previewawsdevopsguruevents


// Type definition for RelatedAnomalySourceDetail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   relatedAnomalySourceDetail := &RelatedAnomalySourceDetail{
//   	CloudWatchMetrics: []CloudWatchDimension{
//   		&CloudWatchDimension{
//   			MetricName: []*string{
//   				jsii.String("metricName"),
//   			},
//   			Namespace: []*string{
//   				jsii.String("namespace"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewRecommendationCreated_RelatedAnomalySourceDetail struct {
	// cloudWatchMetrics property.
	//
	// Specify an array of string values to match this event if the actual value of cloudWatchMetrics is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CloudWatchMetrics *[]*DevOpsGuruNewRecommendationCreated_CloudWatchDimension `field:"optional" json:"cloudWatchMetrics" yaml:"cloudWatchMetrics"`
}

