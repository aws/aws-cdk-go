package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cognito::UserPoolResourceServer` resource creates a new OAuth2.0 resource server and defines custom scopes in it.
//
// > If you don't specify a value for a parameter, Amazon Cognito sets it to a default value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnUserPoolResourceServerPropsMixin := awscdkcfnpropertymixins.Aws_cognito.NewCfnUserPoolResourceServerPropsMixin(&CfnUserPoolResourceServerMixinProps{
//   	Identifier: jsii.String("identifier"),
//   	Name: jsii.String("name"),
//   	Scopes: []interface{}{
//   		&ResourceServerScopeTypeProperty{
//   			ScopeDescription: jsii.String("scopeDescription"),
//   			ScopeName: jsii.String("scopeName"),
//   		},
//   	},
//   	UserPoolId: jsii.String("userPoolId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolresourceserver.html
//
type CfnUserPoolResourceServerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUserPoolResourceServerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserPoolResourceServerPropsMixin
type jsiiProxy_CfnUserPoolResourceServerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnUserPoolResourceServerPropsMixin) Props() *CfnUserPoolResourceServerMixinProps {
	var returns *CfnUserPoolResourceServerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolResourceServerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::UserPoolResourceServer`.
func NewCfnUserPoolResourceServerPropsMixin(props *CfnUserPoolResourceServerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUserPoolResourceServerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPoolResourceServerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPoolResourceServerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolResourceServerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::UserPoolResourceServer`.
func NewCfnUserPoolResourceServerPropsMixin_Override(c CfnUserPoolResourceServerPropsMixin, props *CfnUserPoolResourceServerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolResourceServerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUserPoolResourceServerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPoolResourceServerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolResourceServerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolResourceServerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolResourceServerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolResourceServerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUserPoolResourceServerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

