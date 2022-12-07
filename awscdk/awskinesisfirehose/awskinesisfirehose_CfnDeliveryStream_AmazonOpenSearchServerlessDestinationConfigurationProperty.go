package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonOpenSearchServerlessDestinationConfigurationProperty := &amazonOpenSearchServerlessDestinationConfigurationProperty{
//   	indexName: jsii.String("indexName"),
//   	roleArn: jsii.String("roleArn"),
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
//   	bufferingHints: &amazonOpenSearchServerlessBufferingHintsProperty{
//   		intervalInSeconds: jsii.Number(123),
//   		sizeInMBs: jsii.Number(123),
//   	},
//   	cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		logGroupName: jsii.String("logGroupName"),
//   		logStreamName: jsii.String("logStreamName"),
//   	},
//   	collectionEndpoint: jsii.String("collectionEndpoint"),
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
//   	retryOptions: &amazonOpenSearchServerlessRetryOptionsProperty{
//   		durationInSeconds: jsii.Number(123),
//   	},
//   	s3BackupMode: jsii.String("s3BackupMode"),
//   	vpcConfiguration: &vpcConfigurationProperty{
//   		roleArn: jsii.String("roleArn"),
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDeliveryStream_AmazonOpenSearchServerlessDestinationConfigurationProperty struct {
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.IndexName`.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.RoleARN`.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.S3Configuration`.
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.BufferingHints`.
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.CloudWatchLoggingOptions`.
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.CollectionEndpoint`.
	CollectionEndpoint *string `field:"optional" json:"collectionEndpoint" yaml:"collectionEndpoint"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.ProcessingConfiguration`.
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.RetryOptions`.
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.S3BackupMode`.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.VpcConfiguration`.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

