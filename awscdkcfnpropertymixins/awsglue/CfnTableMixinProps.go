package awsglue


// Properties for CfnTablePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var parameters interface{}
//   var properties interface{}
//   var skewedColumnValueLocationMaps interface{}
//
//   cfnTableMixinProps := &CfnTableMixinProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Name: jsii.String("name"),
//   	OpenTableFormatInput: &OpenTableFormatInputProperty{
//   		IcebergInput: &IcebergInputProperty{
//   			IcebergTableInput: &IcebergTableInputProperty{
//   				Location: jsii.String("location"),
//   				PartitionSpec: &IcebergPartitionSpecProperty{
//   					Fields: []interface{}{
//   						&IcebergPartitionFieldProperty{
//   							FieldId: jsii.Number(123),
//   							Name: jsii.String("name"),
//   							SourceId: jsii.Number(123),
//   							Transform: jsii.String("transform"),
//   						},
//   					},
//   					SpecId: jsii.Number(123),
//   				},
//   				Properties: properties,
//   				Schema: &IcebergSchemaProperty{
//   					Fields: []interface{}{
//   						&IcebergStructFieldProperty{
//   							Doc: jsii.String("doc"),
//   							Id: jsii.Number(123),
//   							Name: jsii.String("name"),
//   							Required: jsii.Boolean(false),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					IdentifierFieldIds: []interface{}{
//   						jsii.Number(123),
//   					},
//   					SchemaId: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   				WriteOrder: &IcebergSortOrderProperty{
//   					Fields: []interface{}{
//   						&IcebergSortFieldProperty{
//   							Direction: jsii.String("direction"),
//   							NullOrder: jsii.String("nullOrder"),
//   							SourceId: jsii.Number(123),
//   							Transform: jsii.String("transform"),
//   						},
//   					},
//   					OrderId: jsii.Number(123),
//   				},
//   			},
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
//   		ViewDefinition: &ViewDefinitionProperty{
//   			Definer: jsii.String("definer"),
//   			IsProtected: jsii.Boolean(false),
//   			Representations: []interface{}{
//   				&ViewRepresentationProperty{
//   					Dialect: jsii.String("dialect"),
//   					DialectVersion: jsii.String("dialectVersion"),
//   					ValidationConnection: jsii.String("validationConnection"),
//   					ViewExpandedText: jsii.String("viewExpandedText"),
//   					ViewOriginalText: jsii.String("viewOriginalText"),
//   				},
//   			},
//   			SubObjects: []*string{
//   				jsii.String("subObjects"),
//   			},
//   		},
//   		ViewExpandedText: jsii.String("viewExpandedText"),
//   		ViewOriginalText: jsii.String("viewOriginalText"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html
//
type CfnTableMixinProps struct {
	// The ID of the Data Catalog in which to create the `Table` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html#cfn-glue-table-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database where the table metadata resides.
	//
	// For Hive compatibility, this must be all lowercase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html#cfn-glue-table-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html#cfn-glue-table-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies an `OpenTableFormatInput` structure when creating an open format table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html#cfn-glue-table-opentableformatinput
	//
	OpenTableFormatInput interface{} `field:"optional" json:"openTableFormatInput" yaml:"openTableFormatInput"`
	// A structure used to define a table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html#cfn-glue-table-tableinput
	//
	TableInput interface{} `field:"optional" json:"tableInput" yaml:"tableInput"`
}

