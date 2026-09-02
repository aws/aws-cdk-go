package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::BedrockAgentCore::GatewayRule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnGatewayRulePropsMixin := awscdkcfnpropertymixins.Aws_bedrockagentcore.NewCfnGatewayRulePropsMixin(&CfnGatewayRuleMixinProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			ConfigurationBundle: &ConfigurationBundleActionProperty{
//   				StaticOverride: &StaticOverrideProperty{
//   					BundleArn: jsii.String("bundleArn"),
//   					BundleVersion: jsii.String("bundleVersion"),
//   				},
//   				WeightedOverride: &WeightedOverrideProperty{
//   					TrafficSplit: []interface{}{
//   						&TrafficSplitEntryProperty{
//   							ConfigurationBundle: &ConfigurationBundleReferenceProperty{
//   								BundleArn: jsii.String("bundleArn"),
//   								BundleVersion: jsii.String("bundleVersion"),
//   							},
//   							Description: jsii.String("description"),
//   							Metadata: map[string]*string{
//   								"metadataKey": jsii.String("metadata"),
//   							},
//   							Name: jsii.String("name"),
//   							Weight: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   			RouteToTarget: &RouteToTargetActionProperty{
//   				StaticRoute: &StaticRouteProperty{
//   					TargetName: jsii.String("targetName"),
//   				},
//   				WeightedRoute: &WeightedRouteProperty{
//   					TrafficSplit: []interface{}{
//   						&TargetTrafficSplitEntryProperty{
//   							Description: jsii.String("description"),
//   							Metadata: map[string]*string{
//   								"metadataKey": jsii.String("metadata"),
//   							},
//   							Name: jsii.String("name"),
//   							TargetName: jsii.String("targetName"),
//   							Weight: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			MatchPaths: &MatchPathsProperty{
//   				AnyOf: []*string{
//   					jsii.String("anyOf"),
//   				},
//   			},
//   			MatchPrincipals: &MatchPrincipalsProperty{
//   				AnyOf: []interface{}{
//   					&MatchPrincipalEntryProperty{
//   						IamPrincipal: &IamPrincipalProperty{
//   							Arn: jsii.String("arn"),
//   							Operator: jsii.String("operator"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   	Priority: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html
//
type CfnGatewayRulePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnGatewayRuleMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGatewayRulePropsMixin
type jsiiProxy_CfnGatewayRulePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnGatewayRulePropsMixin) Props() *CfnGatewayRuleMixinProps {
	var returns *CfnGatewayRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayRulePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::GatewayRule`.
func NewCfnGatewayRulePropsMixin(props *CfnGatewayRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnGatewayRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGatewayRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGatewayRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::GatewayRule`.
func NewCfnGatewayRulePropsMixin_Override(c CfnGatewayRulePropsMixin, props *CfnGatewayRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnGatewayRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGatewayRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGatewayRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGatewayRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGatewayRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

