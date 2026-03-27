package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Encryption types that can be blocked on an S3 bucket.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MySsecBlockedBucket"), &BucketProps{
//   	BlockedEncryptionTypes: []BlockedEncryptionType{
//   		s3.BlockedEncryptionType_SSE_C(),
//   	},
//   })
//
type BlockedEncryptionType interface {
	// The name for this blocked encryption type used in the API.
	Name() *string
}

// The jsii proxy struct for BlockedEncryptionType
type jsiiProxy_BlockedEncryptionType struct {
	_ byte // padding
}

func (j *jsiiProxy_BlockedEncryptionType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Use this constructor only if S3 releases a new BlockedEncryptionType that is unknown to CDK.
//
// Otherwise, use this class's static constants.
func BlockedEncryptionType_Custom(name *string) BlockedEncryptionType {
	_init_.Initialize()

	if err := validateBlockedEncryptionType_CustomParameters(name); err != nil {
		panic(err)
	}
	var returns BlockedEncryptionType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BlockedEncryptionType",
		"custom",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func BlockedEncryptionType_NONE() BlockedEncryptionType {
	_init_.Initialize()
	var returns BlockedEncryptionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.BlockedEncryptionType",
		"NONE",
		&returns,
	)
	return returns
}

func BlockedEncryptionType_SSE_C() BlockedEncryptionType {
	_init_.Initialize()
	var returns BlockedEncryptionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.BlockedEncryptionType",
		"SSE_C",
		&returns,
	)
	return returns
}

