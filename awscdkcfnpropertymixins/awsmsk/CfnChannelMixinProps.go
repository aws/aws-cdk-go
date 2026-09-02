package awsmsk


// Properties for CfnChannelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnChannelMixinProps := &CfnChannelMixinProps{
//   	ChannelName: jsii.String("channelName"),
//   	ClusterArn: jsii.String("clusterArn"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	IcebergDestinationConfiguration: &IcebergDestinationConfigurationProperty{
//   		AppendOnly: jsii.Boolean(false),
//   		Catalog: &CatalogProperty{
//   			CatalogArn: jsii.String("catalogArn"),
//   			WarehouseLocation: jsii.String("warehouseLocation"),
//   		},
//   		CompressionType: jsii.String("compressionType"),
//   		DataFreshnessInSeconds: jsii.Number(123),
//   		DeadLetterQueueS3: &DeadLetterQueueS3Property{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		DestinationTableList: []interface{}{
//   			&DestinationTableProperty{
//   				DestinationDatabaseName: jsii.String("destinationDatabaseName"),
//   				DestinationTableName: jsii.String("destinationTableName"),
//   				PartitionSpec: &PartitionSpecProperty{
//   					PartitionStrategy: jsii.String("partitionStrategy"),
//   					SourceList: []interface{}{
//   						&PartitionSourceProperty{
//   							SourceName: jsii.String("sourceName"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		SchemaEvolution: &SchemaEvolutionProperty{
//   			EnableSchemaEvolution: jsii.Boolean(false),
//   		},
//   		ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   		TableCreation: &TableCreationProperty{
//   			EnableTableCreation: jsii.Boolean(false),
//   		},
//   	},
//   	LoggingInfo: &ChannelLoggingInfoProperty{
//   		CloudWatchLogs: &CloudWatchLogsLogDestinationProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		Firehose: &FirehoseLogDestinationProperty{
//   			DeliveryStream: jsii.String("deliveryStream"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   		S3: &S3LogDestinationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Enabled: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		DataFreshnessInSeconds: jsii.Number(123),
//   		DeadLetterQueueS3: &DeadLetterQueueS3Property{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   		Storage: &S3StorageProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			CompressionType: jsii.String("compressionType"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   			OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   			OutputPrefix: jsii.String("outputPrefix"),
//   			StorageClass: jsii.String("storageClass"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TopicConfigurationList: []interface{}{
//   		&TopicConfigurationProperty{
//   			RecordConverter: &RecordConverterProperty{
//   				ValueConverter: jsii.String("valueConverter"),
//   			},
//   			RecordSchema: &RecordSchemaProperty{
//   				GsrArn: jsii.String("gsrArn"),
//   			},
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html
//
type CfnChannelMixinProps struct {
	// Name of the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The Amazon Resource Name (ARN) of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// Encryption configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Iceberg destination configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-icebergdestinationconfiguration
	//
	IcebergDestinationConfiguration interface{} `field:"optional" json:"icebergDestinationConfiguration" yaml:"icebergDestinationConfiguration"`
	// Log configuration details for Channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-logginginfo
	//
	LoggingInfo interface{} `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// S3 destination configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-s3destinationconfiguration
	//
	S3DestinationConfiguration interface{} `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	// Tags attached to the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Topic configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-topicconfigurationlist
	//
	TopicConfigurationList interface{} `field:"optional" json:"topicConfigurationList" yaml:"topicConfigurationList"`
}

