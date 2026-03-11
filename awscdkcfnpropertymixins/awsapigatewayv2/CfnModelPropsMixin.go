package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGatewayV2::Model` resource updates data model for a WebSocket API.
//
// For more information, see [Model Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) in the *API Gateway Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var schema interface{}
//
//   cfnModelPropsMixin := awscdkcfnpropertymixins.Aws_apigatewayv2.NewCfnModelPropsMixin(&CfnModelMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	ContentType: jsii.String("contentType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Schema: schema,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html
//
type CfnModelPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnModelMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnModelPropsMixin
type jsiiProxy_CfnModelPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnModelPropsMixin) Props() *CfnModelMixinProps {
	var returns *CfnModelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGatewayV2::Model`.
func NewCfnModelPropsMixin(props *CfnModelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnModelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnModelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apigatewayv2.CfnModelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGatewayV2::Model`.
func NewCfnModelPropsMixin_Override(c CfnModelPropsMixin, props *CfnModelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apigatewayv2.CfnModelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnModelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_apigatewayv2.CfnModelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_apigatewayv2.CfnModelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnModelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

