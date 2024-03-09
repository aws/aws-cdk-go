package awskinesisfirehose


// The `ExtendedS3DestinationConfiguration` property type configures an Amazon S3 destination for an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extendedS3DestinationConfigurationProperty := &ExtendedS3DestinationConfigurationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	BufferingHints: &BufferingHintsProperty{
//   		IntervalInSeconds: jsii.Number(123),
//   		SizeInMBs: jsii.Number(123),
//   	},
//   	CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamName: jsii.String("logStreamName"),
//   	},
//   	CompressionFormat: jsii.String("compressionFormat"),
//   	CustomTimeZone: jsii.String("customTimeZone"),
//   	DataFormatConversionConfiguration: &DataFormatConversionConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		InputFormatConfiguration: &InputFormatConfigurationProperty{
//   			Deserializer: &DeserializerProperty{
//   				HiveJsonSerDe: &HiveJsonSerDeProperty{
//   					TimestampFormats: []*string{
//   						jsii.String("timestampFormats"),
//   					},
//   				},
//   				OpenXJsonSerDe: &OpenXJsonSerDeProperty{
//   					CaseInsensitive: jsii.Boolean(false),
//   					ColumnToJsonKeyMappings: map[string]*string{
//   						"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   					},
//   					ConvertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		OutputFormatConfiguration: &OutputFormatConfigurationProperty{
//   			Serializer: &SerializerProperty{
//   				OrcSerDe: &OrcSerDeProperty{
//   					BlockSizeBytes: jsii.Number(123),
//   					BloomFilterColumns: []*string{
//   						jsii.String("bloomFilterColumns"),
//   					},
//   					BloomFilterFalsePositiveProbability: jsii.Number(123),
//   					Compression: jsii.String("compression"),
//   					DictionaryKeyThreshold: jsii.Number(123),
//   					EnablePadding: jsii.Boolean(false),
//   					FormatVersion: jsii.String("formatVersion"),
//   					PaddingTolerance: jsii.Number(123),
//   					RowIndexStride: jsii.Number(123),
//   					StripeSizeBytes: jsii.Number(123),
//   				},
//   				ParquetSerDe: &ParquetSerDeProperty{
//   					BlockSizeBytes: jsii.Number(123),
//   					Compression: jsii.String("compression"),
//   					EnableDictionaryCompression: jsii.Boolean(false),
//   					MaxPaddingBytes: jsii.Number(123),
//   					PageSizeBytes: jsii.Number(123),
//   					WriterVersion: jsii.String("writerVersion"),
//   				},
//   			},
//   		},
//   		SchemaConfiguration: &SchemaConfigurationProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Region: jsii.String("region"),
//   			RoleArn: jsii.String("roleArn"),
//   			TableName: jsii.String("tableName"),
//   			VersionId: jsii.String("versionId"),
//   		},
//   	},
//   	DynamicPartitioningConfiguration: &DynamicPartitioningConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		RetryOptions: &RetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   	},
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   			AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   		},
//   		NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   	},
//   	ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   	FileExtension: jsii.String("fileExtension"),
//   	Prefix: jsii.String("prefix"),
//   	ProcessingConfiguration: &ProcessingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		Processors: []interface{}{
//   			&ProcessorProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Parameters: []interface{}{
//   					&ProcessorParameterProperty{
//   						ParameterName: jsii.String("parameterName"),
//   						ParameterValue: jsii.String("parameterValue"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	S3BackupConfiguration: &S3DestinationConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
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
//   	},
//   	S3BackupMode: jsii.String("s3BackupMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html
//
type CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket.
	//
	// For constraints, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The Amazon Resource Name (ARN) of the AWS credentials.
	//
	// For constraints, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The buffering option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-bufferinghints
	//
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// The Amazon CloudWatch logging options for your delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-cloudwatchloggingoptions
	//
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The compression format.
	//
	// If no value is specified, the default is `UNCOMPRESSED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-compressionformat
	//
	CompressionFormat *string `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	// The time zone you prefer.
	//
	// UTC is the default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-customtimezone
	//
	CustomTimeZone *string `field:"optional" json:"customTimeZone" yaml:"customTimeZone"`
	// The serializer, deserializer, and schema for converting data from the JSON format to the Parquet or ORC format before writing it to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-dataformatconversionconfiguration
	//
	DataFormatConversionConfiguration interface{} `field:"optional" json:"dataFormatConversionConfiguration" yaml:"dataFormatConversionConfiguration"`
	// The configuration of the dynamic partitioning mechanism that creates targeted data sets from the streaming data by partitioning it based on partition keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-dynamicpartitioningconfiguration
	//
	DynamicPartitioningConfiguration interface{} `field:"optional" json:"dynamicPartitioningConfiguration" yaml:"dynamicPartitioningConfiguration"`
	// The encryption configuration for the Kinesis Data Firehose delivery stream.
	//
	// The default value is `NoEncryption` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-erroroutputprefix
	//
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// Specify a file extension.
	//
	// It will override the default file extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-fileextension
	//
	FileExtension *string `field:"optional" json:"fileExtension" yaml:"fileExtension"`
	// The `YYYY/MM/DD/HH` time format prefix is automatically used for delivered Amazon S3 files.
	//
	// For more information, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The data processing configuration for the Kinesis Data Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-processingconfiguration
	//
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// The configuration for backup in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-s3backupconfiguration
	//
	S3BackupConfiguration interface{} `field:"optional" json:"s3BackupConfiguration" yaml:"s3BackupConfiguration"`
	// The Amazon S3 backup mode.
	//
	// After you create a delivery stream, you can update it to enable Amazon S3 backup if it is disabled. If backup is enabled, you can't update the delivery stream to disable it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-s3backupmode
	//
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
}

