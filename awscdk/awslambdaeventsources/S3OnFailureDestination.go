package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// An S3 dead letter bucket destination configuration for a Lambda event source.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket IBucket
//   var myFunction Function
//
//
//   // Your MSK cluster arn
//   clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//
//   s3OnFailureDestination := awscdk.NewS3OnFailureDestination(bucket)
//
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String(ClusterArn),
//   	Topic: jsii.String(Topic),
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   	OnFailure: s3OnFailureDestination,
//   }))
//
type S3OnFailureDestination interface {
	awslambda.IEventSourceDlq
	// Returns a destination configuration for the DLQ.
	Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig
}

// The jsii proxy struct for S3OnFailureDestination
type jsiiProxy_S3OnFailureDestination struct {
	internal.Type__awslambdaIEventSourceDlq
}

func NewS3OnFailureDestination(bucket awss3.IBucket) S3OnFailureDestination {
	_init_.Initialize()

	if err := validateNewS3OnFailureDestinationParameters(bucket); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3OnFailureDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.S3OnFailureDestination",
		[]interface{}{bucket},
		&j,
	)

	return &j
}

func NewS3OnFailureDestination_Override(s S3OnFailureDestination, bucket awss3.IBucket) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.S3OnFailureDestination",
		[]interface{}{bucket},
		s,
	)
}

func (s *jsiiProxy_S3OnFailureDestination) Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig {
	if err := s.validateBindParameters(_target, targetHandler); err != nil {
		panic(err)
	}
	var returns *awslambda.DlqDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_target, targetHandler},
		&returns,
	)

	return returns
}

