package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDeliveryStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeliveryStreamProps := &cfnDeliveryStreamProps{
//   	amazonOpenSearchServerlessDestinationConfiguration: &amazonOpenSearchServerlessDestinationConfigurationProperty{
//   		indexName: jsii.String("indexName"),
//   		roleArn: jsii.String("roleArn"),
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		bufferingHints: &amazonOpenSearchServerlessBufferingHintsProperty{
//   			intervalInSeconds: jsii.Number(123),
//   			sizeInMBs: jsii.Number(123),
//   		},
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		collectionEndpoint: jsii.String("collectionEndpoint"),
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		retryOptions: &amazonOpenSearchServerlessRetryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   		vpcConfiguration: &vpcConfigurationProperty{
//   			roleArn: jsii.String("roleArn"),
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	amazonopensearchserviceDestinationConfiguration: &amazonopensearchserviceDestinationConfigurationProperty{
//   		indexName: jsii.String("indexName"),
//   		roleArn: jsii.String("roleArn"),
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		bufferingHints: &amazonopensearchserviceBufferingHintsProperty{
//   			intervalInSeconds: jsii.Number(123),
//   			sizeInMBs: jsii.Number(123),
//   		},
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		clusterEndpoint: jsii.String("clusterEndpoint"),
//   		domainArn: jsii.String("domainArn"),
//   		indexRotationPeriod: jsii.String("indexRotationPeriod"),
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		retryOptions: &amazonopensearchserviceRetryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   		typeName: jsii.String("typeName"),
//   		vpcConfiguration: &vpcConfigurationProperty{
//   			roleArn: jsii.String("roleArn"),
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	deliveryStreamEncryptionConfigurationInput: &deliveryStreamEncryptionConfigurationInputProperty{
//   		keyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		keyArn: jsii.String("keyArn"),
//   	},
//   	deliveryStreamName: jsii.String("deliveryStreamName"),
//   	deliveryStreamType: jsii.String("deliveryStreamType"),
//   	elasticsearchDestinationConfiguration: &elasticsearchDestinationConfigurationProperty{
//   		indexName: jsii.String("indexName"),
//   		roleArn: jsii.String("roleArn"),
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		bufferingHints: &elasticsearchBufferingHintsProperty{
//   			intervalInSeconds: jsii.Number(123),
//   			sizeInMBs: jsii.Number(123),
//   		},
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		clusterEndpoint: jsii.String("clusterEndpoint"),
//   		domainArn: jsii.String("domainArn"),
//   		indexRotationPeriod: jsii.String("indexRotationPeriod"),
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		retryOptions: &elasticsearchRetryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   		typeName: jsii.String("typeName"),
//   		vpcConfiguration: &vpcConfigurationProperty{
//   			roleArn: jsii.String("roleArn"),
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	extendedS3DestinationConfiguration: &extendedS3DestinationConfigurationProperty{
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
//   		dataFormatConversionConfiguration: &dataFormatConversionConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			inputFormatConfiguration: &inputFormatConfigurationProperty{
//   				deserializer: &deserializerProperty{
//   					hiveJsonSerDe: &hiveJsonSerDeProperty{
//   						timestampFormats: []*string{
//   							jsii.String("timestampFormats"),
//   						},
//   					},
//   					openXJsonSerDe: &openXJsonSerDeProperty{
//   						caseInsensitive: jsii.Boolean(false),
//   						columnToJsonKeyMappings: map[string]*string{
//   							"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   						},
//   						convertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			outputFormatConfiguration: &outputFormatConfigurationProperty{
//   				serializer: &serializerProperty{
//   					orcSerDe: &orcSerDeProperty{
//   						blockSizeBytes: jsii.Number(123),
//   						bloomFilterColumns: []*string{
//   							jsii.String("bloomFilterColumns"),
//   						},
//   						bloomFilterFalsePositiveProbability: jsii.Number(123),
//   						compression: jsii.String("compression"),
//   						dictionaryKeyThreshold: jsii.Number(123),
//   						enablePadding: jsii.Boolean(false),
//   						formatVersion: jsii.String("formatVersion"),
//   						paddingTolerance: jsii.Number(123),
//   						rowIndexStride: jsii.Number(123),
//   						stripeSizeBytes: jsii.Number(123),
//   					},
//   					parquetSerDe: &parquetSerDeProperty{
//   						blockSizeBytes: jsii.Number(123),
//   						compression: jsii.String("compression"),
//   						enableDictionaryCompression: jsii.Boolean(false),
//   						maxPaddingBytes: jsii.Number(123),
//   						pageSizeBytes: jsii.Number(123),
//   						writerVersion: jsii.String("writerVersion"),
//   					},
//   				},
//   			},
//   			schemaConfiguration: &schemaConfigurationProperty{
//   				catalogId: jsii.String("catalogId"),
//   				databaseName: jsii.String("databaseName"),
//   				region: jsii.String("region"),
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//   				versionId: jsii.String("versionId"),
//   			},
//   		},
//   		dynamicPartitioningConfiguration: &dynamicPartitioningConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			retryOptions: &retryOptionsProperty{
//   				durationInSeconds: jsii.Number(123),
//   			},
//   		},
//   		encryptionConfiguration: &encryptionConfigurationProperty{
//   			kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   				awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		prefix: jsii.String("prefix"),
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		s3BackupConfiguration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	httpEndpointDestinationConfiguration: &httpEndpointDestinationConfigurationProperty{
//   		endpointConfiguration: &httpEndpointConfigurationProperty{
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			accessKey: jsii.String("accessKey"),
//   			name: jsii.String("name"),
//   		},
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
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
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		requestConfiguration: &httpEndpointRequestConfigurationProperty{
//   			commonAttributes: []interface{}{
//   				&httpEndpointCommonAttributeProperty{
//   					attributeName: jsii.String("attributeName"),
//   					attributeValue: jsii.String("attributeValue"),
//   				},
//   			},
//   			contentEncoding: jsii.String("contentEncoding"),
//   		},
//   		retryOptions: &retryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		roleArn: jsii.String("roleArn"),
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	kinesisStreamSourceConfiguration: &kinesisStreamSourceConfigurationProperty{
//   		kinesisStreamArn: jsii.String("kinesisStreamArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	redshiftDestinationConfiguration: &redshiftDestinationConfigurationProperty{
//   		clusterJdbcurl: jsii.String("clusterJdbcurl"),
//   		copyCommand: &copyCommandProperty{
//   			dataTableName: jsii.String("dataTableName"),
//
//   			// the properties below are optional
//   			copyOptions: jsii.String("copyOptions"),
//   			dataTableColumns: jsii.String("dataTableColumns"),
//   		},
//   		password: jsii.String("password"),
//   		roleArn: jsii.String("roleArn"),
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//   		username: jsii.String("username"),
//
//   		// the properties below are optional
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		retryOptions: &redshiftRetryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		s3BackupConfiguration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	s3DestinationConfiguration: &s3DestinationConfigurationProperty{
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
//   	splunkDestinationConfiguration: &splunkDestinationConfigurationProperty{
//   		hecEndpoint: jsii.String("hecEndpoint"),
//   		hecEndpointType: jsii.String("hecEndpointType"),
//   		hecToken: jsii.String("hecToken"),
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		hecAcknowledgmentTimeoutInSeconds: jsii.Number(123),
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		retryOptions: &splunkRetryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeliveryStreamProps struct {
	// `AWS::KinesisFirehose::DeliveryStream.AmazonOpenSearchServerlessDestinationConfiguration`.
	AmazonOpenSearchServerlessDestinationConfiguration interface{} `field:"optional" json:"amazonOpenSearchServerlessDestinationConfiguration" yaml:"amazonOpenSearchServerlessDestinationConfiguration"`
	// The destination in Amazon OpenSearch Service.
	//
	// You can specify only one destination.
	AmazonopensearchserviceDestinationConfiguration interface{} `field:"optional" json:"amazonopensearchserviceDestinationConfiguration" yaml:"amazonopensearchserviceDestinationConfiguration"`
	// Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput interface{} `field:"optional" json:"deliveryStreamEncryptionConfigurationInput" yaml:"deliveryStreamEncryptionConfigurationInput"`
	// The name of the delivery stream.
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// The delivery stream type. This can be one of the following values:.
	//
	// - `DirectPut` : Provider applications access the delivery stream directly.
	// - `KinesisStreamAsSource` : The delivery stream uses a Kinesis data stream as a source.
	DeliveryStreamType *string `field:"optional" json:"deliveryStreamType" yaml:"deliveryStreamType"`
	// An Amazon ES destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon ES destination to an Amazon S3 or Amazon Redshift destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ElasticsearchDestinationConfiguration interface{} `field:"optional" json:"elasticsearchDestinationConfiguration" yaml:"elasticsearchDestinationConfiguration"`
	// An Amazon S3 destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Extended S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ExtendedS3DestinationConfiguration interface{} `field:"optional" json:"extendedS3DestinationConfiguration" yaml:"extendedS3DestinationConfiguration"`
	// Enables configuring Kinesis Firehose to deliver data to any HTTP endpoint destination.
	//
	// You can specify only one destination.
	HttpEndpointDestinationConfiguration interface{} `field:"optional" json:"httpEndpointDestinationConfiguration" yaml:"httpEndpointDestinationConfiguration"`
	// When a Kinesis stream is used as the source for the delivery stream, a [KinesisStreamSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html) containing the Kinesis stream ARN and the role ARN for the source stream.
	KinesisStreamSourceConfiguration interface{} `field:"optional" json:"kinesisStreamSourceConfiguration" yaml:"kinesisStreamSourceConfiguration"`
	// An Amazon Redshift destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Redshift destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	RedshiftDestinationConfiguration interface{} `field:"optional" json:"redshiftDestinationConfiguration" yaml:"redshiftDestinationConfiguration"`
	// The `S3DestinationConfiguration` property type specifies an Amazon Simple Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	S3DestinationConfiguration interface{} `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	// The configuration of a destination in Splunk for the delivery stream.
	SplunkDestinationConfiguration interface{} `field:"optional" json:"splunkDestinationConfiguration" yaml:"splunkDestinationConfiguration"`
	// A set of tags to assign to the delivery stream.
	//
	// A tag is a key-value pair that you can define and assign to AWS resources. Tags are metadata. For example, you can add friendly names and descriptions or other types of information that can help you distinguish the delivery stream. For more information about tags, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the AWS Billing and Cost Management User Guide.
	//
	// You can specify up to 50 tags when creating a delivery stream.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

