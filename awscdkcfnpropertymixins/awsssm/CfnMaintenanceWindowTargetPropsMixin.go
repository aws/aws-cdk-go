package awsssm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SSM::MaintenanceWindowTarget` resource registers a target with a maintenance window for AWS Systems Manager .
//
// For more information, see [RegisterTargetWithMaintenanceWindow](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_RegisterTargetWithMaintenanceWindow.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMaintenanceWindowTargetPropsMixin := awscdkcfnpropertymixins.Aws_ssm.NewCfnMaintenanceWindowTargetPropsMixin(&CfnMaintenanceWindowTargetMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	OwnerInformation: jsii.String("ownerInformation"),
//   	ResourceType: jsii.String("resourceType"),
//   	Targets: []interface{}{
//   		&TargetsProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	WindowId: jsii.String("windowId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html
//
type CfnMaintenanceWindowTargetPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMaintenanceWindowTargetMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMaintenanceWindowTargetPropsMixin
type jsiiProxy_CfnMaintenanceWindowTargetPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMaintenanceWindowTargetPropsMixin) Props() *CfnMaintenanceWindowTargetMixinProps {
	var returns *CfnMaintenanceWindowTargetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTargetPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSM::MaintenanceWindowTarget`.
func NewCfnMaintenanceWindowTargetPropsMixin(props *CfnMaintenanceWindowTargetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMaintenanceWindowTargetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMaintenanceWindowTargetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMaintenanceWindowTargetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTargetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSM::MaintenanceWindowTarget`.
func NewCfnMaintenanceWindowTargetPropsMixin_Override(c CfnMaintenanceWindowTargetPropsMixin, props *CfnMaintenanceWindowTargetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTargetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMaintenanceWindowTargetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMaintenanceWindowTargetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTargetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMaintenanceWindowTargetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTargetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTargetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTargetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

