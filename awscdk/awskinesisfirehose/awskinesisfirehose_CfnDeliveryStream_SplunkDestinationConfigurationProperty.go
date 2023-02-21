package awskinesisfirehose


// The `SplunkDestinationConfiguration` property type specifies the configuration of a destination in Splunk for a Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   splunkDestinationConfigurationProperty := &splunkDestinationConfigurationProperty{
//   	hecEndpoint: jsii.String("hecEndpoint"),
//   	hecEndpointType: jsii.String("hecEndpointType"),
//   	hecToken: jsii.String("hecToken"),
//   	s3Configuration: &s3DestinationConfigurationProperty{
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
//
//   	// the properties below are optional
//   	cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		logGroupName: jsii.String("logGroupName"),
//   		logStreamName: jsii.String("logStreamName"),
//   	},
//   	hecAcknowledgmentTimeoutInSeconds: jsii.Number(123),
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
//   	retryOptions: &splunkRetryOptionsProperty{
//   		durationInSeconds: jsii.Number(123),
//   	},
//   	s3BackupMode: jsii.String("s3BackupMode"),
//   }
//
type CfnDeliveryStream_SplunkDestinationConfigurationProperty struct {
	// The HTTP Event Collector (HEC) endpoint to which Kinesis Data Firehose sends your data.
	HecEndpoint *string `field:"required" json:"hecEndpoint" yaml:"hecEndpoint"`
	// This type can be either `Raw` or `Event` .
	HecEndpointType *string `field:"required" json:"hecEndpointType" yaml:"hecEndpointType"`
	// This is a GUID that you obtain from your Splunk cluster when you create a new HEC endpoint.
	HecToken *string `field:"required" json:"hecToken" yaml:"hecToken"`
	// The configuration for the backup Amazon S3 location.
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// The Amazon CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The amount of time that Kinesis Data Firehose waits to receive an acknowledgment from Splunk after it sends it data.
	//
	// At the end of the timeout period, Kinesis Data Firehose either tries to send the data again or considers it an error, based on your retry settings.
	HecAcknowledgmentTimeoutInSeconds *float64 `field:"optional" json:"hecAcknowledgmentTimeoutInSeconds" yaml:"hecAcknowledgmentTimeoutInSeconds"`
	// The data processing configuration.
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// The retry behavior in case Kinesis Data Firehose is unable to deliver data to Splunk, or if it doesn't receive an acknowledgment of receipt from Splunk.
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Defines how documents should be delivered to Amazon S3.
	//
	// When set to `FailedEventsOnly` , Kinesis Data Firehose writes any data that could not be indexed to the configured Amazon S3 destination. When set to `AllEvents` , Kinesis Data Firehose delivers all incoming records to Amazon S3, and also writes failed documents to Amazon S3. The default value is `FailedEventsOnly` .
	//
	// You can update this backup mode from `FailedEventsOnly` to `AllEvents` . You can't update it from `AllEvents` to `FailedEventsOnly` .
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
}

