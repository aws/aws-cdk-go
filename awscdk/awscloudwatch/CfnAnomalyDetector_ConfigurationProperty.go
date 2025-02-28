package awscloudwatch


// Specifies details about how the anomaly detection model is to be trained, including time ranges to exclude when training and updating the model.
//
// The configuration can also include the time zone to use for the metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationProperty := &ConfigurationProperty{
//   	ExcludedTimeRanges: []interface{}{
//   		&RangeProperty{
//   			EndTime: jsii.String("endTime"),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	MetricTimeZone: jsii.String("metricTimeZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-configuration.html
//
type CfnAnomalyDetector_ConfigurationProperty struct {
	// Specifies an array of time ranges to exclude from use when the anomaly detection model is trained and updated.
	//
	// Use this to make sure that events that could cause unusual values for the metric, such as deployments, aren't used when CloudWatch creates or updates the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-configuration.html#cfn-cloudwatch-anomalydetector-configuration-excludedtimeranges
	//
	ExcludedTimeRanges interface{} `field:"optional" json:"excludedTimeRanges" yaml:"excludedTimeRanges"`
	// The time zone to use for the metric.
	//
	// This is useful to enable the model to automatically account for daylight savings time changes if the metric is sensitive to such time changes.
	//
	// To specify a time zone, use the name of the time zone as specified in the standard tz database. For more information, see [tz database](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Tz_database) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-configuration.html#cfn-cloudwatch-anomalydetector-configuration-metrictimezone
	//
	MetricTimeZone *string `field:"optional" json:"metricTimeZone" yaml:"metricTimeZone"`
}

