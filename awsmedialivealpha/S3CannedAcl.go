package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// S3 canned ACL for output destinations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   s3CannedAcl := medialive_alpha.S3CannedAcl_AUTHENTICATED_READ()
//
// Experimental.
type S3CannedAcl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for S3CannedAcl
type jsiiProxy_S3CannedAcl struct {
	_ byte // padding
}

func (j *jsiiProxy_S3CannedAcl) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func S3CannedAcl_Of(value *string) S3CannedAcl {
	_init_.Initialize()

	if err := validateS3CannedAcl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns S3CannedAcl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.S3CannedAcl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func S3CannedAcl_AUTHENTICATED_READ() S3CannedAcl {
	_init_.Initialize()
	var returns S3CannedAcl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.S3CannedAcl",
		"AUTHENTICATED_READ",
		&returns,
	)
	return returns
}

func S3CannedAcl_BUCKET_OWNER_FULL_CONTROL() S3CannedAcl {
	_init_.Initialize()
	var returns S3CannedAcl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.S3CannedAcl",
		"BUCKET_OWNER_FULL_CONTROL",
		&returns,
	)
	return returns
}

func S3CannedAcl_BUCKET_OWNER_READ() S3CannedAcl {
	_init_.Initialize()
	var returns S3CannedAcl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.S3CannedAcl",
		"BUCKET_OWNER_READ",
		&returns,
	)
	return returns
}

func S3CannedAcl_PUBLIC_READ() S3CannedAcl {
	_init_.Initialize()
	var returns S3CannedAcl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.S3CannedAcl",
		"PUBLIC_READ",
		&returns,
	)
	return returns
}

