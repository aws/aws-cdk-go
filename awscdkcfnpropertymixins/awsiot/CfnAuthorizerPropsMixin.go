package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an authorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAuthorizerPropsMixin := awscdkcfnpropertymixins.Aws_iot.NewCfnAuthorizerPropsMixin(&CfnAuthorizerMixinProps{
//   	AuthorizerFunctionArn: jsii.String("authorizerFunctionArn"),
//   	AuthorizerName: jsii.String("authorizerName"),
//   	EnableCachingForHttp: jsii.Boolean(false),
//   	SigningDisabled: jsii.Boolean(false),
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TokenKeyName: jsii.String("tokenKeyName"),
//   	TokenSigningPublicKeys: map[string]*string{
//   		"tokenSigningPublicKeysKey": jsii.String("tokenSigningPublicKeys"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html
//
type CfnAuthorizerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAuthorizerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAuthorizerPropsMixin
type jsiiProxy_CfnAuthorizerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAuthorizerPropsMixin) Props() *CfnAuthorizerMixinProps {
	var returns *CfnAuthorizerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::Authorizer`.
func NewCfnAuthorizerPropsMixin(props *CfnAuthorizerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAuthorizerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAuthorizerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAuthorizerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnAuthorizerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::Authorizer`.
func NewCfnAuthorizerPropsMixin_Override(c CfnAuthorizerPropsMixin, props *CfnAuthorizerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnAuthorizerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAuthorizerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAuthorizerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnAuthorizerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAuthorizerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnAuthorizerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAuthorizerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAuthorizerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

