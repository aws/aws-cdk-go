package awslicensemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::LicenseManager::LicenseAssetRuleSet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLicenseAssetRuleSetPropsMixin := awscdkcfnpropertymixins.Aws_licensemanager.NewCfnLicenseAssetRuleSetPropsMixin(&CfnLicenseAssetRuleSetMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Rules: []interface{}{
//   		&LicenseAssetRuleProperty{
//   			RuleStatement: &RuleStatementProperty{
//   				InstanceRuleStatement: &InstanceRuleStatementProperty{
//   					AndRuleStatement: &AndRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   					MatchingRuleStatement: &MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   					OrRuleStatement: &OrRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				LicenseConfigurationRuleStatement: &LicenseConfigurationRuleStatementProperty{
//   					AndRuleStatement: &AndRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   					MatchingRuleStatement: &MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   					OrRuleStatement: &OrRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				LicenseRuleStatement: &LicenseRuleStatementProperty{
//   					AndRuleStatement: &AndRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   					MatchingRuleStatement: &MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   					OrRuleStatement: &OrRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetruleset.html
//
type CfnLicenseAssetRuleSetPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLicenseAssetRuleSetMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLicenseAssetRuleSetPropsMixin
type jsiiProxy_CfnLicenseAssetRuleSetPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLicenseAssetRuleSetPropsMixin) Props() *CfnLicenseAssetRuleSetMixinProps {
	var returns *CfnLicenseAssetRuleSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicenseAssetRuleSetPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LicenseManager::LicenseAssetRuleSet`.
func NewCfnLicenseAssetRuleSetPropsMixin(props *CfnLicenseAssetRuleSetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLicenseAssetRuleSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLicenseAssetRuleSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLicenseAssetRuleSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LicenseManager::LicenseAssetRuleSet`.
func NewCfnLicenseAssetRuleSetPropsMixin_Override(c CfnLicenseAssetRuleSetPropsMixin, props *CfnLicenseAssetRuleSetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLicenseAssetRuleSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLicenseAssetRuleSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLicenseAssetRuleSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLicenseAssetRuleSetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLicenseAssetRuleSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

