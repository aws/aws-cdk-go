package previewawss3vectorsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3vectors/previewawss3vectorsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an Amazon S3 vector bucket in the same AWS Region where you create the AWS CloudFormation stack.
//
// Vector buckets are specialized storage containers designed for storing and managing vector data used in machine learning and AI applications. They provide optimized storage and retrieval capabilities for high-dimensional vector data.
//
// To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to *retain* the bucket or to *delete* the bucket. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// > You can only delete empty vector buckets. Deletion fails for buckets that have contents.
//
// - **Permissions** - The required permissions for CloudFormation to use are based on the operations that are performed on the stack.
//
// - Create
//
// - s3vectors:CreateVectorBucket
// - s3vectors:GetVectorBucket
// - kms:GenerateDataKey (if using KMS encryption)
// - Read
//
// - s3vectors:GetVectorBucket
// - kms:GenerateDataKey (if using KMS encryption)
// - Delete
//
// - s3vectors:DeleteVectorBucket
// - s3vectors:GetVectorBucket
// - kms:GenerateDataKey (if using KMS encryption)
// - List
//
// - s3vectors:ListVectorBuckets
// - kms:GenerateDataKey (if using KMS encryption).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVectorBucketPropsMixin := awscdkmixinspreview.Mixins.NewCfnVectorBucketPropsMixin(&CfnVectorBucketMixinProps{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SseType: jsii.String("sseType"),
//   	},
//   	VectorBucketName: jsii.String("vectorBucketName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-vectorbucket.html
//
type CfnVectorBucketPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVectorBucketMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVectorBucketPropsMixin
type jsiiProxy_CfnVectorBucketPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVectorBucketPropsMixin) Props() *CfnVectorBucketMixinProps {
	var returns *CfnVectorBucketMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVectorBucketPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Vectors::VectorBucket`.
func NewCfnVectorBucketPropsMixin(props *CfnVectorBucketMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVectorBucketPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVectorBucketPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVectorBucketPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnVectorBucketPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Vectors::VectorBucket`.
func NewCfnVectorBucketPropsMixin_Override(c CfnVectorBucketPropsMixin, props *CfnVectorBucketMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnVectorBucketPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVectorBucketPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVectorBucketPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnVectorBucketPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVectorBucketPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnVectorBucketPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVectorBucketPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnVectorBucketPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

