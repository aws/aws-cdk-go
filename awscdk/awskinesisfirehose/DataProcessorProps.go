package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configure the LambdaFunctionProcessor.
//
// Example:
//   var bucket Bucket
//   // Provide a Lambda function that will transform records before delivery, with custom
//   // buffering and retry configuration
//   lambdaFunction := lambda.NewFunction(this, jsii.String("Processor"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("process-records"))),
//   })
//   lambdaProcessor := firehose.NewLambdaFunctionProcessor(lambdaFunction, &DataProcessorProps{
//   	BufferInterval: awscdk.Duration_Minutes(jsii.Number(5)),
//   	BufferSize: awscdk.Size_Mebibytes(jsii.Number(5)),
//   	Retries: jsii.Number(5),
//   })
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	Processors: []IDataProcessor{
//   		lambdaProcessor,
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: s3Destination,
//   })
//
type DataProcessorProps struct {
	// The length of time Amazon Data Firehose will buffer incoming data before calling the processor.
	//
	// s.
	// Default: Duration.minutes(1)
	//
	BufferInterval awscdk.Duration `field:"optional" json:"bufferInterval" yaml:"bufferInterval"`
	// The amount of incoming data Amazon Data Firehose will buffer before calling the processor.
	// Default: Size.mebibytes(3)
	//
	BufferSize awscdk.Size `field:"optional" json:"bufferSize" yaml:"bufferSize"`
	// The number of times Amazon Data Firehose will retry the processor invocation after a failure due to network timeout or invocation limits.
	// Default: 3.
	//
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
}

