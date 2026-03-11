package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a launch constraint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLaunchRoleConstraintPropsMixin := awscdkcfnpropertymixins.Aws_servicecatalog.NewCfnLaunchRoleConstraintPropsMixin(&CfnLaunchRoleConstraintMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   	LocalRoleName: jsii.String("localRoleName"),
//   	PortfolioId: jsii.String("portfolioId"),
//   	ProductId: jsii.String("productId"),
//   	RoleArn: jsii.String("roleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html
//
type CfnLaunchRoleConstraintPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLaunchRoleConstraintMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLaunchRoleConstraintPropsMixin
type jsiiProxy_CfnLaunchRoleConstraintPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLaunchRoleConstraintPropsMixin) Props() *CfnLaunchRoleConstraintMixinProps {
	var returns *CfnLaunchRoleConstraintMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraintPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ServiceCatalog::LaunchRoleConstraint`.
func NewCfnLaunchRoleConstraintPropsMixin(props *CfnLaunchRoleConstraintMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLaunchRoleConstraintPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLaunchRoleConstraintPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLaunchRoleConstraintPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnLaunchRoleConstraintPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ServiceCatalog::LaunchRoleConstraint`.
func NewCfnLaunchRoleConstraintPropsMixin_Override(c CfnLaunchRoleConstraintPropsMixin, props *CfnLaunchRoleConstraintMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnLaunchRoleConstraintPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLaunchRoleConstraintPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLaunchRoleConstraintPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnLaunchRoleConstraintPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchRoleConstraintPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnLaunchRoleConstraintPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchRoleConstraintPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraintPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

