package previewawss3vectorsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3vectors/previewawss3vectorsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::S3Vectors::Index` resource defines a vector index within an Amazon S3 vector bucket.
//
// For more information, see [Creating a vector index in a vector bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors-create-index.html) in the *Amazon Simple Storage Service User Guide* .
//
// You must specify either `VectorBucketName` or `VectorBucketArn` to identify the bucket that contains the index.
//
// To control how AWS CloudFormation handles the vector index when the stack is deleted, you can set a deletion policy for your index. You can choose to *retain* the index or to *delete* the index. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// - **Permissions** - The required permissions for CloudFormation to use are based on the operations that are performed on the stack.
//
// - Create
//
// - s3vectors:CreateIndex
// - s3vectors:GetIndex
// - Read
//
// - s3vectors:GetIndex
// - Delete
//
// - s3vectors:DeleteIndex
// - s3vectors:GetIndex
// - List
//
// - s3vectors:ListIndexes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIndexPropsMixin := awscdkmixinspreview.Mixins.NewCfnIndexPropsMixin(&CfnIndexMixinProps{
//   	DataType: jsii.String("dataType"),
//   	Dimension: jsii.Number(123),
//   	DistanceMetric: jsii.String("distanceMetric"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SseType: jsii.String("sseType"),
//   	},
//   	IndexName: jsii.String("indexName"),
//   	MetadataConfiguration: &MetadataConfigurationProperty{
//   		NonFilterableMetadataKeys: []*string{
//   			jsii.String("nonFilterableMetadataKeys"),
//   		},
//   	},
//   	VectorBucketArn: jsii.String("vectorBucketArn"),
//   	VectorBucketName: jsii.String("vectorBucketName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html
//
type CfnIndexPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIndexMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIndexPropsMixin
type jsiiProxy_CfnIndexPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIndexPropsMixin) Props() *CfnIndexMixinProps {
	var returns *CfnIndexMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndexPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Vectors::Index`.
func NewCfnIndexPropsMixin(props *CfnIndexMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIndexPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIndexPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIndexPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnIndexPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Vectors::Index`.
func NewCfnIndexPropsMixin_Override(c CfnIndexPropsMixin, props *CfnIndexMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnIndexPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIndexPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIndexPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnIndexPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIndexPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3vectors.mixins.CfnIndexPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIndexPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIndexPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

