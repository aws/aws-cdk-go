package previewawsdatazonemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatazone/previewawsdatazonemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The configuration details of an environment blueprint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentBlueprintConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnEnvironmentBlueprintConfigurationPropsMixin(&CfnEnvironmentBlueprintConfigurationMixinProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnabledRegions: []*string{
//   		jsii.String("enabledRegions"),
//   	},
//   	EnvironmentBlueprintIdentifier: jsii.String("environmentBlueprintIdentifier"),
//   	EnvironmentRolePermissionBoundary: jsii.String("environmentRolePermissionBoundary"),
//   	ManageAccessRoleArn: jsii.String("manageAccessRoleArn"),
//   	ProvisioningConfigurations: []interface{}{
//   		&ProvisioningConfigurationProperty{
//   			LakeFormationConfiguration: &LakeFormationConfigurationProperty{
//   				LocationRegistrationExcludeS3Locations: []*string{
//   					jsii.String("locationRegistrationExcludeS3Locations"),
//   				},
//   				LocationRegistrationRole: jsii.String("locationRegistrationRole"),
//   			},
//   		},
//   	},
//   	ProvisioningRoleArn: jsii.String("provisioningRoleArn"),
//   	RegionalParameters: []interface{}{
//   		&RegionalParameterProperty{
//   			Parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			Region: jsii.String("region"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html
//
type CfnEnvironmentBlueprintConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEnvironmentBlueprintConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentBlueprintConfigurationPropsMixin
type jsiiProxy_CfnEnvironmentBlueprintConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEnvironmentBlueprintConfigurationPropsMixin) Props() *CfnEnvironmentBlueprintConfigurationMixinProps {
	var returns *CfnEnvironmentBlueprintConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentBlueprintConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataZone::EnvironmentBlueprintConfiguration`.
func NewCfnEnvironmentBlueprintConfigurationPropsMixin(props *CfnEnvironmentBlueprintConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEnvironmentBlueprintConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentBlueprintConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentBlueprintConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentBlueprintConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataZone::EnvironmentBlueprintConfiguration`.
func NewCfnEnvironmentBlueprintConfigurationPropsMixin_Override(c CfnEnvironmentBlueprintConfigurationPropsMixin, props *CfnEnvironmentBlueprintConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentBlueprintConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEnvironmentBlueprintConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentBlueprintConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentBlueprintConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentBlueprintConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnEnvironmentBlueprintConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentBlueprintConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEnvironmentBlueprintConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

