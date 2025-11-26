package previewawsgluemixins


// A structure used to define a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//   var skewedColumnValueLocationMaps interface{}
//
//   tableInputProperty := &TableInputProperty{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Owner: jsii.String("owner"),
//   	Parameters: parameters,
//   	PartitionKeys: []interface{}{
//   		&ColumnProperty{
//   			Comment: jsii.String("comment"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Retention: jsii.Number(123),
//   	StorageDescriptor: &StorageDescriptorProperty{
//   		BucketColumns: []*string{
//   			jsii.String("bucketColumns"),
//   		},
//   		Columns: []interface{}{
//   			&ColumnProperty{
//   				Comment: jsii.String("comment"),
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Compressed: jsii.Boolean(false),
//   		InputFormat: jsii.String("inputFormat"),
//   		Location: jsii.String("location"),
//   		NumberOfBuckets: jsii.Number(123),
//   		OutputFormat: jsii.String("outputFormat"),
//   		Parameters: parameters,
//   		SchemaReference: &SchemaReferenceProperty{
//   			SchemaId: &SchemaIdProperty{
//   				RegistryName: jsii.String("registryName"),
//   				SchemaArn: jsii.String("schemaArn"),
//   				SchemaName: jsii.String("schemaName"),
//   			},
//   			SchemaVersionId: jsii.String("schemaVersionId"),
//   			SchemaVersionNumber: jsii.Number(123),
//   		},
//   		SerdeInfo: &SerdeInfoProperty{
//   			Name: jsii.String("name"),
//   			Parameters: parameters,
//   			SerializationLibrary: jsii.String("serializationLibrary"),
//   		},
//   		SkewedInfo: &SkewedInfoProperty{
//   			SkewedColumnNames: []*string{
//   				jsii.String("skewedColumnNames"),
//   			},
//   			SkewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   			SkewedColumnValues: []*string{
//   				jsii.String("skewedColumnValues"),
//   			},
//   		},
//   		SortColumns: []interface{}{
//   			&OrderProperty{
//   				Column: jsii.String("column"),
//   				SortOrder: jsii.Number(123),
//   			},
//   		},
//   		StoredAsSubDirectories: jsii.Boolean(false),
//   	},
//   	TableType: jsii.String("tableType"),
//   	TargetTable: &TableIdentifierProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Name: jsii.String("name"),
//   		Region: jsii.String("region"),
//   	},
//   	ViewExpandedText: jsii.String("viewExpandedText"),
//   	ViewOriginalText: jsii.String("viewOriginalText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html
//
type CfnTablePropsMixin_TableInputProperty struct {
	// A description of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The table name.
	//
	// For Hive compatibility, this is folded to lowercase when it is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The table owner.
	//
	// Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-owner
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// These key-value pairs define properties associated with the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A list of columns by which the table is partitioned. Only primitive types are supported as partition keys.
	//
	// When you create a table used by Amazon Athena, and you do not specify any `partitionKeys` , you must at least set the value of `partitionKeys` to an empty list. For example:
	//
	// `"PartitionKeys": []`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-partitionkeys
	//
	PartitionKeys interface{} `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	// The retention time for this table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-retention
	//
	Retention *float64 `field:"optional" json:"retention" yaml:"retention"`
	// A storage descriptor containing information about the physical storage of this table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-storagedescriptor
	//
	StorageDescriptor interface{} `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
	// The type of this table.
	//
	// AWS Glue will create tables with the `EXTERNAL_TABLE` type. Other services, such as Athena, may create tables with additional table types.
	//
	// AWS Glue related table types:
	//
	// - **EXTERNAL_TABLE** - Hive compatible attribute - indicates a non-Hive managed table.
	// - **GOVERNED** - Used by AWS Lake Formation . The AWS Glue Data Catalog understands `GOVERNED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-tabletype
	//
	TableType *string `field:"optional" json:"tableType" yaml:"tableType"`
	// A `TableIdentifier` structure that describes a target table for resource linking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-targettable
	//
	TargetTable interface{} `field:"optional" json:"targetTable" yaml:"targetTable"`
	// Included for Apache Hive compatibility.
	//
	// Not used in the normal course of AWS Glue operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-viewexpandedtext
	//
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// Included for Apache Hive compatibility.
	//
	// Not used in the normal course of AWS Glue operations. If the table is a `VIRTUAL_VIEW` , certain Athena configuration encoded in base64.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-vieworiginaltext
	//
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

