package awss3notifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3notifications/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an SNS topic as a bucket notification destination.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//   bucket.AddEventNotification(s3.EventType_OBJECT_CREATED, s3n.NewSnsDestination(topic))
//
type SnsDestination interface {
	awss3.IBucketNotificationDestination
	// Registers this resource to receive notifications for the specified bucket.
	//
	// This method will only be called once for each destination/bucket
	// pair and the result will be cached, so there is no need to implement
	// idempotency in each destination.
	Bind(scope constructs.Construct, bucket awss3.IBucketRef) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for SnsDestination
type jsiiProxy_SnsDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

func NewSnsDestination(topic awssns.ITopic) SnsDestination {
	_init_.Initialize()

	if err := validateNewSnsDestinationParameters(topic); err != nil {
		panic(err)
	}
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

func (s *jsiiProxy_SnsDestination) Bind(scope constructs.Construct, bucket awss3.IBucketRef) *awss3.BucketNotificationDestinationConfig {
	if err := s.validateBindParameters(scope, bucket); err != nil {
		panic(err)
	}
	var returns *awss3.BucketNotificationDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, bucket},
		&returns,
	)

	return returns
}

