package awskinesisanalyticsv2


// Properties for defining a `CfnApplicationOutput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationOutputProps := &CfnApplicationOutputProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	Output: &OutputProperty{
//   		DestinationSchema: &DestinationSchemaProperty{
//   			RecordFormatType: jsii.String("recordFormatType"),
//   		},
//
//   		// the properties below are optional
//   		KinesisFirehoseOutput: &KinesisFirehoseOutputProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   		KinesisStreamsOutput: &KinesisStreamsOutputProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   		LambdaOutput: &LambdaOutputProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationoutput.html
//
type CfnApplicationOutputProps struct {
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationoutput.html#cfn-kinesisanalyticsv2-applicationoutput-applicationname
	//
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Describes a SQL-based Kinesis Data Analytics application's output configuration, in which you identify an in-application stream and a destination where you want the in-application stream data to be written.
	//
	// The destination can be a Kinesis data stream or a Kinesis Data Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationoutput.html#cfn-kinesisanalyticsv2-applicationoutput-output
	//
	Output interface{} `field:"required" json:"output" yaml:"output"`
}

