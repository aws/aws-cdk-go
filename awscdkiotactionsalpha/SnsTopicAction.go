// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
)

// The action to write the data from an MQTT message to an Amazon SNS topic.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewSnsTopicAction(topic, &SnsTopicActionProps{
//   			MessageFormat: actions.SnsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/iot/latest/developerguide/sns-rule-action.html
//
// Experimental.
type SnsTopicAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for SnsTopicAction
type jsiiProxy_SnsTopicAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewSnsTopicAction(topic awssns.ITopic, props *SnsTopicActionProps) SnsTopicAction {
	_init_.Initialize()

	if err := validateNewSnsTopicActionParameters(topic, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsTopicAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.SnsTopicAction",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsTopicAction_Override(s SnsTopicAction, topic awssns.ITopic, props *SnsTopicActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.SnsTopicAction",
		[]interface{}{topic, props},
		s,
	)
}

