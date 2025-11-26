package previewawss3vectorsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3vectors/previewawss3vectorsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::S3Vectors::VectorBucketPolicy` resource defines an Amazon S3 vector bucket policy to control access to an Amazon S3 vector bucket.
//
// Vector bucket policies are written in JSON and allow you to grant or deny permissions across all (or a subset of) objects within a vector bucket.
//
// You must specify either `VectorBucketName` or `VectorBucketArn` to identify the target bucket.
//
// To control how AWS CloudFormation handles the vector bucket policy when the stack is deleted, you can set a deletion policy for your policy. You can choose to *retain* the policy or to *delete* the policy. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// - **Permissions** - The required permissions for CloudFormation to use are based on the operations that are performed on the stack.
//
// - Create
//
// - s3vectors:GetVectorBucketPolicy
// - s3vectors:PutVectorBucketPolicy
// - Read
//
// - s3vectors:GetVectorBucketPolicy
// - Update
//
// - s3vectors:GetVectorBucketPolicy
// - s3vectors:PutVectorBucketPolicy
// - Delete
//
// - s3vectors:GetVectorBucketPolicy
// - s3vectors:DeleteVectorBucketPolicy
// - List
//
// - s3vectors:GetVectorBucketPolicy
// - s3vectors:ListVectorBuckets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnVectorBucketPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnVectorBucketPolicyPropsMixin(&CfnVectorBucketPolicyMixinProps{
//   	Policy: policy,
//   	VectorBucketArn: jsii.String("vectorBucketArn"),
//   	VectorBucketName: jsii.String("vectorBucketName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-vectorbucketpolicy.html
//
type CfnVectorBucketPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVectorBucketPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVectorBucketPolicyPropsMixin
type jsiiProxy_CfnVectorBucketPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVectorBucketPolicyPropsMixin) Props() *CfnVectorBucketPolicyMixinProps {
	var returns *CfnVectorBucketPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVectorBucketPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Vectors::VectorBucketPolicy`.
func NewCfnVectorBucketPolicyPropsMixin(props *CfnVectorBucketPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVectorBucketPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVectorBucketPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVectorBucketPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnVectorBucketPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Vectors::VectorBucketPolicy`.
func NewCfnVectorBucketPolicyPropsMixin_Override(c CfnVectorBucketPolicyPropsMixin, props *CfnVectorBucketPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnVectorBucketPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVectorBucketPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVectorBucketPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnVectorBucketPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVectorBucketPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnVectorBucketPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVectorBucketPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVectorBucketPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

