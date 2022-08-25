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
//   cfnTableProps := &cfnTableProps{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	tableInput: &tableInputProperty{
//   		description: jsii.String("description"),
//   		name: jsii.String("name"),
//   		owner: jsii.String("owner"),
//   		parameters: parameters,
//   		partitionKeys: []interface{}{
//   			&columnProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				comment: jsii.String("comment"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		retention: jsii.Number(123),
//   		storageDescriptor: &storageDescriptorProperty{
//   			bucketColumns: []*string{
//   				jsii.String("bucketColumns"),
//   			},
//   			columns: []interface{}{
//   				&columnProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					comment: jsii.String("comment"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			compressed: jsii.Boolean(false),
//   			inputFormat: jsii.String("inputFormat"),
//   			location: jsii.String("location"),
//   			numberOfBuckets: jsii.Number(123),
//   			outputFormat: jsii.String("outputFormat"),
//   			parameters: parameters,
//   			schemaReference: &schemaReferenceProperty{
//   				schemaId: &schemaIdProperty{
//   					registryName: jsii.String("registryName"),
//   					schemaArn: jsii.String("schemaArn"),
//   					schemaName: jsii.String("schemaName"),
//   				},
//   				schemaVersionId: jsii.String("schemaVersionId"),
//   				schemaVersionNumber: jsii.Number(123),
//   			},
//   			serdeInfo: &serdeInfoProperty{
//   				name: jsii.String("name"),
//   				parameters: parameters,
//   				serializationLibrary: jsii.String("serializationLibrary"),
//   			},
//   			skewedInfo: &skewedInfoProperty{
//   				skewedColumnNames: []*string{
//   					jsii.String("skewedColumnNames"),
//   				},
//   				skewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   				skewedColumnValues: []*string{
//   					jsii.String("skewedColumnValues"),
//   				},
//   			},
//   			sortColumns: []interface{}{
//   				&orderProperty{
//   					column: jsii.String("column"),
//   					sortOrder: jsii.Number(123),
//   				},
//   			},
//   			storedAsSubDirectories: jsii.Boolean(false),
//   		},
//   		tableType: jsii.String("tableType"),
//   		targetTable: &tableIdentifierProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   		},
//   		viewExpandedText: jsii.String("viewExpandedText"),
//   		viewOriginalText: jsii.String("viewOriginalText"),
//   	},
//   }
//
type CfnTableProps struct {
	// The ID of the Data Catalog in which to create the `Table` .
	//
	// If none is supplied, the AWS account ID is used by default.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the database where the table metadata resides.
	//
	// For Hive compatibility, this must be all lowercase.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// A structure used to define a table.
	TableInput interface{} `field:"required" json:"tableInput" yaml:"tableInput"`
}

