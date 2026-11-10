package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CloudWatch::AccessGrant.
//
// An access grant authorizes a principal to perform a set of actions on resources in a space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAccessGrantPropsMixin := awscdkcfnpropertymixins.Aws_cloudwatch.NewCfnAccessGrantPropsMixin(&CfnAccessGrantMixinProps{
//   	DomainId: jsii.String("domainId"),
//   	Name: jsii.String("name"),
//   	Permission: jsii.String("permission"),
//   	Principal: &AccessGrantPrincipalProperty{
//   		PrincipalAttributes: []interface{}{
//   			&AccessGrantPrincipalAttributeProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		PrincipalId: jsii.String("principalId"),
//   		PrincipalType: jsii.String("principalType"),
//   	},
//   	ScopedActions: []interface{}{
//   		&ScopedActionsProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			ContextConditions: map[string][]*string{
//   				"contextConditionsKey": []*string{
//   					jsii.String("contextConditions"),
//   				},
//   			},
//   			Resources: []interface{}{
//   				&ResourceScopeProperty{
//   					ResourceArns: []*string{
//   						jsii.String("resourceArns"),
//   					},
//   					ResourceType: jsii.String("resourceType"),
//   					RowScopeGroups: []interface{}{
//   						[]interface{}{
//   							&RowScopeProperty{
//   								Field: jsii.String("field"),
//   								Operator: jsii.String("operator"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   						},
//   					},
//   					SignalTypes: []*string{
//   						jsii.String("signalTypes"),
//   					},
//   					Tags: map[string]*string{
//   						"tagsKey": jsii.String("tags"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SpaceId: jsii.String("spaceId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-accessgrant.html
//
type CfnAccessGrantPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAccessGrantMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccessGrantPropsMixin
type jsiiProxy_CfnAccessGrantPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAccessGrantPropsMixin) Props() *CfnAccessGrantMixinProps {
	var returns *CfnAccessGrantMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessGrantPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::AccessGrant`.
func NewCfnAccessGrantPropsMixin(props *CfnAccessGrantMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAccessGrantPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccessGrantPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccessGrantPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::AccessGrant`.
func NewCfnAccessGrantPropsMixin_Override(c CfnAccessGrantPropsMixin, props *CfnAccessGrantMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAccessGrantPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccessGrantPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessGrantPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessGrantPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAccessGrantPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

