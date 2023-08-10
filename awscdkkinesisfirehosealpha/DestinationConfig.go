package awscdkkinesisfirehosealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Kinesis Data Firehose delivery stream destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import kinesisfirehose_alpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var dependable iDependable
//
//   destinationConfig := &DestinationConfig{
//   	Dependables: []*iDependable{
//   		dependable,
//   	},
//   	ExtendedS3DestinationConfiguration: &ExtendedS3DestinationConfigurationProperty{
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
//   		DataFormatConversionConfiguration: &DataFormatConversionConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			InputFormatConfiguration: &InputFormatConfigurationProperty{
//   				Deserializer: &DeserializerProperty{
//   					HiveJsonSerDe: &HiveJsonSerDeProperty{
//   						TimestampFormats: []*string{
//   							jsii.String("timestampFormats"),
//   						},
//   					},
//   					OpenXJsonSerDe: &OpenXJsonSerDeProperty{
//   						CaseInsensitive: jsii.Boolean(false),
//   						ColumnToJsonKeyMappings: map[string]*string{
//   							"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   						},
//   						ConvertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			OutputFormatConfiguration: &OutputFormatConfigurationProperty{
//   				Serializer: &SerializerProperty{
//   					OrcSerDe: &OrcSerDeProperty{
//   						BlockSizeBytes: jsii.Number(123),
//   						BloomFilterColumns: []*string{
//   							jsii.String("bloomFilterColumns"),
//   						},
//   						BloomFilterFalsePositiveProbability: jsii.Number(123),
//   						Compression: jsii.String("compression"),
//   						DictionaryKeyThreshold: jsii.Number(123),
//   						EnablePadding: jsii.Boolean(false),
//   						FormatVersion: jsii.String("formatVersion"),
//   						PaddingTolerance: jsii.Number(123),
//   						RowIndexStride: jsii.Number(123),
//   						StripeSizeBytes: jsii.Number(123),
//   					},
//   					ParquetSerDe: &ParquetSerDeProperty{
//   						BlockSizeBytes: jsii.Number(123),
//   						Compression: jsii.String("compression"),
//   						EnableDictionaryCompression: jsii.Boolean(false),
//   						MaxPaddingBytes: jsii.Number(123),
//   						PageSizeBytes: jsii.Number(123),
//   						WriterVersion: jsii.String("writerVersion"),
//   					},
//   				},
//   			},
//   			SchemaConfiguration: &SchemaConfigurationProperty{
//   				CatalogId: jsii.String("catalogId"),
//   				DatabaseName: jsii.String("databaseName"),
//   				Region: jsii.String("region"),
//   				RoleArn: jsii.String("roleArn"),
//   				TableName: jsii.String("tableName"),
//   				VersionId: jsii.String("versionId"),
//   			},
//   		},
//   		DynamicPartitioningConfiguration: &DynamicPartitioningConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			RetryOptions: &RetryOptionsProperty{
//   				DurationInSeconds: jsii.Number(123),
//   			},
//   		},
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   				AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		Prefix: jsii.String("prefix"),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		S3BackupConfiguration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   }
//
// Experimental.
type DestinationConfig struct {
	// Any resources that were created by the destination when binding it to the stack that must be deployed before the delivery stream is deployed.
	// Experimental.
	Dependables *[]constructs.IDependable `field:"optional" json:"dependables" yaml:"dependables"`
	// S3 destination configuration properties.
	// Experimental.
	ExtendedS3DestinationConfiguration *awskinesisfirehose.CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty `field:"optional" json:"extendedS3DestinationConfiguration" yaml:"extendedS3DestinationConfiguration"`
}

