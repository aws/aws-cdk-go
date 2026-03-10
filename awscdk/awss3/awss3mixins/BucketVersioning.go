package awss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/awss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// S3-specific mixin for enabling versioning.
//
// Example:
//   s3.NewCfnBucket(this, jsii.String("Bucket")).With(#error#.NewBucketVersioning())
//
type BucketVersioning interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for BucketVersioning
type jsiiProxy_BucketVersioning struct {
	internal.Type__awscdkMixin
}

func NewBucketVersioning(enabled *bool) BucketVersioning {
	_init_.Initialize()

	j := jsiiProxy_BucketVersioning{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.mixins.BucketVersioning",
		[]interface{}{enabled},
		&j,
	)

	return &j
}

func NewBucketVersioning_Override(b BucketVersioning, enabled *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.mixins.BucketVersioning",
		[]interface{}{enabled},
		b,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func BucketVersioning_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBucketVersioning_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.mixins.BucketVersioning",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketVersioning) ApplyTo(construct constructs.IConstruct) {
	if err := b.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyTo",
		[]interface{}{construct},
	)
}

func (b *jsiiProxy_BucketVersioning) Supports(construct constructs.IConstruct) *bool {
	if err := b.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		b,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

