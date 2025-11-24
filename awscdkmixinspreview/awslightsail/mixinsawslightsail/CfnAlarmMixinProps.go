package mixinsawslightsail


// Properties for CfnAlarmPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAlarmMixinProps := &CfnAlarmMixinProps{
//   	AlarmName: jsii.String("alarmName"),
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	ContactProtocols: []*string{
//   		jsii.String("contactProtocols"),
//   	},
//   	DatapointsToAlarm: jsii.Number(123),
//   	EvaluationPeriods: jsii.Number(123),
//   	MetricName: jsii.String("metricName"),
//   	MonitoredResourceName: jsii.String("monitoredResourceName"),
//   	NotificationEnabled: jsii.Boolean(false),
//   	NotificationTriggers: []*string{
//   		jsii.String("notificationTriggers"),
//   	},
//   	Threshold: jsii.Number(123),
//   	TreatMissingData: jsii.String("treatMissingData"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html
//
type CfnAlarmMixinProps struct {
	// The name of the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-alarmname
	//
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The arithmetic operation to use when comparing the specified statistic and threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-comparisonoperator
	//
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The contact protocols for the alarm, such as `Email` , `SMS` (text messaging), or both.
	//
	// *Allowed Values* : `Email` | `SMS`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-contactprotocols
	//
	ContactProtocols *[]*string `field:"optional" json:"contactProtocols" yaml:"contactProtocols"`
	// The number of data points within the evaluation periods that must be breaching to cause the alarm to go to the `ALARM` state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-datapointstoalarm
	//
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// The number of periods over which data is compared to the specified threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-evaluationperiods
	//
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The name of the metric associated with the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The name of the Lightsail resource that the alarm monitors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-monitoredresourcename
	//
	MonitoredResourceName *string `field:"optional" json:"monitoredResourceName" yaml:"monitoredResourceName"`
	// A Boolean value indicating whether the alarm is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-notificationenabled
	//
	NotificationEnabled interface{} `field:"optional" json:"notificationEnabled" yaml:"notificationEnabled"`
	// The alarm states that trigger a notification.
	//
	// > To specify the `OK` and `INSUFFICIENT_DATA` values, you must also specify `ContactProtocols` values. Otherwise, the `OK` and `INSUFFICIENT_DATA` values will not take effect and the stack will drift.
	//
	// *Allowed Values* : `OK` | `ALARM` | `INSUFFICIENT_DATA`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-notificationtriggers
	//
	NotificationTriggers *[]*string `field:"optional" json:"notificationTriggers" yaml:"notificationTriggers"`
	// The value against which the specified statistic is compared.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Specifies how the alarm handles missing data points.
	//
	// An alarm can treat missing data in the following ways:
	//
	// - `breaching` - Assumes the missing data is not within the threshold. Missing data counts towards the number of times that the metric is not within the threshold.
	// - `notBreaching` - Assumes the missing data is within the threshold. Missing data does not count towards the number of times that the metric is not within the threshold.
	// - `ignore` - Ignores the missing data. Maintains the current alarm state.
	// - `missing` - Missing data is treated as missing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-treatmissingdata
	//
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

