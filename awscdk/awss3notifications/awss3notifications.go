package awss3notifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3notifications/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Lambda function as a bucket notification destination.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myLambda function
//   bucket := s3.bucket.fromBucketAttributes(this, jsii.String("ImportedBucket"), &bucketAttributes{
//   	bucketArn: jsii.String("arn:aws:s3:::my-bucket"),
//   })
//
//   // now you can just call methods on the bucket
//   bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &notificationKeyFilter{
//   	prefix: jsii.String("home/myusername/*"),
//   })
//
type LambdaDestination interface {
	awss3.IBucketNotificationDestination
	// Registers this resource to receive notifications for the specified bucket.
	//
	// This method will only be called once for each destination/bucket
	// pair and the result will be cached, so there is no need to implement
	// idempotency in each destination.
	Bind(_scope constructs.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

func NewLambdaDestination(fn awslambda.IFunction) LambdaDestination {
	_init_.Initialize()

	j := jsiiProxy_LambdaDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_notifications.LambdaDestination",
		[]interface{}{fn},
		&j,
	)

	return &j
}

func NewLambdaDestination_Override(l LambdaDestination, fn awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_notifications.LambdaDestination",
		[]interface{}{fn},
		l,
	)
}

func (l *jsiiProxy_LambdaDestination) Bind(_scope constructs.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig {
	var returns *awss3.BucketNotificationDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, bucket},
		&returns,
	)

	return returns
}

// Use an SNS topic as a bucket notification destination.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//   bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewSnsDestination(topic))
//
type SnsDestination interface {
	awss3.IBucketNotificationDestination
	// Registers this resource to receive notifications for the specified bucket.
	//
	// This method will only be called once for each destination/bucket
	// pair and the result will be cached, so there is no need to implement
	// idempotency in each destination.
	Bind(_scope constructs.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for SnsDestination
type jsiiProxy_SnsDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

func NewSnsDestination(topic awssns.ITopic) SnsDestination {
	_init_.Initialize()

	j := jsiiProxy_SnsDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_notifications.SnsDestination",
		[]interface{}{topic},
		&j,
	)

	return &j
}

func NewSnsDestination_Override(s SnsDestination, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_notifications.SnsDestination",
		[]interface{}{topic},
		s,
	)
}

func (s *jsiiProxy_SnsDestination) Bind(_scope constructs.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig {
	var returns *awss3.BucketNotificationDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, bucket},
		&returns,
	)

	return returns
}

// Use an SQS queue as a bucket notification destination.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myQueue queue
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   bucket.addEventNotification(s3.eventType_OBJECT_REMOVED,
//   s3n.NewSqsDestination(myQueue), &notificationKeyFilter{
//   	prefix: jsii.String("foo/"),
//   	suffix: jsii.String(".jpg"),
//   })
//
type SqsDestination interface {
	awss3.IBucketNotificationDestination
	// Allows using SQS queues as destinations for bucket notifications.
	//
	// Use `bucket.onEvent(event, queue)` to subscribe.
	Bind(_scope constructs.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for SqsDestination
type jsiiProxy_SqsDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

func NewSqsDestination(queue awssqs.IQueue) SqsDestination {
	_init_.Initialize()

	j := jsiiProxy_SqsDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_notifications.SqsDestination",
		[]interface{}{queue},
		&j,
	)

	return &j
}

func NewSqsDestination_Override(s SqsDestination, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_notifications.SqsDestination",
		[]interface{}{queue},
		s,
	)
}

func (s *jsiiProxy_SqsDestination) Bind(_scope constructs.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig {
	var returns *awss3.BucketNotificationDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, bucket},
		&returns,
	)

	return returns
}

