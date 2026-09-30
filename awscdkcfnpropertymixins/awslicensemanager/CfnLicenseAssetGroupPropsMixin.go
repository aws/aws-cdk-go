package awslicensemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::LicenseManager::LicenseAssetGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLicenseAssetGroupPropsMixin := awscdkcfnpropertymixins.Aws_licensemanager.NewCfnLicenseAssetGroupPropsMixin(&CfnLicenseAssetGroupMixinProps{
//   	AssociatedLicenseAssetRulesetArns: []*string{
//   		jsii.String("associatedLicenseAssetRulesetArns"),
//   	},
//   	Description: jsii.String("description"),
//   	LicenseAssetGroupConfigurations: []interface{}{
//   		&LicenseAssetGroupConfigurationProperty{
//   			UsageDimension: jsii.String("usageDimension"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Properties: []interface{}{
//   		&LicenseAssetGroupPropertyProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetgroup.html
//
type CfnLicenseAssetGroupPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLicenseAssetGroupMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLicenseAssetGroupPropsMixin
type jsiiProxy_CfnLicenseAssetGroupPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLicenseAssetGroupPropsMixin) Props() *CfnLicenseAssetGroupMixinProps {
	var returns *CfnLicenseAssetGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicenseAssetGroupPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LicenseManager::LicenseAssetGroup`.
func NewCfnLicenseAssetGroupPropsMixin(props *CfnLicenseAssetGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLicenseAssetGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLicenseAssetGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLicenseAssetGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LicenseManager::LicenseAssetGroup`.
func NewCfnLicenseAssetGroupPropsMixin_Override(c CfnLicenseAssetGroupPropsMixin, props *CfnLicenseAssetGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLicenseAssetGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLicenseAssetGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLicenseAssetGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLicenseAssetGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLicenseAssetGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

