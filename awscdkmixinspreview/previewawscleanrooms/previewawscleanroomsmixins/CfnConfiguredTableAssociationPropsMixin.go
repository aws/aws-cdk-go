package previewawscleanroomsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscleanrooms/previewawscleanroomsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a configured table association.
//
// A configured table association links a configured table with a collaboration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfiguredTableAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnConfiguredTableAssociationPropsMixin(&CfnConfiguredTableAssociationMixinProps{
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html
//
type CfnConfiguredTableAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfiguredTableAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfiguredTableAssociationPropsMixin
type jsiiProxy_CfnConfiguredTableAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnConfiguredTableAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::ConfiguredTableAssociation`.
func NewCfnConfiguredTableAssociationPropsMixin(props *CfnConfiguredTableAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfiguredTableAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfiguredTableAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfiguredTableAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::ConfiguredTableAssociation`.
func NewCfnConfiguredTableAssociationPropsMixin_Override(c CfnConfiguredTableAssociationPropsMixin, props *CfnConfiguredTableAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfiguredTableAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfiguredTableAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin",
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

