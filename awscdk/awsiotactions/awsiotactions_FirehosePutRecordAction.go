package awsiotactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awsiotactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehose"
)

// The action to put the record from an MQTT message to the Kinesis Data Firehose stream.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &DeliveryStreamProps{
//   	Destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket),
//   	},
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewFirehosePutRecordAction(stream, &FirehosePutRecordActionProps{
//   			BatchMode: jsii.Boolean(true),
//   			RecordSeparator: actions.FirehoseRecordSeparator_NEWLINE,
//   		}),
//   	},
//   })
//
// Experimental.
type FirehosePutRecordAction interface {
	awsiot.IAction
	// Returns the topic rule action specification.
	// Experimental.
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for FirehosePutRecordAction
type jsiiProxy_FirehosePutRecordAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewFirehosePutRecordAction(stream awskinesisfirehose.IDeliveryStream, props *FirehosePutRecordActionProps) FirehosePutRecordAction {
	_init_.Initialize()

	if err := validateNewFirehosePutRecordActionParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehosePutRecordAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.FirehosePutRecordAction",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirehosePutRecordAction_Override(f FirehosePutRecordAction, stream awskinesisfirehose.IDeliveryStream, props *FirehosePutRecordActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.FirehosePutRecordAction",
		[]interface{}{stream, props},
		f,
	)
}

func (f *jsiiProxy_FirehosePutRecordAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	if err := f.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

