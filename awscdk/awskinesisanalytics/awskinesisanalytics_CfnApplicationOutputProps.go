package awskinesisanalytics


// Properties for defining a `CfnApplicationOutput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationOutputProps := &cfnApplicationOutputProps{
//   	applicationName: jsii.String("applicationName"),
//   	output: &outputProperty{
//   		destinationSchema: &destinationSchemaProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//   		},
//
//   		// the properties below are optional
//   		kinesisFirehoseOutput: &kinesisFirehoseOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		kinesisStreamsOutput: &kinesisStreamsOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		lambdaOutput: &lambdaOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnApplicationOutputProps struct {
	// Name of the application to which you want to add the output configuration.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// An array of objects, each describing one output configuration.
	//
	// In the output configuration, you specify the name of an in-application stream, a destination (that is, an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an AWS Lambda function), and record the formation to use when writing to the destination.
	Output interface{} `field:"required" json:"output" yaml:"output"`
}

