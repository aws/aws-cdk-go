package awsglue


// Properties for defining a `CfnTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var skewedColumnValueLocationMaps interface{}
//
//   cfnTableProps := &CfnTableProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	TableInput: &TableInputProperty{
//   		Description: jsii.String("description"),
//   		Name: jsii.String("name"),
//   		Owner: jsii.String("owner"),
//   		Parameters: parameters,
//   		PartitionKeys: []interface{}{
//   			&ColumnProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Comment: jsii.String("comment"),
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
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Comment: jsii.String("comment"),
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
//
//   	// the properties below are optional
//   	OpenTableFormatInput: &OpenTableFormatInputProperty{
//   		IcebergInput: &IcebergInputProperty{
//   			MetadataOperation: jsii.String("metadataOperation"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html
//
type CfnTableProps struct {
	// The ID of the Data Catalog in which to create the `Table` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html#cfn-glue-table-catalogid
	//
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the database where the table metadata resides.
	//
	// For Hive compatibility, this must be all lowercase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html#cfn-glue-table-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// A structure used to define a table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html#cfn-glue-table-tableinput
	//
	TableInput interface{} `field:"required" json:"tableInput" yaml:"tableInput"`
	// Specifies an `OpenTableFormatInput` structure when creating an open format table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html#cfn-glue-table-opentableformatinput
	//
	OpenTableFormatInput interface{} `field:"optional" json:"openTableFormatInput" yaml:"openTableFormatInput"`
}

