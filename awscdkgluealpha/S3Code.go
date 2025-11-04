package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Glue job Code from an S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//
//   s3Code := glue_alpha.NewS3Code(bucket, jsii.String("key"))
//
// Experimental.
type S3Code interface {
	Code
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig
}

// The jsii proxy struct for S3Code
type jsiiProxy_S3Code struct {
	jsiiProxy_Code
}

// Experimental.
func NewS3Code(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	if err := validateNewS3CodeParameters(bucket, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Code{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.S3Code",
		[]interface{}{bucket, key},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Code_Override(s S3Code, bucket awss3.IBucket, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.S3Code",
		[]interface{}{bucket, key},
		s,
	)
}

// Job code from a local disk path.
// Experimental.
func S3Code_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateS3Code_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Job code as an S3 object.
// Experimental.
func S3Code_FromBucket(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	if err := validateS3Code_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3Code",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Code) Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig {
	if err := s.validateBindParameters(scope, grantable); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, grantable},
		&returns,
	)

	return returns
}

