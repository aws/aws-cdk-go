package awskinesisanalyticsv2


// Describes a SQL-based Kinesis Data Analytics application's output configuration, in which you identify an in-application stream and a destination where you want the in-application stream data to be written.
//
// The destination can be a Kinesis data stream or a Kinesis Data Firehose delivery stream.
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
//   	},
//   	KinesisStreamsOutput: &KinesisStreamsOutputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   	},
//   	LambdaOutput: &LambdaOutputProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-output.html
//
type CfnApplicationOutput_OutputProperty struct {
	// Describes the data format when records are written to the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-output.html#cfn-kinesisanalyticsv2-applicationoutput-output-destinationschema
	//
	DestinationSchema interface{} `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Identifies a Kinesis Data Firehose delivery stream as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-output.html#cfn-kinesisanalyticsv2-applicationoutput-output-kinesisfirehoseoutput
	//
	KinesisFirehoseOutput interface{} `field:"optional" json:"kinesisFirehoseOutput" yaml:"kinesisFirehoseOutput"`
	// Identifies a Kinesis data stream as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-output.html#cfn-kinesisanalyticsv2-applicationoutput-output-kinesisstreamsoutput
	//
	KinesisStreamsOutput interface{} `field:"optional" json:"kinesisStreamsOutput" yaml:"kinesisStreamsOutput"`
	// Identifies an Amazon Lambda function as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-output.html#cfn-kinesisanalyticsv2-applicationoutput-output-lambdaoutput
	//
	LambdaOutput interface{} `field:"optional" json:"lambdaOutput" yaml:"lambdaOutput"`
	// The name of the in-application stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-output.html#cfn-kinesisanalyticsv2-applicationoutput-output-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

