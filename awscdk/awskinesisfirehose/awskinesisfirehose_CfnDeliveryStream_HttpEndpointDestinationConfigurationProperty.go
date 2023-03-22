package awskinesisfirehose


// Describes the configuration of the HTTP endpoint destination.
//
// Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including Datadog, MongoDB, and New Relic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpEndpointDestinationConfigurationProperty := &httpEndpointDestinationConfigurationProperty{
//   	endpointConfiguration: &httpEndpointConfigurationProperty{
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		accessKey: jsii.String("accessKey"),
//   		name: jsii.String("name"),
//   	},
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
//   	bufferingHints: &bufferingHintsProperty{
//   		intervalInSeconds: jsii.Number(123),
//   		sizeInMBs: jsii.Number(123),
//   	},
//   	cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		logGroupName: jsii.String("logGroupName"),
//   		logStreamName: jsii.String("logStreamName"),
//   	},
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
//   	requestConfiguration: &httpEndpointRequestConfigurationProperty{
//   		commonAttributes: []interface{}{
//   			&httpEndpointCommonAttributeProperty{
//   				attributeName: jsii.String("attributeName"),
//   				attributeValue: jsii.String("attributeValue"),
//   			},
//   		},
//   		contentEncoding: jsii.String("contentEncoding"),
//   	},
//   	retryOptions: &retryOptionsProperty{
//   		durationInSeconds: jsii.Number(123),
//   	},
//   	roleArn: jsii.String("roleArn"),
//   	s3BackupMode: jsii.String("s3BackupMode"),
//   }
//
type CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty struct {
	// The configuration of the HTTP endpoint selected as the destination.
	EndpointConfiguration interface{} `field:"required" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Describes the configuration of a destination in Amazon S3.
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// The buffering options that can be used before data is delivered to the specified destination.
	//
	// Kinesis Data Firehose treats these options as hints, and it might choose to use more optimal values. The SizeInMBs and IntervalInSeconds parameters are optional. However, if you specify a value for one of them, you must also provide a value for the other.
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Describes the Amazon CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// Describes the data processing configuration.
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// The configuration of the request sent to the HTTP endpoint specified as the destination.
	RequestConfiguration interface{} `field:"optional" json:"requestConfiguration" yaml:"requestConfiguration"`
	// Describes the retry behavior in case Kinesis Data Firehose is unable to deliver data to the specified HTTP endpoint destination, or if it doesn't receive a valid acknowledgment of receipt from the specified HTTP endpoint destination.
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Kinesis Data Firehose uses this IAM role for all the permissions that the delivery stream needs.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Describes the S3 bucket backup options for the data that Kinesis Data Firehose delivers to the HTTP endpoint destination.
	//
	// You can back up all documents (AllData) or only the documents that Kinesis Data Firehose could not deliver to the specified HTTP endpoint destination (FailedDataOnly).
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
}

