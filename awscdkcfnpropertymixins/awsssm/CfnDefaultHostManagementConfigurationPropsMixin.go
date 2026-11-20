package awsssm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::SSM::DefaultHostManagementConfiguration.
//
// Enables or disables Default Host Management Configuration (DHMC) for an AWS account and Region by managing the '/ssm/managed-instance/default-ec2-instance-management-role' service setting. This is a singleton resource - there is at most one DHMC configuration per account/Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDefaultHostManagementConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_ssm.NewCfnDefaultHostManagementConfigurationPropsMixin(&CfnDefaultHostManagementConfigurationMixinProps{
//   	CustomerManagedRoleArn: jsii.String("customerManagedRoleArn"),
//   	PermissionModel: jsii.String("permissionModel"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-defaulthostmanagementconfiguration.html
//
type CfnDefaultHostManagementConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDefaultHostManagementConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDefaultHostManagementConfigurationPropsMixin
type jsiiProxy_CfnDefaultHostManagementConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDefaultHostManagementConfigurationPropsMixin) Props() *CfnDefaultHostManagementConfigurationMixinProps {
	var returns *CfnDefaultHostManagementConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDefaultHostManagementConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSM::DefaultHostManagementConfiguration`.
func NewCfnDefaultHostManagementConfigurationPropsMixin(props *CfnDefaultHostManagementConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDefaultHostManagementConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDefaultHostManagementConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDefaultHostManagementConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnDefaultHostManagementConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSM::DefaultHostManagementConfiguration`.
func NewCfnDefaultHostManagementConfigurationPropsMixin_Override(c CfnDefaultHostManagementConfigurationPropsMixin, props *CfnDefaultHostManagementConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnDefaultHostManagementConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDefaultHostManagementConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDefaultHostManagementConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnDefaultHostManagementConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDefaultHostManagementConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnDefaultHostManagementConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDefaultHostManagementConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDefaultHostManagementConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

