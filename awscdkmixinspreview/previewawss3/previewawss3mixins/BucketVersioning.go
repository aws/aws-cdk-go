package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3/previewawss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// S3-specific mixin for enabling versioning.
//
// Example:
//   import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"
//
//
//   bucket := s3.NewCfnBucket(scope, jsii.String("MyBucket")).With(awscdkmixinspreview.NewBucketVersioning()).With(awscdkmixinspreview.NewAutoDeleteObjects())
//
// Experimental.
type BucketVersioning interface {
	core.IMixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for BucketVersioning
type jsiiProxy_BucketVersioning struct {
	internal.Type__coreIMixin
}

// Experimental.
func NewBucketVersioning(enabled *bool) BucketVersioning {
	_init_.Initialize()

	j := jsiiProxy_BucketVersioning{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.BucketVersioning",
		[]interface{}{enabled},
		&j,
	)

	return &j
}

// Experimental.
func NewBucketVersioning_Override(b BucketVersioning, enabled *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.BucketVersioning",
		[]interface{}{enabled},
		b,
	)
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

