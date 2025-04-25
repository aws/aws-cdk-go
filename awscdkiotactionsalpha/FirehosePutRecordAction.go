package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
)

// The action to put the record from an MQTT message to the Amazon Data Firehose stream.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &DeliveryStreamProps{
//   	Destination: firehose.NewS3Bucket(bucket),
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
	awscdkiotalpha.IAction
}

// The jsii proxy struct for FirehosePutRecordAction
type jsiiProxy_FirehosePutRecordAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewFirehosePutRecordAction(stream awskinesisfirehose.IDeliveryStream, props *FirehosePutRecordActionProps) FirehosePutRecordAction {
	_init_.Initialize()

	if err := validateNewFirehosePutRecordActionParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehosePutRecordAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.FirehosePutRecordAction",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirehosePutRecordAction_Override(f FirehosePutRecordAction, stream awskinesisfirehose.IDeliveryStream, props *FirehosePutRecordActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.FirehosePutRecordAction",
		[]interface{}{stream, props},
		f,
	)
}

