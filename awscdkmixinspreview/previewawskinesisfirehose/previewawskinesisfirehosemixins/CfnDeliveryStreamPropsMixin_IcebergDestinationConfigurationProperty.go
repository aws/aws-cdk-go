package previewawskinesisfirehosemixins


// Specifies the destination configure settings for Apache Iceberg Table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   icebergDestinationConfigurationProperty := &IcebergDestinationConfigurationProperty{
//   	AppendOnly: jsii.Boolean(false),
//   	BufferingHints: &BufferingHintsProperty{
//   		IntervalInSeconds: jsii.Number(123),
//   		SizeInMBs: jsii.Number(123),
//   	},
//   	CatalogConfiguration: &CatalogConfigurationProperty{
//   		CatalogArn: jsii.String("catalogArn"),
//   		WarehouseLocation: jsii.String("warehouseLocation"),
//   	},
//   	CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamName: jsii.String("logStreamName"),
//   	},
//   	DestinationTableConfigurationList: []interface{}{
//   		&DestinationTableConfigurationProperty{
//   			DestinationDatabaseName: jsii.String("destinationDatabaseName"),
//   			DestinationTableName: jsii.String("destinationTableName"),
//   			PartitionSpec: &PartitionSpecProperty{
//   				Identity: []interface{}{
//   					&PartitionFieldProperty{
//   						SourceName: jsii.String("sourceName"),
//   					},
//   				},
//   			},
//   			S3ErrorOutputPrefix: jsii.String("s3ErrorOutputPrefix"),
//   			UniqueKeys: []*string{
//   				jsii.String("uniqueKeys"),
//   			},
//   		},
//   	},
//   	ProcessingConfiguration: &ProcessingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		Processors: []interface{}{
//   			&ProcessorProperty{
//   				Parameters: []interface{}{
//   					&ProcessorParameterProperty{
//   						ParameterName: jsii.String("parameterName"),
//   						ParameterValue: jsii.String("parameterValue"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	RetryOptions: &RetryOptionsProperty{
//   		DurationInSeconds: jsii.Number(123),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	S3BackupMode: jsii.String("s3BackupMode"),
//   	S3Configuration: &S3DestinationConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		BufferingHints: &BufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		CompressionFormat: jsii.String("compressionFormat"),
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   				AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		Prefix: jsii.String("prefix"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	SchemaEvolutionConfiguration: &SchemaEvolutionConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	TableCreationConfiguration: &TableCreationConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html
//
type CfnDeliveryStreamPropsMixin_IcebergDestinationConfigurationProperty struct {
	// Describes whether all incoming data for this delivery stream will be append only (inserts only and not for updates and deletes) for Iceberg delivery.
	//
	// This feature is only applicable for Apache Iceberg Tables.
	//
	// The default value is false. If you set this value to true, Firehose automatically increases the throughput limit of a stream based on the throttling levels of the stream. If you set this parameter to true for a stream with updates and deletes, you will see out of order delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-appendonly
	//
	AppendOnly interface{} `field:"optional" json:"appendOnly" yaml:"appendOnly"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-bufferinghints
	//
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Configuration describing where the destination Apache Iceberg Tables are persisted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-catalogconfiguration
	//
	CatalogConfiguration interface{} `field:"optional" json:"catalogConfiguration" yaml:"catalogConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-cloudwatchloggingoptions
	//
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// Provides a list of `DestinationTableConfigurations` which Firehose uses to deliver data to Apache Iceberg Tables.
	//
	// Firehose will write data with insert if table specific configuration is not provided here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-destinationtableconfigurationlist
	//
	DestinationTableConfigurationList interface{} `field:"optional" json:"destinationTableConfigurationList" yaml:"destinationTableConfigurationList"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-processingconfiguration
	//
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-retryoptions
	//
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// The Amazon Resource Name (ARN) of the IAM role to be assumed by Firehose for calling Apache Iceberg Tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Describes how Firehose will backup records.
	//
	// Currently,S3 backup only supports `FailedDataOnly` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-s3backupmode
	//
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// The configuration to enable automatic schema evolution.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-schemaevolutionconfiguration
	//
	SchemaEvolutionConfiguration interface{} `field:"optional" json:"schemaEvolutionConfiguration" yaml:"schemaEvolutionConfiguration"`
	// The configuration to enable automatic table creation.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration-tablecreationconfiguration
	//
	TableCreationConfiguration interface{} `field:"optional" json:"tableCreationConfiguration" yaml:"tableCreationConfiguration"`
}

