package awss3notifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3notifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Lambda function as a bucket notification destination.
//
// Example:
//   var myLambda Function
//
//   bucket := s3.Bucket_FromBucketAttributes(this, jsii.String("ImportedBucket"), &BucketAttributes{
//   	BucketArn: jsii.String("arn:aws:s3:::amzn-s3-demo-bucket"),
//   })
//
//   // now you can just call methods on the bucket
//   bucket.AddEventNotification(s3.EventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &NotificationKeyFilter{
//   	Prefix: jsii.String("home/myusername/*"),
//   })
//
type LambdaDestination interface {
	awss3.IBucketNotificationDestination
	// Registers this resource to receive notifications for the specified bucket.
	//
	// This method will only be called once for each destination/bucket
	// pair and the result will be cached, so there is no need to implement
	// idempotency in each destination.
	Bind(scope constructs.Construct, bucket awss3.IBucketRef) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

func NewLambdaDestination(fn awslambda.IFunction) LambdaDestination {
	_init_.Initialize()

	if err := validateNewLambdaDestinationParameters(fn); err != nil {
		panic(err)
	}
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

func (l *jsiiProxy_LambdaDestination) Bind(scope constructs.Construct, bucket awss3.IBucketRef) *awss3.BucketNotificationDestinationConfig {
	if err := l.validateBindParameters(scope, bucket); err != nil {
		panic(err)
	}
	var returns *awss3.BucketNotificationDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, bucket},
		&returns,
	)

	return returns
}

