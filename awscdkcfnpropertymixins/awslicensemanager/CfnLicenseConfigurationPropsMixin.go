package awslicensemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::LicenseManager::LicenseConfiguration.
//
// A license configuration is an abstraction of a customer license agreement that can be consumed and enforced by License Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLicenseConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_licensemanager.NewCfnLicenseConfigurationPropsMixin(&CfnLicenseConfigurationMixinProps{
//   	Description: jsii.String("description"),
//   	DisassociateWhenNotFound: jsii.Boolean(false),
//   	LicenseCount: jsii.Number(123),
//   	LicenseCountHardLimit: jsii.Boolean(false),
//   	LicenseCountingType: jsii.String("licenseCountingType"),
//   	LicenseExpiry: jsii.Number(123),
//   	LicenseRules: []*string{
//   		jsii.String("licenseRules"),
//   	},
//   	Name: jsii.String("name"),
//   	ProductInformationList: []interface{}{
//   		&ProductInformationProperty{
//   			ProductInformationFilterList: []interface{}{
//   				&ProductInformationFilterProperty{
//   					ProductInformationFilterComparator: jsii.String("productInformationFilterComparator"),
//   					ProductInformationFilterName: jsii.String("productInformationFilterName"),
//   					ProductInformationFilterValue: []*string{
//   						jsii.String("productInformationFilterValue"),
//   					},
//   				},
//   			},
//   			ResourceType: jsii.String("resourceType"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseconfiguration.html
//
type CfnLicenseConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLicenseConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLicenseConfigurationPropsMixin
type jsiiProxy_CfnLicenseConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLicenseConfigurationPropsMixin) Props() *CfnLicenseConfigurationMixinProps {
	var returns *CfnLicenseConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicenseConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LicenseManager::LicenseConfiguration`.
func NewCfnLicenseConfigurationPropsMixin(props *CfnLicenseConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLicenseConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLicenseConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLicenseConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LicenseManager::LicenseConfiguration`.
func NewCfnLicenseConfigurationPropsMixin_Override(c CfnLicenseConfigurationPropsMixin, props *CfnLicenseConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLicenseConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLicenseConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLicenseConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLicenseConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLicenseConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

