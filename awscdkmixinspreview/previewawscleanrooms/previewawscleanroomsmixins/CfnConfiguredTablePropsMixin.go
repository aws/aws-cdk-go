package previewawscleanroomsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscleanrooms/previewawscleanroomsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new configured table resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfiguredTablePropsMixin := awscdkmixinspreview.Mixins.NewCfnConfiguredTablePropsMixin(&CfnConfiguredTableMixinProps{
//   	AllowedColumns: []*string{
//   		jsii.String("allowedColumns"),
//   	},
//   	AnalysisMethod: jsii.String("analysisMethod"),
//   	AnalysisRules: []interface{}{
//   		&AnalysisRuleProperty{
//   			Policy: &ConfiguredTableAnalysisRulePolicyProperty{
//   				V1: &ConfiguredTableAnalysisRulePolicyV1Property{
//   					Aggregation: &AnalysisRuleAggregationProperty{
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AggregateColumns: []interface{}{
//   							&AggregateColumnProperty{
//   								ColumnNames: []*string{
//   									jsii.String("columnNames"),
//   								},
//   								Function: jsii.String("function"),
//   							},
//   						},
//   						AllowedJoinOperators: []*string{
//   							jsii.String("allowedJoinOperators"),
//   						},
//   						DimensionColumns: []*string{
//   							jsii.String("dimensionColumns"),
//   						},
//   						JoinColumns: []*string{
//   							jsii.String("joinColumns"),
//   						},
//   						JoinRequired: jsii.String("joinRequired"),
//   						OutputConstraints: []interface{}{
//   							&AggregationConstraintProperty{
//   								ColumnName: jsii.String("columnName"),
//   								Minimum: jsii.Number(123),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						ScalarFunctions: []*string{
//   							jsii.String("scalarFunctions"),
//   						},
//   					},
//   					Custom: &AnalysisRuleCustomProperty{
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AllowedAnalyses: []*string{
//   							jsii.String("allowedAnalyses"),
//   						},
//   						AllowedAnalysisProviders: []*string{
//   							jsii.String("allowedAnalysisProviders"),
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
//   					List: &AnalysisRuleListProperty{
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AllowedJoinOperators: []*string{
//   							jsii.String("allowedJoinOperators"),
//   						},
//   						JoinColumns: []*string{
//   							jsii.String("joinColumns"),
//   						},
//   						ListColumns: []*string{
//   							jsii.String("listColumns"),
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	SelectedAnalysisMethods: []*string{
//   		jsii.String("selectedAnalysisMethods"),
//   	},
//   	TableReference: &TableReferenceProperty{
//   		Athena: &AthenaTableReferenceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			OutputLocation: jsii.String("outputLocation"),
//   			Region: jsii.String("region"),
//   			TableName: jsii.String("tableName"),
//   			WorkGroup: jsii.String("workGroup"),
//   		},
//   		Glue: &GlueTableReferenceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Region: jsii.String("region"),
//   			TableName: jsii.String("tableName"),
//   		},
//   		Snowflake: &SnowflakeTableReferenceProperty{
//   			AccountIdentifier: jsii.String("accountIdentifier"),
//   			DatabaseName: jsii.String("databaseName"),
//   			SchemaName: jsii.String("schemaName"),
//   			SecretArn: jsii.String("secretArn"),
//   			TableName: jsii.String("tableName"),
//   			TableSchema: &SnowflakeTableSchemaProperty{
//   				V1: []interface{}{
//   					&SnowflakeTableSchemaV1Property{
//   						ColumnName: jsii.String("columnName"),
//   						ColumnType: jsii.String("columnType"),
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html
//
type CfnConfiguredTablePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfiguredTableMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfiguredTablePropsMixin
type jsiiProxy_CfnConfiguredTablePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfiguredTablePropsMixin) Props() *CfnConfiguredTableMixinProps {
	var returns *CfnConfiguredTableMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguredTablePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::ConfiguredTable`.
func NewCfnConfiguredTablePropsMixin(props *CfnConfiguredTableMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfiguredTablePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfiguredTablePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfiguredTablePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::ConfiguredTable`.
func NewCfnConfiguredTablePropsMixin_Override(c CfnConfiguredTablePropsMixin, props *CfnConfiguredTableMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfiguredTablePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfiguredTablePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfiguredTablePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfiguredTablePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConfiguredTablePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

