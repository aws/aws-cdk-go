package awss3notifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3notifications/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Use a Lambda function as a bucket notification destination.
// Experimental.
type LambdaDestination interface {
	awss3.IBucketNotificationDestination
	Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

// Experimental.
func NewLambdaDestination(fn awslambda.IFunction) LambdaDestination {
	_init_.Initialize()

	j := jsiiProxy_LambdaDestination{}

	_jsii_.Create(
		"monocdk.aws_s3_notifications.LambdaDestination",
		[]interface{}{fn},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaDestination_Override(l LambdaDestination, fn awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_s3_notifications.LambdaDestination",
		[]interface{}{fn},
		l,
	)
}

// Registers this resource to receive notifications for the specified bucket.
//
// This method will only be called once for each destination/bucket
// pair and the result will be cached, so there is no need to implement
// idempotency in each destination.
// Experimental.
func (l *jsiiProxy_LambdaDestination) Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig {
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
// Experimental.
type SnsDestination interface {
	awss3.IBucketNotificationDestination
	Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for SnsDestination
type jsiiProxy_SnsDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

// Experimental.
func NewSnsDestination(topic awssns.ITopic) SnsDestination {
	_init_.Initialize()

	j := jsiiProxy_SnsDestination{}

	_jsii_.Create(
		"monocdk.aws_s3_notifications.SnsDestination",
		[]interface{}{topic},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsDestination_Override(s SnsDestination, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_s3_notifications.SnsDestination",
		[]interface{}{topic},
		s,
	)
}

// Registers this resource to receive notifications for the specified bucket.
//
// This method will only be called once for each destination/bucket
// pair and the result will be cached, so there is no need to implement
// idempotency in each destination.
// Experimental.
func (s *jsiiProxy_SnsDestination) Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig {
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
// Experimental.
type SqsDestination interface {
	awss3.IBucketNotificationDestination
	Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for SqsDestination
type jsiiProxy_SqsDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

// Experimental.
func NewSqsDestination(queue awssqs.IQueue) SqsDestination {
	_init_.Initialize()

	j := jsiiProxy_SqsDestination{}

	_jsii_.Create(
		"monocdk.aws_s3_notifications.SqsDestination",
		[]interface{}{queue},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsDestination_Override(s SqsDestination, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_s3_notifications.SqsDestination",
		[]interface{}{queue},
		s,
	)
}

// Allows using SQS queues as destinations for bucket notifications.
//
// Use `bucket.onEvent(event, queue)` to subscribe.
// Experimental.
func (s *jsiiProxy_SqsDestination) Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig {
	var returns *awss3.BucketNotificationDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, bucket},
		&returns,
	)

	return returns
}

