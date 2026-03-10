package awss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/awss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds statements to a bucket policy.
//
// Example:
//   s3.NewCfnBucketPolicy(this, jsii.String("Policy"), &CfnBucketPolicyProps{
//   	Bucket: s3.NewCfnBucket(this, jsii.String("Bucket")).ref,
//   	PolicyDocument: iam.NewPolicyDocument(),
//   }).With(#error#.NewBucketPolicyStatements([]PolicyStatement{
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
type BucketPolicyStatements interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for BucketPolicyStatements
type jsiiProxy_BucketPolicyStatements struct {
	internal.Type__awscdkMixin
}

func NewBucketPolicyStatements(statements *[]awsiam.PolicyStatement) BucketPolicyStatements {
	_init_.Initialize()

	if err := validateNewBucketPolicyStatementsParameters(statements); err != nil {
		panic(err)
	}
	j := jsiiProxy_BucketPolicyStatements{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.mixins.BucketPolicyStatements",
		[]interface{}{statements},
		&j,
	)

	return &j
}

func NewBucketPolicyStatements_Override(b BucketPolicyStatements, statements *[]awsiam.PolicyStatement) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.mixins.BucketPolicyStatements",
		[]interface{}{statements},
		b,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func BucketPolicyStatements_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBucketPolicyStatements_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.mixins.BucketPolicyStatements",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketPolicyStatements) ApplyTo(construct constructs.IConstruct) {
	if err := b.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyTo",
		[]interface{}{construct},
	)
}

func (b *jsiiProxy_BucketPolicyStatements) Supports(construct constructs.IConstruct) *bool {
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

