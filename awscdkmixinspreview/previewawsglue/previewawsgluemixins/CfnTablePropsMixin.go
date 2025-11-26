package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::Table` resource specifies tabular data in the AWS Glue data catalog.
//
// For more information, see [Defining Tables in the AWS Glue Data Catalog](https://docs.aws.amazon.com/glue/latest/dg/tables-described.html) and [Table Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-Table) in the *AWS Glue Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//   var skewedColumnValueLocationMaps interface{}
//
//   cfnTablePropsMixin := awscdkmixinspreview.Mixins.NewCfnTablePropsMixin(&CfnTableMixinProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	OpenTableFormatInput: &OpenTableFormatInputProperty{
//   		IcebergInput: &IcebergInputProperty{
//   			MetadataOperation: jsii.String("metadataOperation"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	TableInput: &TableInputProperty{
//   		Description: jsii.String("description"),
//   		Name: jsii.String("name"),
//   		Owner: jsii.String("owner"),
//   		Parameters: parameters,
//   		PartitionKeys: []interface{}{
//   			&ColumnProperty{
//   				Comment: jsii.String("comment"),
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Retention: jsii.Number(123),
//   		StorageDescriptor: &StorageDescriptorProperty{
//   			BucketColumns: []*string{
//   				jsii.String("bucketColumns"),
//   			},
//   			Columns: []interface{}{
//   				&ColumnProperty{
//   					Comment: jsii.String("comment"),
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Compressed: jsii.Boolean(false),
//   			InputFormat: jsii.String("inputFormat"),
//   			Location: jsii.String("location"),
//   			NumberOfBuckets: jsii.Number(123),
//   			OutputFormat: jsii.String("outputFormat"),
//   			Parameters: parameters,
//   			SchemaReference: &SchemaReferenceProperty{
//   				SchemaId: &SchemaIdProperty{
//   					RegistryName: jsii.String("registryName"),
//   					SchemaArn: jsii.String("schemaArn"),
//   					SchemaName: jsii.String("schemaName"),
//   				},
//   				SchemaVersionId: jsii.String("schemaVersionId"),
//   				SchemaVersionNumber: jsii.Number(123),
//   			},
//   			SerdeInfo: &SerdeInfoProperty{
//   				Name: jsii.String("name"),
//   				Parameters: parameters,
//   				SerializationLibrary: jsii.String("serializationLibrary"),
//   			},
//   			SkewedInfo: &SkewedInfoProperty{
//   				SkewedColumnNames: []*string{
//   					jsii.String("skewedColumnNames"),
//   				},
//   				SkewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   				SkewedColumnValues: []*string{
//   					jsii.String("skewedColumnValues"),
//   				},
//   			},
//   			SortColumns: []interface{}{
//   				&OrderProperty{
//   					Column: jsii.String("column"),
//   					SortOrder: jsii.Number(123),
//   				},
//   			},
//   			StoredAsSubDirectories: jsii.Boolean(false),
//   		},
//   		TableType: jsii.String("tableType"),
//   		TargetTable: &TableIdentifierProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   			Region: jsii.String("region"),
//   		},
//   		ViewExpandedText: jsii.String("viewExpandedText"),
//   		ViewOriginalText: jsii.String("viewOriginalText"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html
//
type CfnTablePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTableMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTablePropsMixin
type jsiiProxy_CfnTablePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTablePropsMixin) Props() *CfnTableMixinProps {
	var returns *CfnTableMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTablePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::Table`.
func NewCfnTablePropsMixin(props *CfnTableMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTablePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTablePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTablePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::Table`.
func NewCfnTablePropsMixin_Override(c CfnTablePropsMixin, props *CfnTableMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTablePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTablePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTablePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTablePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTablePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTablePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

