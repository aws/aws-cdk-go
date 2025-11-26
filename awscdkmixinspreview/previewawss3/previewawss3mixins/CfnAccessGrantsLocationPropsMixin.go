package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3/previewawss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::S3::AccessGrantsLocation` resource creates the S3 data location that you would like to register in your S3 Access Grants instance.
//
// Your S3 data must be in the same Region as your S3 Access Grants instance. The location can be one of the following:
//
// - The default S3 location `s3://`
// - A bucket - `S3://<bucket-name>`
// - A bucket and prefix - `S3://<bucket-name>/<prefix>`
//
// When you register a location, you must include the IAM role that has permission to manage the S3 location that you are registering. Give S3 Access Grants permission to assume this role [using a policy](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-location.html) . S3 Access Grants assumes this role to manage access to the location and to vend temporary credentials to grantees or client applications.
//
// - **Permissions** - You must have the `s3:CreateAccessGrantsLocation` permission to use this resource.
// - **Additional Permissions** - You must also have the following permission for the specified IAM role: `iam:PassRole`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccessGrantsLocationPropsMixin := awscdkmixinspreview.Mixins.NewCfnAccessGrantsLocationPropsMixin(&CfnAccessGrantsLocationMixinProps{
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	LocationScope: jsii.String("locationScope"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrantslocation.html
//
type CfnAccessGrantsLocationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAccessGrantsLocationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccessGrantsLocationPropsMixin
type jsiiProxy_CfnAccessGrantsLocationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAccessGrantsLocationPropsMixin) Props() *CfnAccessGrantsLocationMixinProps {
	var returns *CfnAccessGrantsLocationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessGrantsLocationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3::AccessGrantsLocation`.
func NewCfnAccessGrantsLocationPropsMixin(props *CfnAccessGrantsLocationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAccessGrantsLocationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccessGrantsLocationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccessGrantsLocationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsLocationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3::AccessGrantsLocation`.
func NewCfnAccessGrantsLocationPropsMixin_Override(c CfnAccessGrantsLocationPropsMixin, props *CfnAccessGrantsLocationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsLocationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAccessGrantsLocationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccessGrantsLocationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsLocationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessGrantsLocationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsLocationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessGrantsLocationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAccessGrantsLocationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

