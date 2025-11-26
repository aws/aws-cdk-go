package previewawss3outpostsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3outposts/previewawss3outpostsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// This resource applies a bucket policy to an Amazon S3 on Outposts bucket.
//
// If you are using an identity other than the root user of the AWS account that owns the S3 on Outposts bucket, the calling identity must have the `s3-outposts:PutBucketPolicy` permissions on the specified Outposts bucket and belong to the bucket owner's account in order to use this resource.
//
// If you don't have `s3-outposts:PutBucketPolicy` permissions, S3 on Outposts returns a `403 Access Denied` error.
//
// > The root user of the AWS account that owns an Outposts bucket can *always* use this resource, even if the policy explicitly denies the root user the ability to perform actions on this resource.
//
// For more information, see the AWS::IAM::Policy [PolicyDocument](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument) resource description in this guide and [Access Policy Language Overview](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-policy-language-overview.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   cfnBucketPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnBucketPolicyPropsMixin(&CfnBucketPolicyMixinProps{
//   	Bucket: jsii.String("bucket"),
//   	PolicyDocument: policyDocument,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-bucketpolicy.html
//
type CfnBucketPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBucketPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBucketPolicyPropsMixin
type jsiiProxy_CfnBucketPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBucketPolicyPropsMixin) Props() *CfnBucketPolicyMixinProps {
	var returns *CfnBucketPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Outposts::BucketPolicy`.
func NewCfnBucketPolicyPropsMixin(props *CfnBucketPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBucketPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBucketPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBucketPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3outposts.mixins.CfnBucketPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Outposts::BucketPolicy`.
func NewCfnBucketPolicyPropsMixin_Override(c CfnBucketPolicyPropsMixin, props *CfnBucketPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3outposts.mixins.CfnBucketPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBucketPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBucketPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3outposts.mixins.CfnBucketPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBucketPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3outposts.mixins.CfnBucketPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBucketPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnBucketPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

