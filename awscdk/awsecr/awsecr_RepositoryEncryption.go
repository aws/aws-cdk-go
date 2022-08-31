package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Indicates whether server-side encryption is enabled for the object, and whether that encryption is from the AWS Key Management Service (AWS KMS) or from Amazon S3 managed encryption (SSE-S3).
//
// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &repositoryProps{
//   	encryption: ecr.repositoryEncryption_KMS(),
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Experimental.
type RepositoryEncryption interface {
	// the string value of the encryption.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for RepositoryEncryption
type jsiiProxy_RepositoryEncryption struct {
	_ byte // padding
}

func (j *jsiiProxy_RepositoryEncryption) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewRepositoryEncryption(value *string) RepositoryEncryption {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEncryption{}

	_jsii_.Create(
		"monocdk.aws_ecr.RepositoryEncryption",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEncryption_Override(r RepositoryEncryption, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr.RepositoryEncryption",
		[]interface{}{value},
		r,
	)
}

func RepositoryEncryption_AES_256() RepositoryEncryption {
	_init_.Initialize()
	var returns RepositoryEncryption
	_jsii_.StaticGet(
		"monocdk.aws_ecr.RepositoryEncryption",
		"AES_256",
		&returns,
	)
	return returns
}

func RepositoryEncryption_KMS() RepositoryEncryption {
	_init_.Initialize()
	var returns RepositoryEncryption
	_jsii_.StaticGet(
		"monocdk.aws_ecr.RepositoryEncryption",
		"KMS",
		&returns,
	)
	return returns
}

