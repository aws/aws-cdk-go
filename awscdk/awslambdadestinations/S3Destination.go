package awslambdadestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdadestinations/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a S3 bucket as a Lambda destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//
//   s3Destination := awscdk.Aws_lambda_destinations.NewS3Destination(bucket)
//
type S3Destination interface {
	awslambda.IDestination
	// Returns a destination configuration.
	Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for S3Destination
type jsiiProxy_S3Destination struct {
	internal.Type__awslambdaIDestination
}

func NewS3Destination(bucket awss3.IBucket) S3Destination {
	_init_.Initialize()

	if err := validateNewS3DestinationParameters(bucket); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Destination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.S3Destination",
		[]interface{}{bucket},
		&j,
	)

	return &j
}

func NewS3Destination_Override(s S3Destination, bucket awss3.IBucket) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.S3Destination",
		[]interface{}{bucket},
		s,
	)
}

func (s *jsiiProxy_S3Destination) Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
	if err := s.validateBindParameters(_scope, fn, _options); err != nil {
		panic(err)
	}
	var returns *awslambda.DestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, fn, _options},
		&returns,
	)

	return returns
}

