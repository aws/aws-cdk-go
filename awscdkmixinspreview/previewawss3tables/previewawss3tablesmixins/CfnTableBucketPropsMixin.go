package previewawss3tablesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3tables/previewawss3tablesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a table bucket.
//
// For more information, see [Creating a table bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-buckets-create.html) in the *Amazon Simple Storage Service User Guide* .
//
// - **Permissions** - - You must have the `s3tables:CreateTableBucket` permission to use this operation.
// - If you use this operation with the optional `encryptionConfiguration` parameter you must have the `s3tables:PutTableBucketEncryption` permission.
// - **Cloud Development Kit** - To use S3 Tables AWS CDK constructs, add the `@aws-cdk/aws-s3tables-alpha` dependency with one of the following options:
//
// - NPM: `npm i.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTableBucketPropsMixin := awscdkmixinspreview.Mixins.NewCfnTableBucketPropsMixin(&CfnTableBucketMixinProps{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SseAlgorithm: jsii.String("sseAlgorithm"),
//   	},
//   	MetricsConfiguration: &MetricsConfigurationProperty{
//   		Status: jsii.String("status"),
//   	},
//   	StorageClassConfiguration: &StorageClassConfigurationProperty{
//   		StorageClass: jsii.String("storageClass"),
//   	},
//   	TableBucketName: jsii.String("tableBucketName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UnreferencedFileRemoval: &UnreferencedFileRemovalProperty{
//   		NoncurrentDays: jsii.Number(123),
//   		Status: jsii.String("status"),
//   		UnreferencedDays: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucket.html
//
type CfnTableBucketPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTableBucketMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTableBucketPropsMixin
type jsiiProxy_CfnTableBucketPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTableBucketPropsMixin) Props() *CfnTableBucketMixinProps {
	var returns *CfnTableBucketMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTableBucketPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Tables::TableBucket`.
func NewCfnTableBucketPropsMixin(props *CfnTableBucketMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTableBucketPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTableBucketPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTableBucketPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTableBucketPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Tables::TableBucket`.
func NewCfnTableBucketPropsMixin_Override(c CfnTableBucketPropsMixin, props *CfnTableBucketMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTableBucketPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTableBucketPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTableBucketPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTableBucketPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTableBucketPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTableBucketPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTableBucketPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTableBucketPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

