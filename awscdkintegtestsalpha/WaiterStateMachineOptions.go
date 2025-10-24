package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for creating a WaiterStateMachine.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import lambdanodejs "github.com/aws/aws-cdk-go/awscdk"
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   app := cdk.NewApp(&AppProps{
//   	PostCliContext: map[string]interface{}{
//   		"@aws-cdk/aws-lambda:useCdkManagedLogGroup": jsii.Boolean(false),
//   	},
//   })
//
//   stack := cdk.NewStack(app, jsii.String("aws-cdk-firehose-delivery-stream-s3-all-properties"))
//
//   bucket := s3.NewBucket(stack, jsii.String("FirehoseDeliveryStreamS3AllPropertiesBucket"), &BucketProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//
//   backupBucket := s3.NewBucket(stack, jsii.String("FirehoseDeliveryStreamS3AllPropertiesBackupBucket"), &BucketProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//   logGroup := logs.NewLogGroup(stack, jsii.String("LogGroup"), &LogGroupProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   })
//
//   dataProcessorFunction := lambdanodejs.NewNodejsFunction(stack, jsii.String("DataProcessorFunction"), &NodejsFunctionProps{
//   	Entry: path.join(__dirname, jsii.String("lambda-data-processor.js")),
//   	Timeout: cdk.Duration_Minutes(jsii.Number(1)),
//   })
//
//   processor := firehose.NewLambdaFunctionProcessor(dataProcessorFunction, &DataProcessorProps{
//   	BufferInterval: cdk.Duration_Seconds(jsii.Number(60)),
//   	BufferSize: cdk.Size_Mebibytes(jsii.Number(1)),
//   	Retries: jsii.Number(1),
//   })
//
//   key := kms.NewKey(stack, jsii.String("Key"), &KeyProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   })
//
//   backupKey := kms.NewKey(stack, jsii.String("BackupKey"), &KeyProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   })
//
//   deliveryStream := firehose.NewDeliveryStream(stack, jsii.String("DeliveryStream"), &DeliveryStreamProps{
//   	Destination: firehose.NewS3Bucket(bucket, &S3BucketProps{
//   		LoggingConfig: firehose.NewEnableLogging(logGroup),
//   		Processor: processor,
//   		Compression: firehose.Compression_GZIP(),
//   		DataOutputPrefix: jsii.String("regularPrefix"),
//   		ErrorOutputPrefix: jsii.String("errorPrefix"),
//   		FileExtension: jsii.String(".log.gz"),
//   		TimeZone: cdk.TimeZone_ASIA_TOKYO(),
//   		BufferingInterval: cdk.Duration_*Seconds(jsii.Number(60)),
//   		BufferingSize: cdk.Size_*Mebibytes(jsii.Number(1)),
//   		EncryptionKey: key,
//   		S3Backup: &DestinationS3BackupProps{
//   			Mode: firehose.BackupMode_ALL,
//   			Bucket: backupBucket,
//   			Compression: firehose.Compression_ZIP(),
//   			DataOutputPrefix: jsii.String("backupPrefix"),
//   			ErrorOutputPrefix: jsii.String("backupErrorPrefix"),
//   			BufferingInterval: cdk.Duration_*Seconds(jsii.Number(60)),
//   			BufferingSize: cdk.Size_*Mebibytes(jsii.Number(1)),
//   			EncryptionKey: backupKey,
//   		},
//   	}),
//   })
//
//   firehose.NewDeliveryStream(stack, jsii.String("ZeroBufferingDeliveryStream"), &DeliveryStreamProps{
//   	Destination: firehose.NewS3Bucket(bucket, &S3BucketProps{
//   		Compression: firehose.Compression_GZIP(),
//   		DataOutputPrefix: jsii.String("regularPrefix"),
//   		ErrorOutputPrefix: jsii.String("errorPrefix"),
//   		BufferingInterval: cdk.Duration_*Seconds(jsii.Number(0)),
//   	}),
//   })
//
//   testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("integ-tests"), &IntegTestProps{
//   	TestCases: []Stack{
//   		stack,
//   	},
//   	Regions: []*string{
//   		jsii.String("us-east-1"),
//   	},
//   })
//
//   testCase.Assertions.AwsApiCall(jsii.String("Firehose"), jsii.String("putRecord"), map[string]interface{}{
//   	"DeliveryStreamName": deliveryStream.deliveryStreamName,
//   	"Record": map[string]*string{
//   		"Data": jsii.String("testData123"),
//   	},
//   })
//
//   s3ApiCall := testCase.Assertions.AwsApiCall(jsii.String("S3"), jsii.String("listObjectsV2"), map[string]interface{}{
//   	"Bucket": bucket.bucketName,
//   	"MaxKeys": jsii.Number(1),
//   }).Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
//   	"KeyCount": jsii.Number(1),
//   })).WaitForAssertions(&WaiterStateMachineOptions{
//   	Interval: cdk.Duration_*Seconds(jsii.Number(30)),
//   	TotalTimeout: cdk.Duration_*Minutes(jsii.Number(10)),
//   })
//
//   if s3ApiCall instanceof awscdkintegtestsalpha.AwsApiCall && s3ApiCall.WaiterProvider {
//   	s3ApiCall.WaiterProvider.AddToRolePolicy(map[string]interface{}{
//   		"Effect": jsii.String("Allow"),
//   		"Action": []*string{
//   			jsii.String("s3:GetObject"),
//   			jsii.String("s3:ListBucket"),
//   		},
//   		"Resource": []*string{
//   			jsii.String("*"),
//   		},
//   	})
//   }
//
// Experimental.
type WaiterStateMachineOptions struct {
	// Backoff between attempts.
	//
	// This is the multiplier by which the retry interval increases
	// after each retry attempt.
	//
	// By default there is no backoff. Each retry will wait the amount of time
	// specified by `interval`.
	// Default: 1 (no backoff).
	//
	// Experimental.
	BackoffRate *float64 `field:"optional" json:"backoffRate" yaml:"backoffRate"`
	// The interval (number of seconds) to wait between attempts.
	// Default: Duration.seconds(5)
	//
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The total time that the state machine will wait for a successful response.
	// Default: Duration.minutes(30)
	//
	// Experimental.
	TotalTimeout awscdk.Duration `field:"optional" json:"totalTimeout" yaml:"totalTimeout"`
}

