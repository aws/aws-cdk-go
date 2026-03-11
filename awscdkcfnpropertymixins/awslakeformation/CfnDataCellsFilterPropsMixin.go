package awslakeformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A structure that represents a data cell filter with column-level, row-level, and/or cell-level security.
//
// Data cell filters belong to a specific table in a Data Catalog . During a stack operation, AWS CloudFormation calls the AWS Lake Formation `CreateDataCellsFilter` API operation to create a `DataCellsFilter` resource, and calls the `DeleteDataCellsFilter` API operation to delete it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var allRowsWildcard interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnDataCellsFilterPropsMixin := awscdkcfnpropertymixins.Aws_lakeformation.NewCfnDataCellsFilterPropsMixin(&CfnDataCellsFilterMixinProps{
//   	ColumnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	ColumnWildcard: &ColumnWildcardProperty{
//   		ExcludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   	DatabaseName: jsii.String("databaseName"),
//   	Name: jsii.String("name"),
//   	RowFilter: &RowFilterProperty{
//   		AllRowsWildcard: allRowsWildcard,
//   		FilterExpression: jsii.String("filterExpression"),
//   	},
//   	TableCatalogId: jsii.String("tableCatalogId"),
//   	TableName: jsii.String("tableName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datacellsfilter.html
//
type CfnDataCellsFilterPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDataCellsFilterMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataCellsFilterPropsMixin
type jsiiProxy_CfnDataCellsFilterPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDataCellsFilterPropsMixin) Props() *CfnDataCellsFilterMixinProps {
	var returns *CfnDataCellsFilterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilterPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LakeFormation::DataCellsFilter`.
func NewCfnDataCellsFilterPropsMixin(props *CfnDataCellsFilterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDataCellsFilterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataCellsFilterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataCellsFilterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataCellsFilterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LakeFormation::DataCellsFilter`.
func NewCfnDataCellsFilterPropsMixin_Override(c CfnDataCellsFilterPropsMixin, props *CfnDataCellsFilterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataCellsFilterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDataCellsFilterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataCellsFilterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataCellsFilterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataCellsFilterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataCellsFilterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataCellsFilterPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataCellsFilterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

