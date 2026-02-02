package previewawsbcmdataexportsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbcmdataexports/previewawsbcmdataexportsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a data export and specifies the data query, the delivery preference, and any optional resource tags.
//
// A `DataQuery` consists of both a `QueryStatement` and `TableConfigurations` .
//
// The `QueryStatement` is an SQL statement. Data Exports only supports a limited subset of the SQL syntax. For more information on the SQL syntax that is supported, see [Data query](https://docs.aws.amazon.com/cur/latest/userguide/de-data-query.html) . To view the available tables and columns, see the [Data Exports table dictionary](https://docs.aws.amazon.com/cur/latest/userguide/de-table-dictionary.html) .
//
// The `TableConfigurations` is a collection of specified `TableProperties` for the table being queried in the `QueryStatement` . TableProperties are additional configurations you can provide to change the data and schema of a table. Each table can have different TableProperties. However, tables are not required to have any TableProperties. Each table property has a default value that it assumes if not specified. For more information on table configurations, see [Data query](https://docs.aws.amazon.com/cur/latest/userguide/de-data-query.html) . To view the table properties available for each table, see the [Data Exports table dictionary](https://docs.aws.amazon.com/cur/latest/userguide/de-table-dictionary.html) or use the `ListTables` API to get a response of all tables and their available properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnExportPropsMixin := awscdkmixinspreview.Mixins.NewCfnExportPropsMixin(&CfnExportMixinProps{
//   	Export: &ExportProperty{
//   		DataQuery: &DataQueryProperty{
//   			QueryStatement: jsii.String("queryStatement"),
//   			TableConfigurations: map[string]interface{}{
//   				"tableConfigurationsKey": map[string]*string{
//   					"tableConfigurationsKey": jsii.String("tableConfigurations"),
//   				},
//   			},
//   		},
//   		Description: jsii.String("description"),
//   		DestinationConfigurations: &DestinationConfigurationsProperty{
//   			S3Destination: &S3DestinationProperty{
//   				S3Bucket: jsii.String("s3Bucket"),
//   				S3OutputConfigurations: &S3OutputConfigurationsProperty{
//   					Compression: jsii.String("compression"),
//   					Format: jsii.String("format"),
//   					OutputType: jsii.String("outputType"),
//   					Overwrite: jsii.String("overwrite"),
//   				},
//   				S3Prefix: jsii.String("s3Prefix"),
//   				S3Region: jsii.String("s3Region"),
//   			},
//   		},
//   		ExportArn: jsii.String("exportArn"),
//   		Name: jsii.String("name"),
//   		RefreshCadence: &RefreshCadenceProperty{
//   			Frequency: jsii.String("frequency"),
//   		},
//   	},
//   	Tags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmdataexports-export.html
//
type CfnExportPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnExportMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnExportPropsMixin
type jsiiProxy_CfnExportPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnExportPropsMixin) Props() *CfnExportMixinProps {
	var returns *CfnExportMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExportPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BCMDataExports::Export`.
func NewCfnExportPropsMixin(props *CfnExportMixinProps, options *mixins.CfnPropertyMixinOptions) CfnExportPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnExportPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnExportPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BCMDataExports::Export`.
func NewCfnExportPropsMixin_Override(c CfnExportPropsMixin, props *CfnExportMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnExportPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExportPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExportPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExportPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnExportPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

