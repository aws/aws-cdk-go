package awss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/awss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// S3-specific Mixin to automatically delete all objects from a bucket when the bucket is removed from the stack or when the stack is deleted.
//
// Requires the `removalPolicy` to be set to `RemovalPolicy.DESTROY`.
//
// Apply this mixin to a bucket will add `s3:PutBucketPolicy` to the
// bucket policy. This is because during bucket deletion, the custom resource provider
// needs to update the bucket policy by adding a deny policy for `s3:PutObject` to
// prevent race conditions with external bucket writers.
//
// Example:
//   // Apply mixins fluently with .with()
//   // Apply mixins fluently with .with()
//   s3.NewCfnBucket(scope, jsii.String("MyL1Bucket")).With(awscdk.NewBucketBlockPublicAccess()).With(awscdk.NewBucketAutoDeleteObjects())
//
//   // Apply multiple mixins to the same construct
//   // Apply multiple mixins to the same construct
//   s3.NewCfnBucket(scope, jsii.String("MyL1Bucket")).With(awscdk.NewBucketBlockPublicAccess(), awscdk.NewBucketAutoDeleteObjects())
//
//   // Mixins work with all types of constructs:
//   // L1, L2 and even custom constructs
//   // Mixins work with all types of constructs:
//   // L1, L2 and even custom constructs
//   s3.NewBucket(stack, jsii.String("MyL2Bucket")).With(awscdk.NewBucketBlockPublicAccess())
//   NewCustomBucket(stack, jsii.String("MyCustomBucket")).With(awscdk.NewBucketBlockPublicAccess())
//
type BucketAutoDeleteObjects interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for BucketAutoDeleteObjects
type jsiiProxy_BucketAutoDeleteObjects struct {
	internal.Type__awscdkMixin
}

func NewBucketAutoDeleteObjects() BucketAutoDeleteObjects {
	_init_.Initialize()

	j := jsiiProxy_BucketAutoDeleteObjects{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.mixins.BucketAutoDeleteObjects",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewBucketAutoDeleteObjects_Override(b BucketAutoDeleteObjects) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.mixins.BucketAutoDeleteObjects",
		nil, // no parameters
		b,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func BucketAutoDeleteObjects_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBucketAutoDeleteObjects_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.mixins.BucketAutoDeleteObjects",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketAutoDeleteObjects) ApplyTo(construct constructs.IConstruct) {
	if err := b.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyTo",
		[]interface{}{construct},
	)
}

func (b *jsiiProxy_BucketAutoDeleteObjects) Supports(construct constructs.IConstruct) *bool {
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

