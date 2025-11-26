package previewawscloudwatchevents


// Type definition for MetricStat.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricStat := &MetricStat{
//   	Metric: &Metric{
//   		Dimensions: []*string{
//   			jsii.String("dimensions"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		Namespace: []*string{
//   			jsii.String("namespace"),
//   		},
//   	},
//   	Period: []*string{
//   		jsii.String("period"),
//   	},
//   	Stat: []*string{
//   		jsii.String("stat"),
//   	},
//   }
//
// Experimental.
type AlarmEvents_CloudWatchAlarmStateChange_MetricStat struct {
	// metric property.
	//
	// Specify an array of string values to match this event if the actual value of metric is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Metric *AlarmEvents_CloudWatchAlarmStateChange_Metric `field:"optional" json:"metric" yaml:"metric"`
	// period property.
	//
	// Specify an array of string values to match this event if the actual value of period is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Period *[]*string `field:"optional" json:"period" yaml:"period"`
	// stat property.
	//
	// Specify an array of string values to match this event if the actual value of stat is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Stat *[]*string `field:"optional" json:"stat" yaml:"stat"`
}

