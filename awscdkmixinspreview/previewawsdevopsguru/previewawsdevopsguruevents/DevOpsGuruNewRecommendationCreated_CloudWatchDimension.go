package previewawsdevopsguruevents


// Type definition for CloudWatchDimension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchDimension := &CloudWatchDimension{
//   	MetricName: []*string{
//   		jsii.String("metricName"),
//   	},
//   	Namespace: []*string{
//   		jsii.String("namespace"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewRecommendationCreated_CloudWatchDimension struct {
	// metricName property.
	//
	// Specify an array of string values to match this event if the actual value of metricName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MetricName *[]*string `field:"optional" json:"metricName" yaml:"metricName"`
	// namespace property.
	//
	// Specify an array of string values to match this event if the actual value of namespace is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Namespace *[]*string `field:"optional" json:"namespace" yaml:"namespace"`
}

