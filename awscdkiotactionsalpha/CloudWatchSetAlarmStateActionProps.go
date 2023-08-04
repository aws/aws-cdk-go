package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an action for CloudWatch alarm.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   metric := cloudwatch.NewMetric(&MetricProps{
//   	Namespace: jsii.String("MyNamespace"),
//   	MetricName: jsii.String("MyMetric"),
//   	Dimensions: map[string]interface{}{
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
//   	Actions: []iAction{
//   		actions.NewCloudWatchSetAlarmStateAction(alarm, &CloudWatchSetAlarmStateActionProps{
//   			Reason: jsii.String("AWS Iot Rule action is triggered"),
//   			AlarmStateToSet: cloudwatch.AlarmState_ALARM,
//   		}),
//   	},
//   })
//
// Experimental.
type CloudWatchSetAlarmStateActionProps struct {
	// The IAM role that allows access to AWS service.
	// Default: a new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The value of the alarm state to set.
	// Experimental.
	AlarmStateToSet awscloudwatch.AlarmState `field:"required" json:"alarmStateToSet" yaml:"alarmStateToSet"`
	// The reason for the alarm change.
	// Default: None.
	//
	// Experimental.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

