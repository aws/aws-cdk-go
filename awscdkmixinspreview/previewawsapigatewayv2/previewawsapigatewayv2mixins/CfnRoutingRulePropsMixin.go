package previewawsapigatewayv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapigatewayv2/previewawsapigatewayv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a routing rule.
//
// When the incoming request to a domain name matches the conditions for a rule, API Gateway invokes a stage of a target API. Supported only for REST APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoutingRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnRoutingRulePropsMixin(&CfnRoutingRuleMixinProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			InvokeApi: &ActionInvokeApiProperty{
//   				ApiId: jsii.String("apiId"),
//   				Stage: jsii.String("stage"),
//   				StripBasePath: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			MatchBasePaths: &MatchBasePathsProperty{
//   				AnyOf: []*string{
//   					jsii.String("anyOf"),
//   				},
//   			},
//   			MatchHeaders: &MatchHeadersProperty{
//   				AnyOf: []interface{}{
//   					&MatchHeaderValueProperty{
//   						Header: jsii.String("header"),
//   						ValueGlob: jsii.String("valueGlob"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DomainNameArn: jsii.String("domainNameArn"),
//   	Priority: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html
//
type CfnRoutingRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRoutingRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRoutingRulePropsMixin
type jsiiProxy_CfnRoutingRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRoutingRulePropsMixin) Props() *CfnRoutingRuleMixinProps {
	var returns *CfnRoutingRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoutingRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGatewayV2::RoutingRule`.
func NewCfnRoutingRulePropsMixin(props *CfnRoutingRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRoutingRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRoutingRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRoutingRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnRoutingRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGatewayV2::RoutingRule`.
func NewCfnRoutingRulePropsMixin_Override(c CfnRoutingRulePropsMixin, props *CfnRoutingRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnRoutingRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRoutingRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRoutingRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnRoutingRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRoutingRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnRoutingRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRoutingRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRoutingRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

