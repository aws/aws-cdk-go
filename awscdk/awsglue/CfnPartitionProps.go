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
//   cfnPartitionProps := &CfnPartitionProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	PartitionInput: &PartitionInputProperty{
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//
//   		// the properties below are optional
//   		Parameters: parameters,
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
//
//   					// the properties below are optional
//   					SortOrder: jsii.Number(123),
//   				},
//   			},
//   			StoredAsSubDirectories: jsii.Boolean(false),
//   		},
//   	},
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-partition.html
//
type CfnPartitionProps struct {
	// The AWS account ID of the catalog in which the partion is to be created.
	//
	// > To specify the account ID, you can use the `Ref` intrinsic function with the `AWS::AccountId` pseudo parameter. For example: `!Ref AWS::AccountId`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-partition.html#cfn-glue-partition-catalogid
	//
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the catalog database in which to create the partition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-partition.html#cfn-glue-partition-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The structure used to create and update a partition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-partition.html#cfn-glue-partition-partitioninput
	//
	PartitionInput interface{} `field:"required" json:"partitionInput" yaml:"partitionInput"`
	// The name of the metadata table in which the partition is to be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-partition.html#cfn-glue-partition-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

