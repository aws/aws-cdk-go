package awsiotactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awsiotactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
)

// The action to put the record from an MQTT message to the Kinesis Data stream.
//
// Example:
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewKinesisPutRecordAction(stream, &KinesisPutRecordActionProps{
//   			PartitionKey: jsii.String("${newuuid()}"),
//   		}),
//   	},
//   })
//
// Experimental.
type KinesisPutRecordAction interface {
	awsiot.IAction
	// Returns the topic rule action specification.
	// Experimental.
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for KinesisPutRecordAction
type jsiiProxy_KinesisPutRecordAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewKinesisPutRecordAction(stream awskinesis.IStream, props *KinesisPutRecordActionProps) KinesisPutRecordAction {
	_init_.Initialize()

	if err := validateNewKinesisPutRecordActionParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisPutRecordAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.KinesisPutRecordAction",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisPutRecordAction_Override(k KinesisPutRecordAction, stream awskinesis.IStream, props *KinesisPutRecordActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.KinesisPutRecordAction",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisPutRecordAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	if err := k.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

