package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Use S3 bucket notifications as an event source for AWS Lambda.
//
// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   fn.AddEventSource(eventsources.NewS3EventSource(bucket, &S3EventSourceProps{
//   	Events: []eventType{
//   		s3.*eventType_OBJECT_CREATED,
//   		s3.*eventType_OBJECT_REMOVED,
//   	},
//   	Filters: []notificationKeyFilter{
//   		&notificationKeyFilter{
//   			Prefix: jsii.String("subdir/"),
//   		},
//   	},
//   }))
//
type S3EventSource interface {
	awslambda.IEventSource
	Bucket() awss3.Bucket
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for S3EventSource
type jsiiProxy_S3EventSource struct {
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_S3EventSource) Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}


func NewS3EventSource(bucket awss3.Bucket, props *S3EventSourceProps) S3EventSource {
	_init_.Initialize()

	if err := validateNewS3EventSourceParameters(bucket, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3EventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.S3EventSource",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

func NewS3EventSource_Override(s S3EventSource, bucket awss3.Bucket, props *S3EventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.S3EventSource",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3EventSource) Bind(target awslambda.IFunction) {
	if err := s.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

