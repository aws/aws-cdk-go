package awsglue


// Properties for defining a `CfnPartition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var skewedColumnValueLocationMaps interface{}
//
//   cfnPartitionProps := &cfnPartitionProps{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	partitionInput: &partitionInputProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//
//   		// the properties below are optional
//   		parameters: parameters,
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
//
//   					// the properties below are optional
//   					sortOrder: jsii.Number(123),
//   				},
//   			},
//   			storedAsSubDirectories: jsii.Boolean(false),
//   		},
//   	},
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnPartitionProps struct {
	// The AWS account ID of the catalog in which the partion is to be created.
	//
	// > To specify the account ID, you can use the `Ref` intrinsic function with the `AWS::AccountId` pseudo parameter. For example: `!Ref AWS::AccountId`
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the catalog database in which to create the partition.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The structure used to create and update a partition.
	PartitionInput interface{} `field:"required" json:"partitionInput" yaml:"partitionInput"`
	// The name of the metadata table in which the partition is to be created.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

