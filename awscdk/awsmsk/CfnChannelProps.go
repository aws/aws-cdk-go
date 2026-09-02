package awsmsk


// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelProps := &CfnChannelProps{
//   	ChannelName: jsii.String("channelName"),
//   	TopicConfigurationList: []interface{}{
//   		&TopicConfigurationProperty{
//   			RecordConverter: &RecordConverterProperty{
//   				ValueConverter: jsii.String("valueConverter"),
//   			},
//   			TopicArn: jsii.String("topicArn"),
//
//   			// the properties below are optional
//   			RecordSchema: &RecordSchemaProperty{
//   				GsrArn: jsii.String("gsrArn"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ClusterArn: jsii.String("clusterArn"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	IcebergDestinationConfiguration: &IcebergDestinationConfigurationProperty{
//   		AppendOnly: jsii.Boolean(false),
//   		DeadLetterQueueS3: &DeadLetterQueueS3Property{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//
//   			// the properties below are optional
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		DestinationTableList: []interface{}{
//   			&DestinationTableProperty{
//   				DestinationDatabaseName: jsii.String("destinationDatabaseName"),
//   				DestinationTableName: jsii.String("destinationTableName"),
//
//   				// the properties below are optional
//   				PartitionSpec: &PartitionSpecProperty{
//   					PartitionStrategy: jsii.String("partitionStrategy"),
//
//   					// the properties below are optional
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
//
//   		// the properties below are optional
//   		Catalog: &CatalogProperty{
//   			CatalogArn: jsii.String("catalogArn"),
//   			WarehouseLocation: jsii.String("warehouseLocation"),
//   		},
//   		CompressionType: jsii.String("compressionType"),
//   		DataFreshnessInSeconds: jsii.Number(123),
//   	},
//   	LoggingInfo: &ChannelLoggingInfoProperty{
//   		CloudWatchLogs: &CloudWatchLogsLogDestinationProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		Firehose: &FirehoseLogDestinationProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			DeliveryStream: jsii.String("deliveryStream"),
//   		},
//   		S3: &S3LogDestinationProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			Bucket: jsii.String("bucket"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		DeadLetterQueueS3: &DeadLetterQueueS3Property{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//
//   			// the properties below are optional
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   		Storage: &S3StorageProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			CompressionType: jsii.String("compressionType"),
//   			StorageClass: jsii.String("storageClass"),
//
//   			// the properties below are optional
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   			OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   			OutputPrefix: jsii.String("outputPrefix"),
//   		},
//
//   		// the properties below are optional
//   		DataFreshnessInSeconds: jsii.Number(123),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html
//
type CfnChannelProps struct {
	// Name of the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-channelname
	//
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Topic configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html#cfn-msk-channel-topicconfigurationlist
	//
	TopicConfigurationList interface{} `field:"required" json:"topicConfigurationList" yaml:"topicConfigurationList"`
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
}

