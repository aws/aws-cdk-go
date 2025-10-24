package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
)

// The action to send data from an MQTT message to a web application or service.
//
// Example:
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   })
//
//   topicRule.AddAction(
//   actions.NewHttpsAction(jsii.String("https://example.com/endpoint"), &HttpsActionProps{
//   	ConfirmationUrl: jsii.String("https://example.com"),
//   	Headers: []HttpActionHeader{
//   		&HttpActionHeader{
//   			Key: jsii.String("key0"),
//   			Value: jsii.String("value0"),
//   		},
//   		&HttpActionHeader{
//   			Key: jsii.String("key1"),
//   			Value: jsii.String("value1"),
//   		},
//   	},
//   	Auth: &HttpActionSigV4Auth{
//   		ServiceName: jsii.String("serviceName"),
//   		SigningRegion: jsii.String("us-east-1"),
//   	},
//   }))
//
// Experimental.
type HttpsAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for HttpsAction
type jsiiProxy_HttpsAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewHttpsAction(url *string, props *HttpsActionProps) HttpsAction {
	_init_.Initialize()

	if err := validateNewHttpsActionParameters(url, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpsAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.HttpsAction",
		[]interface{}{url, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpsAction_Override(h HttpsAction, url *string, props *HttpsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.HttpsAction",
		[]interface{}{url, props},
		h,
	)
}

