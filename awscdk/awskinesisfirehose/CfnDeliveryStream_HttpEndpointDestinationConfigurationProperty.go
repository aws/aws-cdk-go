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
//   httpEndpointDestinationConfigurationProperty := &HttpEndpointDestinationConfigurationProperty{
//   	EndpointConfiguration: &HttpEndpointConfigurationProperty{
//   		Url: jsii.String("url"),
//
//   		// the properties below are optional
//   		AccessKey: jsii.String("accessKey"),
//   		Name: jsii.String("name"),
//   	},
//   	S3Configuration: &S3DestinationConfigurationProperty{
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
//   	RequestConfiguration: &HttpEndpointRequestConfigurationProperty{
//   		CommonAttributes: []interface{}{
//   			&HttpEndpointCommonAttributeProperty{
//   				AttributeName: jsii.String("attributeName"),
//   				AttributeValue: jsii.String("attributeValue"),
//   			},
//   		},
//   		ContentEncoding: jsii.String("contentEncoding"),
//   	},
//   	RetryOptions: &RetryOptionsProperty{
//   		DurationInSeconds: jsii.Number(123),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	S3BackupMode: jsii.String("s3BackupMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html
//
type CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty struct {
	// The configuration of the HTTP endpoint selected as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-endpointconfiguration
	//
	EndpointConfiguration interface{} `field:"required" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Describes the configuration of a destination in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// The buffering options that can be used before data is delivered to the specified destination.
	//
	// Kinesis Data Firehose treats these options as hints, and it might choose to use more optimal values. The SizeInMBs and IntervalInSeconds parameters are optional. However, if you specify a value for one of them, you must also provide a value for the other.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-bufferinghints
	//
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Describes the Amazon CloudWatch logging options for your delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-cloudwatchloggingoptions
	//
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// Describes the data processing configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-processingconfiguration
	//
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// The configuration of the request sent to the HTTP endpoint specified as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-requestconfiguration
	//
	RequestConfiguration interface{} `field:"optional" json:"requestConfiguration" yaml:"requestConfiguration"`
	// Describes the retry behavior in case Kinesis Data Firehose is unable to deliver data to the specified HTTP endpoint destination, or if it doesn't receive a valid acknowledgment of receipt from the specified HTTP endpoint destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-retryoptions
	//
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Kinesis Data Firehose uses this IAM role for all the permissions that the delivery stream needs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Describes the S3 bucket backup options for the data that Kinesis Data Firehose delivers to the HTTP endpoint destination.
	//
	// You can back up all documents (AllData) or only the documents that Kinesis Data Firehose could not deliver to the specified HTTP endpoint destination (FailedDataOnly).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-s3backupmode
	//
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
}

