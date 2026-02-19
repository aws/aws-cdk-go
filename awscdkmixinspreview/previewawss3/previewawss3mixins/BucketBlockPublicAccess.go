package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3/previewawss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// S3-specific mixin for blocking public-access.
//
// Example:
//   bucket := s3.NewCfnBucket(scope, jsii.String("Bucket"))
//   awscdkmixinspreview.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewBucketBlockPublicAccess())
//
// Experimental.
type BucketBlockPublicAccess interface {
	core.IMixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for BucketBlockPublicAccess
type jsiiProxy_BucketBlockPublicAccess struct {
	internal.Type__coreIMixin
}

// Experimental.
func NewBucketBlockPublicAccess(publicAccessConfig awss3.BlockPublicAccess) BucketBlockPublicAccess {
	_init_.Initialize()

	j := jsiiProxy_BucketBlockPublicAccess{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.BucketBlockPublicAccess",
		[]interface{}{publicAccessConfig},
		&j,
	)

	return &j
}

// Experimental.
func NewBucketBlockPublicAccess_Override(b BucketBlockPublicAccess, publicAccessConfig awss3.BlockPublicAccess) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.BucketBlockPublicAccess",
		[]interface{}{publicAccessConfig},
		b,
	)
}

func (b *jsiiProxy_BucketBlockPublicAccess) ApplyTo(construct constructs.IConstruct) {
	if err := b.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyTo",
		[]interface{}{construct},
	)
}

func (b *jsiiProxy_BucketBlockPublicAccess) Supports(construct constructs.IConstruct) *bool {
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

