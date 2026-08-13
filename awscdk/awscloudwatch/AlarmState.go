package awscloudwatch


// Enumeration indicates state of Alarm used in building Alarm Rule.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   metric := cloudwatch.NewMetric(&MetricProps{
//   	Namespace: jsii.String("MyNamespace"),
//   	MetricName: jsii.String("MyMetric"),
//   	DimensionsMap: map[string]*string{
//   		"MyDimension": jsii.String("MyDimensionValue"),
//   	},
//   })
//   alarm := cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &AlarmProps{
//   	Metric: metric,
//   	Threshold: jsii.Number(100),
//   	EvaluationPeriods: jsii.Number(3),
//   	DatapointsToAlarm: jsii.Number(2),
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
//   	Actions: []IAction{
//   		actions.NewCloudWatchSetAlarmStateAction(alarm, &CloudWatchSetAlarmStateActionProps{
//   			Reason: jsii.String("AWS Iot Rule action is triggered"),
//   			AlarmStateToSet: cloudwatch.AlarmState_ALARM,
//   		}),
//   	},
//   })
//
type AlarmState string

const (
	// State indicates resource is in ALARM.
	AlarmState_ALARM AlarmState = "ALARM"
	// State indicates resource is not in ALARM.
	AlarmState_OK AlarmState = "OK"
	// State indicates there is not enough data to determine is resource is in ALARM.
	AlarmState_INSUFFICIENT_DATA AlarmState = "INSUFFICIENT_DATA"
)

