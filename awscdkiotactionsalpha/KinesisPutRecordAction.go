package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
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
	awscdkiotalpha.IAction
}

// The jsii proxy struct for KinesisPutRecordAction
type jsiiProxy_KinesisPutRecordAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewKinesisPutRecordAction(stream awskinesis.IStream, props *KinesisPutRecordActionProps) KinesisPutRecordAction {
	_init_.Initialize()

	if err := validateNewKinesisPutRecordActionParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisPutRecordAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.KinesisPutRecordAction",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisPutRecordAction_Override(k KinesisPutRecordAction, stream awskinesis.IStream, props *KinesisPutRecordActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.KinesisPutRecordAction",
		[]interface{}{stream, props},
		k,
	)
}

