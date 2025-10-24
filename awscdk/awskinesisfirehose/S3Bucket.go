package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// An S3 bucket destination for data from an Amazon Data Firehose delivery stream.
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
//   	Actions: []IAction{
//   		actions.NewFirehosePutRecordAction(stream, &FirehosePutRecordActionProps{
//   			BatchMode: jsii.Boolean(true),
//   			RecordSeparator: actions.FirehoseRecordSeparator_NEWLINE,
//   		}),
//   	},
//   })
//
type S3Bucket interface {
	IDestination
	// Binds this destination to the Amazon Data Firehose delivery stream.
	//
	// Implementers should use this method to bind resources to the stack and initialize values using the provided stream.
	Bind(scope constructs.Construct, _options *DestinationBindOptions) *DestinationConfig
}

// The jsii proxy struct for S3Bucket
type jsiiProxy_S3Bucket struct {
	jsiiProxy_IDestination
}

func NewS3Bucket(bucket awss3.IBucket, props *S3BucketProps) S3Bucket {
	_init_.Initialize()

	if err := validateNewS3BucketParameters(bucket, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Bucket{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.S3Bucket",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

func NewS3Bucket_Override(s S3Bucket, bucket awss3.IBucket, props *S3BucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.S3Bucket",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3Bucket) Bind(scope constructs.Construct, _options *DestinationBindOptions) *DestinationConfig {
	if err := s.validateBindParameters(scope, _options); err != nil {
		panic(err)
	}
	var returns *DestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, _options},
		&returns,
	)

	return returns
}

