package awskinesisfirehose


// The `ExtendedS3DestinationConfiguration` property type configures an Amazon S3 destination for an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extendedS3DestinationConfigurationProperty := &extendedS3DestinationConfigurationProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	bufferingHints: &bufferingHintsProperty{
//   		intervalInSeconds: jsii.Number(123),
//   		sizeInMBs: jsii.Number(123),
//   	},
//   	cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		logGroupName: jsii.String("logGroupName"),
//   		logStreamName: jsii.String("logStreamName"),
//   	},
//   	compressionFormat: jsii.String("compressionFormat"),
//   	dataFormatConversionConfiguration: &dataFormatConversionConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   		inputFormatConfiguration: &inputFormatConfigurationProperty{
//   			deserializer: &deserializerProperty{
//   				hiveJsonSerDe: &hiveJsonSerDeProperty{
//   					timestampFormats: []*string{
//   						jsii.String("timestampFormats"),
//   					},
//   				},
//   				openXJsonSerDe: &openXJsonSerDeProperty{
//   					caseInsensitive: jsii.Boolean(false),
//   					columnToJsonKeyMappings: map[string]*string{
//   						"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   					},
//   					convertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		outputFormatConfiguration: &outputFormatConfigurationProperty{
//   			serializer: &serializerProperty{
//   				orcSerDe: &orcSerDeProperty{
//   					blockSizeBytes: jsii.Number(123),
//   					bloomFilterColumns: []*string{
//   						jsii.String("bloomFilterColumns"),
//   					},
//   					bloomFilterFalsePositiveProbability: jsii.Number(123),
//   					compression: jsii.String("compression"),
//   					dictionaryKeyThreshold: jsii.Number(123),
//   					enablePadding: jsii.Boolean(false),
//   					formatVersion: jsii.String("formatVersion"),
//   					paddingTolerance: jsii.Number(123),
//   					rowIndexStride: jsii.Number(123),
//   					stripeSizeBytes: jsii.Number(123),
//   				},
//   				parquetSerDe: &parquetSerDeProperty{
//   					blockSizeBytes: jsii.Number(123),
//   					compression: jsii.String("compression"),
//   					enableDictionaryCompression: jsii.Boolean(false),
//   					maxPaddingBytes: jsii.Number(123),
//   					pageSizeBytes: jsii.Number(123),
//   					writerVersion: jsii.String("writerVersion"),
//   				},
//   			},
//   		},
//   		schemaConfiguration: &schemaConfigurationProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			region: jsii.String("region"),
//   			roleArn: jsii.String("roleArn"),
//   			tableName: jsii.String("tableName"),
//   			versionId: jsii.String("versionId"),
//   		},
//   	},
//   	dynamicPartitioningConfiguration: &dynamicPartitioningConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   		retryOptions: &retryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   	},
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   			awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   		},
//   		noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   	},
//   	errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   	prefix: jsii.String("prefix"),
//   	processingConfiguration: &processingConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   		processors: []interface{}{
//   			&processorProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				parameters: []interface{}{
//   					&processorParameterProperty{
//   						parameterName: jsii.String("parameterName"),
//   						parameterValue: jsii.String("parameterValue"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	s3BackupConfiguration: &s3DestinationConfigurationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		bufferingHints: &bufferingHintsProperty{
//   			intervalInSeconds: jsii.Number(123),
//   			sizeInMBs: jsii.Number(123),
//   		},
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		compressionFormat: jsii.String("compressionFormat"),
//   		encryptionConfiguration: &encryptionConfigurationProperty{
//   			kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   				awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		prefix: jsii.String("prefix"),
//   	},
//   	s3BackupMode: jsii.String("s3BackupMode"),
//   }
//
type CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket.
	//
	// For constraints, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The Amazon Resource Name (ARN) of the AWS credentials.
	//
	// For constraints, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The buffering option.
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// The Amazon CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The compression format.
	//
	// If no value is specified, the default is `UNCOMPRESSED` .
	CompressionFormat *string `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	// The serializer, deserializer, and schema for converting data from the JSON format to the Parquet or ORC format before writing it to Amazon S3.
	DataFormatConversionConfiguration interface{} `field:"optional" json:"dataFormatConversionConfiguration" yaml:"dataFormatConversionConfiguration"`
	// The configuration of the dynamic partitioning mechanism that creates targeted data sets from the streaming data by partitioning it based on partition keys.
	DynamicPartitioningConfiguration interface{} `field:"optional" json:"dynamicPartitioningConfiguration" yaml:"dynamicPartitioningConfiguration"`
	// The encryption configuration for the Kinesis Data Firehose delivery stream.
	//
	// The default value is `NoEncryption` .
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html) .
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// The `YYYY/MM/DD/HH` time format prefix is automatically used for delivered Amazon S3 files.
	//
	// For more information, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The data processing configuration for the Kinesis Data Firehose delivery stream.
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// The configuration for backup in Amazon S3.
	S3BackupConfiguration interface{} `field:"optional" json:"s3BackupConfiguration" yaml:"s3BackupConfiguration"`
	// The Amazon S3 backup mode.
	//
	// After you create a delivery stream, you can update it to enable Amazon S3 backup if it is disabled. If backup is enabled, you can't update the delivery stream to disable it.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
}

