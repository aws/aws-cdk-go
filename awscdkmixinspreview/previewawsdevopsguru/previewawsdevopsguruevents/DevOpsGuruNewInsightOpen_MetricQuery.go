package previewawsdevopsguruevents


// Type definition for MetricQuery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricQuery := &MetricQuery{
//   	GroupBy: &GroupBy{
//   		Dimensions: []*string{
//   			jsii.String("dimensions"),
//   		},
//   		Group: []*string{
//   			jsii.String("group"),
//   		},
//   	},
//   	Metric: []*string{
//   		jsii.String("metric"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewInsightOpen_MetricQuery struct {
	// groupBy property.
	//
	// Specify an array of string values to match this event if the actual value of groupBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupBy *DevOpsGuruNewInsightOpen_GroupBy `field:"optional" json:"groupBy" yaml:"groupBy"`
	// metric property.
	//
	// Specify an array of string values to match this event if the actual value of metric is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Metric *[]*string `field:"optional" json:"metric" yaml:"metric"`
}

