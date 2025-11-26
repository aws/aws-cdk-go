package previewawselasticbeanstalkmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawselasticbeanstalk/previewawselasticbeanstalkmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specify an AWS Elastic Beanstalk application version by using the AWS::ElasticBeanstalk::ApplicationVersion resource in an AWS CloudFormation template.
//
// The AWS::ElasticBeanstalk::ApplicationVersion resource is an AWS Elastic Beanstalk resource type that specifies an application version, an iteration of deployable code, for an Elastic Beanstalk application.
//
// > After you create an application version with a specified Amazon S3 bucket and key location, you can't change that Amazon S3 location. If you change the Amazon S3 location, an attempt to launch an environment from the application version will fail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationVersionPropsMixin(&CfnApplicationVersionMixinProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	Description: jsii.String("description"),
//   	SourceBundle: &SourceBundleProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html
//
type CfnApplicationVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationVersionPropsMixin
type jsiiProxy_CfnApplicationVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationVersionPropsMixin) Props() *CfnApplicationVersionMixinProps {
	var returns *CfnApplicationVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElasticBeanstalk::ApplicationVersion`.
func NewCfnApplicationVersionPropsMixin(props *CfnApplicationVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticbeanstalk.mixins.CfnApplicationVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElasticBeanstalk::ApplicationVersion`.
func NewCfnApplicationVersionPropsMixin_Override(c CfnApplicationVersionPropsMixin, props *CfnApplicationVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticbeanstalk.mixins.CfnApplicationVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticbeanstalk.mixins.CfnApplicationVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_elasticbeanstalk.mixins.CfnApplicationVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

