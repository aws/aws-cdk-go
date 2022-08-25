// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

// The action to put the message from an MQTT message to the IoT Events input.
//
// Example:
//   import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//   var role iRole
//
//
//   input := iotevents.NewInput(this, jsii.String("MyInput"), &inputProps{
//   	attributeJsonPaths: []*string{
//   		jsii.String("payload.temperature"),
//   		jsii.String("payload.transactionId"),
//   	},
//   })
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewIotEventsPutMessageAction(input, &iotEventsPutMessageActionProps{
//   			batchMode: jsii.Boolean(true),
//   			 // optional property, default is 'false'
//   			messageId: jsii.String("${payload.transactionId}"),
//   			 // optional property, default is a new UUID
//   			role: role,
//   		}),
//   	},
//   })
//
// Experimental.
type IotEventsPutMessageAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for IotEventsPutMessageAction
type jsiiProxy_IotEventsPutMessageAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewIotEventsPutMessageAction(input awscdkioteventsalpha.IInput, props *IotEventsPutMessageActionProps) IotEventsPutMessageAction {
	_init_.Initialize()

	j := jsiiProxy_IotEventsPutMessageAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.IotEventsPutMessageAction",
		[]interface{}{input, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIotEventsPutMessageAction_Override(i IotEventsPutMessageAction, input awscdkioteventsalpha.IInput, props *IotEventsPutMessageActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.IotEventsPutMessageAction",
		[]interface{}{input, props},
		i,
	)
}

