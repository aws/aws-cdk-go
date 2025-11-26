package previewawscloudformationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudformation/previewawscloudformationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudFormation::StackSet` resource contains information about a StackSet.
//
// With StackSets, you can provision stacks across AWS accounts and Regions from a single CloudFormation template. Each stack is based on the same CloudFormation template, but you can customize individual stacks using parameters.
//
// > Run deployments to nested StackSets from the parent stack, not directly through the StackSet API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var managedExecution interface{}
//
//   cfnStackSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnStackSetPropsMixin(&CfnStackSetMixinProps{
//   	AdministrationRoleArn: jsii.String("administrationRoleArn"),
//   	AutoDeployment: &AutoDeploymentProperty{
//   		DependsOn: []*string{
//   			jsii.String("dependsOn"),
//   		},
//   		Enabled: jsii.Boolean(false),
//   		RetainStacksOnAccountRemoval: jsii.Boolean(false),
//   	},
//   	CallAs: jsii.String("callAs"),
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	Description: jsii.String("description"),
//   	ExecutionRoleName: jsii.String("executionRoleName"),
//   	ManagedExecution: managedExecution,
//   	OperationPreferences: &OperationPreferencesProperty{
//   		ConcurrencyMode: jsii.String("concurrencyMode"),
//   		FailureToleranceCount: jsii.Number(123),
//   		FailureTolerancePercentage: jsii.Number(123),
//   		MaxConcurrentCount: jsii.Number(123),
//   		MaxConcurrentPercentage: jsii.Number(123),
//   		RegionConcurrencyType: jsii.String("regionConcurrencyType"),
//   		RegionOrder: []*string{
//   			jsii.String("regionOrder"),
//   		},
//   	},
//   	Parameters: []interface{}{
//   		&ParameterProperty{
//   			ParameterKey: jsii.String("parameterKey"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	PermissionModel: jsii.String("permissionModel"),
//   	StackInstancesGroup: []interface{}{
//   		&StackInstancesProperty{
//   			DeploymentTargets: &DeploymentTargetsProperty{
//   				AccountFilterType: jsii.String("accountFilterType"),
//   				Accounts: []*string{
//   					jsii.String("accounts"),
//   				},
//   				AccountsUrl: jsii.String("accountsUrl"),
//   				OrganizationalUnitIds: []*string{
//   					jsii.String("organizationalUnitIds"),
//   				},
//   			},
//   			ParameterOverrides: []interface{}{
//   				&ParameterProperty{
//   					ParameterKey: jsii.String("parameterKey"),
//   					ParameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   	},
//   	StackSetName: jsii.String("stackSetName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateBody: jsii.String("templateBody"),
//   	TemplateUrl: jsii.String("templateUrl"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html
//
type CfnStackSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStackSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStackSetPropsMixin
type jsiiProxy_CfnStackSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStackSetPropsMixin) Props() *CfnStackSetMixinProps {
	var returns *CfnStackSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFormation::StackSet`.
func NewCfnStackSetPropsMixin(props *CfnStackSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStackSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStackSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStackSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnStackSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFormation::StackSet`.
func NewCfnStackSetPropsMixin_Override(c CfnStackSetPropsMixin, props *CfnStackSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnStackSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStackSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStackSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnStackSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStackSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnStackSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStackSetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

