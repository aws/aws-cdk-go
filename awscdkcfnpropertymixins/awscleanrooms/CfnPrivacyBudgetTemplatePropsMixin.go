package awscleanrooms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An object that defines the privacy budget template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPrivacyBudgetTemplatePropsMixin := awscdkcfnpropertymixins.Aws_cleanrooms.NewCfnPrivacyBudgetTemplatePropsMixin(&CfnPrivacyBudgetTemplateMixinProps{
//   	AutoRefresh: jsii.String("autoRefresh"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Parameters: &ParametersProperty{
//   		BudgetParameters: []interface{}{
//   			&BudgetParameterProperty{
//   				AutoRefresh: jsii.String("autoRefresh"),
//   				Budget: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Epsilon: jsii.Number(123),
//   		ResourceArn: jsii.String("resourceArn"),
//   		UsersNoisePerQuery: jsii.Number(123),
//   	},
//   	PrivacyBudgetType: jsii.String("privacyBudgetType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html
//
type CfnPrivacyBudgetTemplatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPrivacyBudgetTemplateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPrivacyBudgetTemplatePropsMixin
type jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin) Props() *CfnPrivacyBudgetTemplateMixinProps {
	var returns *CfnPrivacyBudgetTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::PrivacyBudgetTemplate`.
func NewCfnPrivacyBudgetTemplatePropsMixin(props *CfnPrivacyBudgetTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPrivacyBudgetTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPrivacyBudgetTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnPrivacyBudgetTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::PrivacyBudgetTemplate`.
func NewCfnPrivacyBudgetTemplatePropsMixin_Override(c CfnPrivacyBudgetTemplatePropsMixin, props *CfnPrivacyBudgetTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnPrivacyBudgetTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPrivacyBudgetTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPrivacyBudgetTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnPrivacyBudgetTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPrivacyBudgetTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnPrivacyBudgetTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

