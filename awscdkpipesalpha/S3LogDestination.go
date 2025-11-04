package awscdkpipesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// S3 bucket for delivery of pipe logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//
//   s3LogDestination := pipes_alpha.NewS3LogDestination(&S3LogDestinationProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	BucketOwner: jsii.String("bucketOwner"),
//   	OutputFormat: pipes_alpha.S3OutputFormat_PLAIN,
//   	Prefix: jsii.String("prefix"),
//   })
//
// Experimental.
type S3LogDestination interface {
	ILogDestination
	// Bind the log destination to the pipe.
	// Experimental.
	Bind(pipe IPipe) *LogDestinationConfig
	// Grant the pipe role to push to the log destination.
	// Experimental.
	GrantPush(grantee awsiam.IRole)
}

// The jsii proxy struct for S3LogDestination
type jsiiProxy_S3LogDestination struct {
	jsiiProxy_ILogDestination
}

// Experimental.
func NewS3LogDestination(parameters *S3LogDestinationProps) S3LogDestination {
	_init_.Initialize()

	if err := validateNewS3LogDestinationParameters(parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3LogDestination{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.S3LogDestination",
		[]interface{}{parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewS3LogDestination_Override(s S3LogDestination, parameters *S3LogDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.S3LogDestination",
		[]interface{}{parameters},
		s,
	)
}

func (s *jsiiProxy_S3LogDestination) Bind(pipe IPipe) *LogDestinationConfig {
	if err := s.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *LogDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3LogDestination) GrantPush(grantee awsiam.IRole) {
	if err := s.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantPush",
		[]interface{}{grantee},
	)
}

