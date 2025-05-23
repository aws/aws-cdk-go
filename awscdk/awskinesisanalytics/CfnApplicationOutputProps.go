package awskinesisanalytics


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
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		KinesisStreamsOutput: &KinesisStreamsOutputProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		LambdaOutput: &LambdaOutputProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html
//
type CfnApplicationOutputProps struct {
	// Name of the application to which you want to add the output configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html#cfn-kinesisanalytics-applicationoutput-applicationname
	//
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// An array of objects, each describing one output configuration.
	//
	// In the output configuration, you specify the name of an in-application stream, a destination (that is, an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an AWS Lambda function), and record the formation to use when writing to the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html#cfn-kinesisanalytics-applicationoutput-output
	//
	Output interface{} `field:"required" json:"output" yaml:"output"`
}

