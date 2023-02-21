package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Environment file from S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3EnvironmentFile := awscdk.Aws_ecs.NewS3EnvironmentFile(bucket, jsii.String("key"), jsii.String("objectVersion"))
//
type S3EnvironmentFile interface {
	EnvironmentFile
	// Called when the container is initialized to allow this object to bind to the stack.
	Bind(_scope constructs.Construct) *EnvironmentFileConfig
}

// The jsii proxy struct for S3EnvironmentFile
type jsiiProxy_S3EnvironmentFile struct {
	jsiiProxy_EnvironmentFile
}

func NewS3EnvironmentFile(bucket awss3.IBucket, key *string, objectVersion *string) S3EnvironmentFile {
	_init_.Initialize()

	if err := validateNewS3EnvironmentFileParameters(bucket, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3EnvironmentFile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.S3EnvironmentFile",
		[]interface{}{bucket, key, objectVersion},
		&j,
	)

	return &j
}

func NewS3EnvironmentFile_Override(s S3EnvironmentFile, bucket awss3.IBucket, key *string, objectVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.S3EnvironmentFile",
		[]interface{}{bucket, key, objectVersion},
		s,
	)
}

// Loads the environment file from a local disk path.
func S3EnvironmentFile_FromAsset(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	if err := validateS3EnvironmentFile_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetEnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.S3EnvironmentFile",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Loads the environment file from an S3 bucket.
//
// Returns: `S3EnvironmentFile` associated with the specified S3 object.
func S3EnvironmentFile_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3EnvironmentFile {
	_init_.Initialize()

	if err := validateS3EnvironmentFile_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3EnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.S3EnvironmentFile",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3EnvironmentFile) Bind(_scope constructs.Construct) *EnvironmentFileConfig {
	if err := s.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *EnvironmentFileConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

