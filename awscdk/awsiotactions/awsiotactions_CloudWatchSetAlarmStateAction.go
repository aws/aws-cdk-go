package awsiotactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awsiotactions/internal"
)

// The action to change the state of an Amazon CloudWatch alarm.
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
type CloudWatchSetAlarmStateAction interface {
	awsiot.IAction
	// Returns the topic rule action specification.
	// Experimental.
	Bind(topicRule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for CloudWatchSetAlarmStateAction
type jsiiProxy_CloudWatchSetAlarmStateAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewCloudWatchSetAlarmStateAction(alarm awscloudwatch.IAlarm, props *CloudWatchSetAlarmStateActionProps) CloudWatchSetAlarmStateAction {
	_init_.Initialize()

	if err := validateNewCloudWatchSetAlarmStateActionParameters(alarm, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudWatchSetAlarmStateAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchSetAlarmStateAction",
		[]interface{}{alarm, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchSetAlarmStateAction_Override(c CloudWatchSetAlarmStateAction, alarm awscloudwatch.IAlarm, props *CloudWatchSetAlarmStateActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchSetAlarmStateAction",
		[]interface{}{alarm, props},
		c,
	)
}

func (c *jsiiProxy_CloudWatchSetAlarmStateAction) Bind(topicRule awsiot.ITopicRule) *awsiot.ActionConfig {
	if err := c.validateBindParameters(topicRule); err != nil {
		panic(err)
	}
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{topicRule},
		&returns,
	)

	return returns
}

