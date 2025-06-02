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
//   outputProperty := &OutputProperty{
//   	DestinationSchema: &DestinationSchemaProperty{
//   		RecordFormatType: jsii.String("recordFormatType"),
//   	},
//
//   	// the properties below are optional
//   	KinesisFirehoseOutput: &KinesisFirehoseOutputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	KinesisStreamsOutput: &KinesisStreamsOutputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	LambdaOutput: &LambdaOutputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html
//
type CfnApplicationOutput_OutputProperty struct {
	// Describes the data format when records are written to the destination.
	//
	// For more information, see [Configuring Application Output](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-destinationschema
	//
	DestinationSchema interface{} `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Identifies an Amazon Kinesis Firehose delivery stream as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-kinesisfirehoseoutput
	//
	KinesisFirehoseOutput interface{} `field:"optional" json:"kinesisFirehoseOutput" yaml:"kinesisFirehoseOutput"`
	// Identifies an Amazon Kinesis stream as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-kinesisstreamsoutput
	//
	KinesisStreamsOutput interface{} `field:"optional" json:"kinesisStreamsOutput" yaml:"kinesisStreamsOutput"`
	// Identifies an AWS Lambda function as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-lambdaoutput
	//
	LambdaOutput interface{} `field:"optional" json:"lambdaOutput" yaml:"lambdaOutput"`
	// Name of the in-application stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

