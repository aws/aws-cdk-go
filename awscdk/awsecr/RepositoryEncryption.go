package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Indicates whether server-side encryption is enabled for the object, and whether that encryption is from the AWS Key Management Service (AWS KMS) or from Amazon S3 managed encryption (SSE-S3).
//
// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
//   	Encryption: ecr.RepositoryEncryption_KMS(),
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
type RepositoryEncryption interface {
	// the string value of the encryption.
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


func NewRepositoryEncryption(value *string) RepositoryEncryption {
	_init_.Initialize()

	if err := validateNewRepositoryEncryptionParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositoryEncryption{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.RepositoryEncryption",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewRepositoryEncryption_Override(r RepositoryEncryption, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.RepositoryEncryption",
		[]interface{}{value},
		r,
	)
}

func RepositoryEncryption_AES_256() RepositoryEncryption {
	_init_.Initialize()
	var returns RepositoryEncryption
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr.RepositoryEncryption",
		"AES_256",
		&returns,
	)
	return returns
}

func RepositoryEncryption_KMS() RepositoryEncryption {
	_init_.Initialize()
	var returns RepositoryEncryption
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr.RepositoryEncryption",
		"KMS",
		&returns,
	)
	return returns
}

