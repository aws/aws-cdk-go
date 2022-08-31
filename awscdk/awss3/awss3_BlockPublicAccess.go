package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &bucketProps{
//   	blockPublicAccess: s3.blockPublicAccess_BLOCK_ALL(),
//   })
//
// Experimental.
type BlockPublicAccess interface {
	// Experimental.
	BlockPublicAcls() *bool
	// Experimental.
	SetBlockPublicAcls(val *bool)
	// Experimental.
	BlockPublicPolicy() *bool
	// Experimental.
	SetBlockPublicPolicy(val *bool)
	// Experimental.
	IgnorePublicAcls() *bool
	// Experimental.
	SetIgnorePublicAcls(val *bool)
	// Experimental.
	RestrictPublicBuckets() *bool
	// Experimental.
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


// Experimental.
func NewBlockPublicAccess(options *BlockPublicAccessOptions) BlockPublicAccess {
	_init_.Initialize()

	j := jsiiProxy_BlockPublicAccess{}

	_jsii_.Create(
		"monocdk.aws_s3.BlockPublicAccess",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewBlockPublicAccess_Override(b BlockPublicAccess, options *BlockPublicAccessOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_s3.BlockPublicAccess",
		[]interface{}{options},
		b,
	)
}

func (j *jsiiProxy_BlockPublicAccess) SetBlockPublicAcls(val *bool) {
	_jsii_.Set(
		j,
		"blockPublicAcls",
		val,
	)
}

func (j *jsiiProxy_BlockPublicAccess) SetBlockPublicPolicy(val *bool) {
	_jsii_.Set(
		j,
		"blockPublicPolicy",
		val,
	)
}

func (j *jsiiProxy_BlockPublicAccess) SetIgnorePublicAcls(val *bool) {
	_jsii_.Set(
		j,
		"ignorePublicAcls",
		val,
	)
}

func (j *jsiiProxy_BlockPublicAccess) SetRestrictPublicBuckets(val *bool) {
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
		"monocdk.aws_s3.BlockPublicAccess",
		"BLOCK_ACLS",
		&returns,
	)
	return returns
}

func BlockPublicAccess_BLOCK_ALL() BlockPublicAccess {
	_init_.Initialize()
	var returns BlockPublicAccess
	_jsii_.StaticGet(
		"monocdk.aws_s3.BlockPublicAccess",
		"BLOCK_ALL",
		&returns,
	)
	return returns
}

