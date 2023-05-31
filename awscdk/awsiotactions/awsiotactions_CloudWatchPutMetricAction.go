package awsiotactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awsiotactions/internal"
)

// The action to capture an Amazon CloudWatch metric.
//
// Example:
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, namespace, unit, value, timestamp FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewCloudWatchPutMetricAction(&CloudWatchPutMetricActionProps{
//   			MetricName: jsii.String("${topic(2)}"),
//   			MetricNamespace: jsii.String("${namespace}"),
//   			MetricUnit: jsii.String("${unit}"),
//   			MetricValue: jsii.String("${value}"),
//   			MetricTimestamp: jsii.String("${timestamp}"),
//   		}),
//   	},
//   })
//
// Experimental.
type CloudWatchPutMetricAction interface {
	awsiot.IAction
	// Returns the topic rule action specification.
	// Experimental.
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for CloudWatchPutMetricAction
type jsiiProxy_CloudWatchPutMetricAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewCloudWatchPutMetricAction(props *CloudWatchPutMetricActionProps) CloudWatchPutMetricAction {
	_init_.Initialize()

	if err := validateNewCloudWatchPutMetricActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudWatchPutMetricAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchPutMetricAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchPutMetricAction_Override(c CloudWatchPutMetricAction, props *CloudWatchPutMetricActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchPutMetricAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudWatchPutMetricAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	if err := c.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

