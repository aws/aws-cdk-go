package awslightsail


// Properties for defining a `CfnAlarm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmProps := &cfnAlarmProps{
//   	alarmName: jsii.String("alarmName"),
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	evaluationPeriods: jsii.Number(123),
//   	metricName: jsii.String("metricName"),
//   	monitoredResourceName: jsii.String("monitoredResourceName"),
//   	threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	contactProtocols: []*string{
//   		jsii.String("contactProtocols"),
//   	},
//   	datapointsToAlarm: jsii.Number(123),
//   	notificationEnabled: jsii.Boolean(false),
//   	notificationTriggers: []*string{
//   		jsii.String("notificationTriggers"),
//   	},
//   	treatMissingData: jsii.String("treatMissingData"),
//   }
//
type CfnAlarmProps struct {
	// The name of the alarm.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// The arithmetic operation to use when comparing the specified statistic and threshold.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The name of the metric associated with the alarm.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The name of the Lightsail resource that the alarm monitors.
	MonitoredResourceName *string `field:"required" json:"monitoredResourceName" yaml:"monitoredResourceName"`
	// The value against which the specified statistic is compared.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The contact protocols for the alarm, such as `Email` , `SMS` (text messaging), or both.
	//
	// *Allowed Values* : `Email` | `SMS`.
	ContactProtocols *[]*string `field:"optional" json:"contactProtocols" yaml:"contactProtocols"`
	// The number of data points within the evaluation periods that must be breaching to cause the alarm to go to the `ALARM` state.
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// A Boolean value indicating whether the alarm is enabled.
	NotificationEnabled interface{} `field:"optional" json:"notificationEnabled" yaml:"notificationEnabled"`
	// The alarm states that trigger a notification.
	//
	// > To specify the `OK` and `INSUFFICIENT_DATA` values, you must also specify `ContactProtocols` values. Otherwise, the `OK` and `INSUFFICIENT_DATA` values will not take effect and the stack will drift.
	//
	// *Allowed Values* : `OK` | `ALARM` | `INSUFFICIENT_DATA`.
	NotificationTriggers *[]*string `field:"optional" json:"notificationTriggers" yaml:"notificationTriggers"`
	// Specifies how the alarm handles missing data points.
	//
	// An alarm can treat missing data in the following ways:
	//
	// - `breaching` - Assumes the missing data is not within the threshold. Missing data counts towards the number of times that the metric is not within the threshold.
	// - `notBreaching` - Assumes the missing data is within the threshold. Missing data does not count towards the number of times that the metric is not within the threshold.
	// - `ignore` - Ignores the missing data. Maintains the current alarm state.
	// - `missing` - Missing data is treated as missing.
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

