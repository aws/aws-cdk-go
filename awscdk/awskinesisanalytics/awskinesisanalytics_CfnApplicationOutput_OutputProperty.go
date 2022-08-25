package awskinesisanalytics


// Describes application output configuration in which you identify an in-application stream and a destination where you want the in-application stream data to be written.
//
// The destination can be an Amazon Kinesis stream or an Amazon Kinesis Firehose delivery stream.
//
// For limits on how many destinations an application can write and other limitations, see [Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputProperty := &outputProperty{
//   	destinationSchema: &destinationSchemaProperty{
//   		recordFormatType: jsii.String("recordFormatType"),
//   	},
//
//   	// the properties below are optional
//   	kinesisFirehoseOutput: &kinesisFirehoseOutputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	kinesisStreamsOutput: &kinesisStreamsOutputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	lambdaOutput: &lambdaOutputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnApplicationOutput_OutputProperty struct {
	// Describes the data format when records are written to the destination.
	//
	// For more information, see [Configuring Application Output](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html) .
	DestinationSchema interface{} `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Identifies an Amazon Kinesis Firehose delivery stream as the destination.
	KinesisFirehoseOutput interface{} `field:"optional" json:"kinesisFirehoseOutput" yaml:"kinesisFirehoseOutput"`
	// Identifies an Amazon Kinesis stream as the destination.
	KinesisStreamsOutput interface{} `field:"optional" json:"kinesisStreamsOutput" yaml:"kinesisStreamsOutput"`
	// Identifies an AWS Lambda function as the destination.
	LambdaOutput interface{} `field:"optional" json:"lambdaOutput" yaml:"lambdaOutput"`
	// Name of the in-application stream.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

