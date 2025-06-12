package awsapplicationinsights


// The `AWS::ApplicationInsights::Application AlarmMetric` property type defines a metric to monitor for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmMetricProperty := &AlarmMetricProperty{
//   	AlarmMetricName: jsii.String("alarmMetricName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-alarmmetric.html
//
type CfnApplication_AlarmMetricProperty struct {
	// The name of the metric to be monitored for the component.
	//
	// For metrics supported by Application Insights, see [Logs and metrics supported by Amazon CloudWatch Application Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/appinsights-logs-and-metrics.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-alarmmetric.html#cfn-applicationinsights-application-alarmmetric-alarmmetricname
	//
	AlarmMetricName *string `field:"required" json:"alarmMetricName" yaml:"alarmMetricName"`
}

