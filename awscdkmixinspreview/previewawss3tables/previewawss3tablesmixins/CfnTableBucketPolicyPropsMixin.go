package previewawss3tablesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3tables/previewawss3tablesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new maintenance configuration or replaces an existing table bucket policy for a table bucket.
//
// For more information, see [Adding a table bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-bucket-policy.html#table-bucket-policy-add) in the *Amazon Simple Storage Service User Guide* .
//
// - **Permissions** - You must have the `s3tables:PutTableBucketPolicy` permission to use this operation.
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
//   var resourcePolicy interface{}
//
//   cfnTableBucketPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnTableBucketPolicyPropsMixin(&CfnTableBucketPolicyMixinProps{
//   	ResourcePolicy: resourcePolicy,
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucketpolicy.html
//
type CfnTableBucketPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTableBucketPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTableBucketPolicyPropsMixin
type jsiiProxy_CfnTableBucketPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTableBucketPolicyPropsMixin) Props() *CfnTableBucketPolicyMixinProps {
	var returns *CfnTableBucketPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTableBucketPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Tables::TableBucketPolicy`.
func NewCfnTableBucketPolicyPropsMixin(props *CfnTableBucketPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTableBucketPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTableBucketPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTableBucketPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTableBucketPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Tables::TableBucketPolicy`.
func NewCfnTableBucketPolicyPropsMixin_Override(c CfnTableBucketPolicyPropsMixin, props *CfnTableBucketPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTableBucketPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTableBucketPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTableBucketPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTableBucketPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTableBucketPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTableBucketPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTableBucketPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTableBucketPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

