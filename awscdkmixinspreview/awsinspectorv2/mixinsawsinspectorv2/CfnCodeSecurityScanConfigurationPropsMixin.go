package mixinsawsinspectorv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsinspectorv2/mixinsawsinspectorv2/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a scan configuration for code security scanning.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeSecurityScanConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnCodeSecurityScanConfigurationPropsMixin(&CfnCodeSecurityScanConfigurationMixinProps{
//   	Configuration: &CodeSecurityScanConfigurationProperty{
//   		ContinuousIntegrationScanConfiguration: &ContinuousIntegrationScanConfigurationProperty{
//   			SupportedEvents: []*string{
//   				jsii.String("supportedEvents"),
//   			},
//   		},
//   		PeriodicScanConfiguration: &PeriodicScanConfigurationProperty{
//   			Frequency: jsii.String("frequency"),
//   			FrequencyExpression: jsii.String("frequencyExpression"),
//   		},
//   		RuleSetCategories: []*string{
//   			jsii.String("ruleSetCategories"),
//   		},
//   	},
//   	Level: jsii.String("level"),
//   	Name: jsii.String("name"),
//   	ScopeSettings: &ScopeSettingsProperty{
//   		ProjectSelectionScope: jsii.String("projectSelectionScope"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityscanconfiguration.html
//
type CfnCodeSecurityScanConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCodeSecurityScanConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCodeSecurityScanConfigurationPropsMixin
type jsiiProxy_CfnCodeSecurityScanConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCodeSecurityScanConfigurationPropsMixin) Props() *CfnCodeSecurityScanConfigurationMixinProps {
	var returns *CfnCodeSecurityScanConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSecurityScanConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::InspectorV2::CodeSecurityScanConfiguration`.
func NewCfnCodeSecurityScanConfigurationPropsMixin(props *CfnCodeSecurityScanConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCodeSecurityScanConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCodeSecurityScanConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCodeSecurityScanConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::InspectorV2::CodeSecurityScanConfiguration`.
func NewCfnCodeSecurityScanConfigurationPropsMixin_Override(c CfnCodeSecurityScanConfigurationPropsMixin, props *CfnCodeSecurityScanConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCodeSecurityScanConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCodeSecurityScanConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCodeSecurityScanConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCodeSecurityScanConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCodeSecurityScanConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

