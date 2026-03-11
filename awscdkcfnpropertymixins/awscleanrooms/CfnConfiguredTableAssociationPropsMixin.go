package awscleanrooms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a configured table association.
//
// A configured table association links a configured table with a collaboration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnConfiguredTableAssociationPropsMixin := awscdkcfnpropertymixins.Aws_cleanrooms.NewCfnConfiguredTableAssociationPropsMixin(&CfnConfiguredTableAssociationMixinProps{
//   	ConfiguredTableAssociationAnalysisRules: []interface{}{
//   		&ConfiguredTableAssociationAnalysisRuleProperty{
//   			Policy: &ConfiguredTableAssociationAnalysisRulePolicyProperty{
//   				V1: &ConfiguredTableAssociationAnalysisRulePolicyV1Property{
//   					Aggregation: &ConfiguredTableAssociationAnalysisRuleAggregationProperty{
//   						AllowedAdditionalAnalyses: []*string{
//   							jsii.String("allowedAdditionalAnalyses"),
//   						},
//   						AllowedResultReceivers: []*string{
//   							jsii.String("allowedResultReceivers"),
//   						},
//   					},
//   					Custom: &ConfiguredTableAssociationAnalysisRuleCustomProperty{
//   						AllowedAdditionalAnalyses: []*string{
//   							jsii.String("allowedAdditionalAnalyses"),
//   						},
//   						AllowedResultReceivers: []*string{
//   							jsii.String("allowedResultReceivers"),
//   						},
//   					},
//   					List: &ConfiguredTableAssociationAnalysisRuleListProperty{
//   						AllowedAdditionalAnalyses: []*string{
//   							jsii.String("allowedAdditionalAnalyses"),
//   						},
//   						AllowedResultReceivers: []*string{
//   							jsii.String("allowedResultReceivers"),
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ConfiguredTableIdentifier: jsii.String("configuredTableIdentifier"),
//   	Description: jsii.String("description"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html
//
type CfnConfiguredTableAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnConfiguredTableAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfiguredTableAssociationPropsMixin
type jsiiProxy_CfnConfiguredTableAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnConfiguredTableAssociationPropsMixin) Props() *CfnConfiguredTableAssociationMixinProps {
	var returns *CfnConfiguredTableAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguredTableAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::ConfiguredTableAssociation`.
func NewCfnConfiguredTableAssociationPropsMixin(props *CfnConfiguredTableAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnConfiguredTableAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfiguredTableAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfiguredTableAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnConfiguredTableAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::ConfiguredTableAssociation`.
func NewCfnConfiguredTableAssociationPropsMixin_Override(c CfnConfiguredTableAssociationPropsMixin, props *CfnConfiguredTableAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnConfiguredTableAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnConfiguredTableAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfiguredTableAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnConfiguredTableAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfiguredTableAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnConfiguredTableAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfiguredTableAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConfiguredTableAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

