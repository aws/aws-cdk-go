package awsmsk


// Iceberg destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   icebergDestinationConfigurationProperty := &IcebergDestinationConfigurationProperty{
//   	AppendOnly: jsii.Boolean(false),
//   	Catalog: &CatalogProperty{
//   		CatalogArn: jsii.String("catalogArn"),
//   		WarehouseLocation: jsii.String("warehouseLocation"),
//   	},
//   	CompressionType: jsii.String("compressionType"),
//   	DataFreshnessInSeconds: jsii.Number(123),
//   	DeadLetterQueueS3: &DeadLetterQueueS3Property{
//   		BucketArn: jsii.String("bucketArn"),
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	},
//   	DestinationTableList: []interface{}{
//   		&DestinationTableProperty{
//   			DestinationDatabaseName: jsii.String("destinationDatabaseName"),
//   			DestinationTableName: jsii.String("destinationTableName"),
//   			PartitionSpec: &PartitionSpecProperty{
//   				PartitionStrategy: jsii.String("partitionStrategy"),
//   				SourceList: []interface{}{
//   					&PartitionSourceProperty{
//   						SourceName: jsii.String("sourceName"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SchemaEvolution: &SchemaEvolutionProperty{
//   		EnableSchemaEvolution: jsii.Boolean(false),
//   	},
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   	TableCreation: &TableCreationProperty{
//   		EnableTableCreation: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html
//
type CfnChannelPropsMixin_IcebergDestinationConfigurationProperty struct {
	// Append only mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html#cfn-msk-channel-icebergdestinationconfiguration-appendonly
	//
	// Default: - true.
	//
	AppendOnly interface{} `field:"optional" json:"appendOnly" yaml:"appendOnly"`
	// Catalog configuration of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html#cfn-msk-channel-icebergdestinationconfiguration-catalog
	//
	Catalog interface{} `field:"optional" json:"catalog" yaml:"catalog"`
	// Compression codec for Iceberg table data files.
	//
	// Defaults to ZSTD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html#cfn-msk-channel-icebergdestinationconfiguration-compressiontype
	//
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Data freshness in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html#cfn-msk-channel-icebergdestinationconfiguration-datafreshnessinseconds
	//
	DataFreshnessInSeconds *float64 `field:"optional" json:"dataFreshnessInSeconds" yaml:"dataFreshnessInSeconds"`
	// Dead letter queue S3 configuration of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html#cfn-msk-channel-icebergdestinationconfiguration-deadletterqueues3
	//
	DeadLetterQueueS3 interface{} `field:"optional" json:"deadLetterQueueS3" yaml:"deadLetterQueueS3"`
	// List of destination tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html#cfn-msk-channel-icebergdestinationconfiguration-destinationtablelist
	//
	DestinationTableList interface{} `field:"optional" json:"destinationTableList" yaml:"destinationTableList"`
	// Schema evolution configuration of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html#cfn-msk-channel-icebergdestinationconfiguration-schemaevolution
	//
	SchemaEvolution interface{} `field:"optional" json:"schemaEvolution" yaml:"schemaEvolution"`
	// The Amazon Resource Name (ARN) of an IAM role used by MSK to access the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html#cfn-msk-channel-icebergdestinationconfiguration-serviceexecutionrolearn
	//
	ServiceExecutionRoleArn *string `field:"optional" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// Table creation configuration of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-icebergdestinationconfiguration.html#cfn-msk-channel-icebergdestinationconfiguration-tablecreation
	//
	TableCreation interface{} `field:"optional" json:"tableCreation" yaml:"tableCreation"`
}

