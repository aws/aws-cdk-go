package previewawss3objectlambdamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3objectlambda/previewawss3objectlambdamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::S3ObjectLambda::AccessPoint` resource specifies an Object Lambda Access Point used to access a bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var contentTransformation interface{}
//
//   cfnAccessPointPropsMixin := awscdkmixinspreview.Mixins.NewCfnAccessPointPropsMixin(&CfnAccessPointMixinProps{
//   	Name: jsii.String("name"),
//   	ObjectLambdaConfiguration: &ObjectLambdaConfigurationProperty{
//   		AllowedFeatures: []*string{
//   			jsii.String("allowedFeatures"),
//   		},
//   		CloudWatchMetricsEnabled: jsii.Boolean(false),
//   		SupportingAccessPoint: jsii.String("supportingAccessPoint"),
//   		TransformationConfigurations: []interface{}{
//   			&TransformationConfigurationProperty{
//   				Actions: []*string{
//   					jsii.String("actions"),
//   				},
//   				ContentTransformation: contentTransformation,
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspoint.html
//
type CfnAccessPointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAccessPointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccessPointPropsMixin
type jsiiProxy_CfnAccessPointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAccessPointPropsMixin) Props() *CfnAccessPointMixinProps {
	var returns *CfnAccessPointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3ObjectLambda::AccessPoint`.
func NewCfnAccessPointPropsMixin(props *CfnAccessPointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAccessPointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccessPointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccessPointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3objectlambda.mixins.CfnAccessPointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3ObjectLambda::AccessPoint`.
func NewCfnAccessPointPropsMixin_Override(c CfnAccessPointPropsMixin, props *CfnAccessPointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3objectlambda.mixins.CfnAccessPointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAccessPointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccessPointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3objectlambda.mixins.CfnAccessPointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessPointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3objectlambda.mixins.CfnAccessPointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessPointPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAccessPointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

