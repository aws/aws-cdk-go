package previewawsapigatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapigateway/previewawsapigatewaymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGateway::BasePathMappingV2` resource creates a base path that clients who call your API must use in the invocation URL.
//
// Supported only for private custom domain names.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBasePathMappingV2PropsMixin := awscdkmixinspreview.Mixins.NewCfnBasePathMappingV2PropsMixin(&CfnBasePathMappingV2MixinProps{
//   	BasePath: jsii.String("basePath"),
//   	DomainNameArn: jsii.String("domainNameArn"),
//   	RestApiId: jsii.String("restApiId"),
//   	Stage: jsii.String("stage"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html
//
type CfnBasePathMappingV2PropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBasePathMappingV2MixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBasePathMappingV2PropsMixin
type jsiiProxy_CfnBasePathMappingV2PropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBasePathMappingV2PropsMixin) Props() *CfnBasePathMappingV2MixinProps {
	var returns *CfnBasePathMappingV2MixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMappingV2PropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGateway::BasePathMappingV2`.
func NewCfnBasePathMappingV2PropsMixin(props *CfnBasePathMappingV2MixinProps, options *mixins.CfnPropertyMixinOptions) CfnBasePathMappingV2PropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBasePathMappingV2PropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBasePathMappingV2PropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnBasePathMappingV2PropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGateway::BasePathMappingV2`.
func NewCfnBasePathMappingV2PropsMixin_Override(c CfnBasePathMappingV2PropsMixin, props *CfnBasePathMappingV2MixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnBasePathMappingV2PropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBasePathMappingV2PropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBasePathMappingV2PropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnBasePathMappingV2PropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBasePathMappingV2PropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnBasePathMappingV2PropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBasePathMappingV2PropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnBasePathMappingV2PropsMixin) Supports(construct constructs.IConstruct) *bool {
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

