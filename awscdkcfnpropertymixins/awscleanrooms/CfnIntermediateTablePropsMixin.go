package awscleanrooms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an intermediate table that stores cached query results within a collaboration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnIntermediateTablePropsMixin := awscdkcfnpropertymixins.Aws_cleanrooms.NewCfnIntermediateTablePropsMixin(&CfnIntermediateTableMixinProps{
//   	AnalysisRules: []interface{}{
//   		&IntermediateTableAnalysisRuleProperty{
//   			Policy: &IntermediateTableAnalysisRulePolicyProperty{
//   				V1: &IntermediateTableAnalysisRulePolicyV1Property{
//   					Custom: &IntermediateTableAnalysisRuleCustomProperty{
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AllowedAnalyses: []*string{
//   							jsii.String("allowedAnalyses"),
//   						},
//   						AllowedAnalysisProviders: []*string{
//   							jsii.String("allowedAnalysisProviders"),
//   						},
//   						AllowedResultReceivers: []*string{
//   							jsii.String("allowedResultReceivers"),
//   						},
//   						DifferentialPrivacy: &DifferentialPrivacyProperty{
//   							Columns: []interface{}{
//   								&DifferentialPrivacyColumnProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   						},
//   						DisallowedOutputColumns: []*string{
//   							jsii.String("disallowedOutputColumns"),
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	PopulationAnalysisConfiguration: &PopulationAnalysisConfigurationProperty{
//   		SqlParameters: &PopulationAnalysisSqlParametersProperty{
//   			AnalysisTemplateArn: jsii.String("analysisTemplateArn"),
//   			QueryString: jsii.String("queryString"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-intermediatetable.html
//
type CfnIntermediateTablePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnIntermediateTableMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIntermediateTablePropsMixin
type jsiiProxy_CfnIntermediateTablePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnIntermediateTablePropsMixin) Props() *CfnIntermediateTableMixinProps {
	var returns *CfnIntermediateTableMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntermediateTablePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::IntermediateTable`.
func NewCfnIntermediateTablePropsMixin(props *CfnIntermediateTableMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnIntermediateTablePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIntermediateTablePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIntermediateTablePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnIntermediateTablePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::IntermediateTable`.
func NewCfnIntermediateTablePropsMixin_Override(c CfnIntermediateTablePropsMixin, props *CfnIntermediateTableMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnIntermediateTablePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnIntermediateTablePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntermediateTablePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnIntermediateTablePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntermediateTablePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cleanrooms.CfnIntermediateTablePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIntermediateTablePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIntermediateTablePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

