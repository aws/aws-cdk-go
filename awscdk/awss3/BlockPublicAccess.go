package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &BucketProps{
//   	BlockPublicAccess: s3.BlockPublicAccess_BLOCK_ALL(),
//   })
//
type BlockPublicAccess interface {
	BlockPublicAcls() *bool
	SetBlockPublicAcls(val *bool)
	BlockPublicPolicy() *bool
	SetBlockPublicPolicy(val *bool)
	IgnorePublicAcls() *bool
	SetIgnorePublicAcls(val *bool)
	RestrictPublicBuckets() *bool
	SetRestrictPublicBuckets(val *bool)
}

// The jsii proxy struct for BlockPublicAccess
type jsiiProxy_BlockPublicAccess struct {
	_ byte // padding
}

func (j *jsiiProxy_BlockPublicAccess) BlockPublicAcls() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"blockPublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockPublicAccess) BlockPublicPolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"blockPublicPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockPublicAccess) IgnorePublicAcls() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ignorePublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockPublicAccess) RestrictPublicBuckets() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"restrictPublicBuckets",
		&returns,
	)
	return returns
}


func NewBlockPublicAccess(options *BlockPublicAccessOptions) BlockPublicAccess {
	_init_.Initialize()

	if err := validateNewBlockPublicAccessParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_BlockPublicAccess{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.BlockPublicAccess",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewBlockPublicAccess_Override(b BlockPublicAccess, options *BlockPublicAccessOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.BlockPublicAccess",
		[]interface{}{options},
		b,
	)
}

func (j *jsiiProxy_BlockPublicAccess)SetBlockPublicAcls(val *bool) {
	_jsii_.Set(
		j,
		"blockPublicAcls",
		val,
	)
}

func (j *jsiiProxy_BlockPublicAccess)SetBlockPublicPolicy(val *bool) {
	_jsii_.Set(
		j,
		"blockPublicPolicy",
		val,
	)
}

func (j *jsiiProxy_BlockPublicAccess)SetIgnorePublicAcls(val *bool) {
	_jsii_.Set(
		j,
		"ignorePublicAcls",
		val,
	)
}

func (j *jsiiProxy_BlockPublicAccess)SetRestrictPublicBuckets(val *bool) {
	_jsii_.Set(
		j,
		"restrictPublicBuckets",
		val,
	)
}

func BlockPublicAccess_BLOCK_ACLS() BlockPublicAccess {
	_init_.Initialize()
	var returns BlockPublicAccess
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.BlockPublicAccess",
		"BLOCK_ACLS",
		&returns,
	)
	return returns
}

func BlockPublicAccess_BLOCK_ALL() BlockPublicAccess {
	_init_.Initialize()
	var returns BlockPublicAccess
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.BlockPublicAccess",
		"BLOCK_ALL",
		&returns,
	)
	return returns
}

