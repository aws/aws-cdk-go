package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// S3EventSourceV2 Use S3 bucket notifications as an event source for AWS Lambda.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var fn Function
//
//
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("amzn-s3-demo-bucket"))
//
//   fn.AddEventSource(awscdk.NewS3EventSourceV2(bucket, &S3EventSourceProps{
//   	Events: []EventType{
//   		s3.EventType_OBJECT_CREATED,
//   		s3.EventType_OBJECT_REMOVED,
//   	},
//   	Filters: []NotificationKeyFilter{
//   		&NotificationKeyFilter{
//   			Prefix: jsii.String("subdir/"),
//   		},
//   	},
//   }))
//
type S3EventSourceV2 interface {
	awslambda.IEventSource
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for S3EventSourceV2
type jsiiProxy_S3EventSourceV2 struct {
	internal.Type__awslambdaIEventSource
}

func NewS3EventSourceV2(bucket awss3.IBucket, props *S3EventSourceProps) S3EventSourceV2 {
	_init_.Initialize()

	if err := validateNewS3EventSourceV2Parameters(bucket, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3EventSourceV2{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.S3EventSourceV2",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

func NewS3EventSourceV2_Override(s S3EventSourceV2, bucket awss3.IBucket, props *S3EventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.S3EventSourceV2",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3EventSourceV2) Bind(target awslambda.IFunction) {
	if err := s.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

