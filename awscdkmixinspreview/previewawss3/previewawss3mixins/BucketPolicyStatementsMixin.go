package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3/previewawss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds statements to a bucket policy.
//
// Example:
//   var bucket IBucketRef
//
//
//   bucketPolicy := s3.NewCfnBucketPolicy(scope, jsii.String("BucketPolicy"), &CfnBucketPolicyProps{
//   	Bucket: bucket,
//   	PolicyDocument: iam.NewPolicyDocument(),
//   })
//   awscdkmixinspreview.Mixins_Of(bucketPolicy).Apply(awscdkmixinspreview.NewBucketPolicyStatementsMixin([]PolicyStatement{
//   	iam.NewPolicyStatement(&PolicyStatementProps{
//   		Actions: []*string{
//   			jsii.String("s3:GetObject"),
//   		},
//   		Resources: []*string{
//   			jsii.String("*"),
//   		},
//   		Principals: []IPrincipal{
//   			iam.NewAnyPrincipal(),
//   		},
//   	}),
//   }))
//
// Experimental.
type BucketPolicyStatementsMixin interface {
	core.Mixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for BucketPolicyStatementsMixin
type jsiiProxy_BucketPolicyStatementsMixin struct {
	internal.Type__coreMixin
}

// Experimental.
func NewBucketPolicyStatementsMixin(statements *[]awsiam.PolicyStatement) BucketPolicyStatementsMixin {
	_init_.Initialize()

	if err := validateNewBucketPolicyStatementsMixinParameters(statements); err != nil {
		panic(err)
	}
	j := jsiiProxy_BucketPolicyStatementsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.BucketPolicyStatementsMixin",
		[]interface{}{statements},
		&j,
	)

	return &j
}

// Experimental.
func NewBucketPolicyStatementsMixin_Override(b BucketPolicyStatementsMixin, statements *[]awsiam.PolicyStatement) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.BucketPolicyStatementsMixin",
		[]interface{}{statements},
		b,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func BucketPolicyStatementsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBucketPolicyStatementsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.mixins.BucketPolicyStatementsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketPolicyStatementsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := b.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyTo",
		[]interface{}{construct},
	)
}

func (b *jsiiProxy_BucketPolicyStatementsMixin) Supports(construct constructs.IConstruct) *bool {
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

