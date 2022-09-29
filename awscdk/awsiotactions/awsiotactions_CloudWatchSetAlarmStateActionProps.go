package awsiotactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Configuration properties of an action for CloudWatch alarm.
//
// Example:
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//
//   metric := cloudwatch.NewMetric(&metricProps{
//   	namespace: jsii.String("MyNamespace"),
//   	metricName: jsii.String("MyMetric"),
//   	dimensions: map[string]interface{}{
//   		"MyDimension": jsii.String("MyDimensionValue"),
//   	},
//   })
//   alarm := cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &alarmProps{
//   	metric: metric,
//   	threshold: jsii.Number(100),
//   	evaluationPeriods: jsii.Number(3),
//   	datapointsToAlarm: jsii.Number(2),
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewCloudWatchSetAlarmStateAction(alarm, &cloudWatchSetAlarmStateActionProps{
//   			reason: jsii.String("AWS Iot Rule action is triggered"),
//   			alarmStateToSet: cloudwatch.alarmState_ALARM,
//   		}),
//   	},
//   })
//
// Experimental.
type CloudWatchSetAlarmStateActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The value of the alarm state to set.
	// Experimental.
	AlarmStateToSet awscloudwatch.AlarmState `field:"required" json:"alarmStateToSet" yaml:"alarmStateToSet"`
	// The reason for the alarm change.
	// Experimental.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

