package awskinesisanalytics


// Properties for defining a `CfnApplicationOutputV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationOutputV2Props := &cfnApplicationOutputV2Props{
//   	applicationName: jsii.String("applicationName"),
//   	output: &outputProperty{
//   		destinationSchema: &destinationSchemaProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//   		},
//
//   		// the properties below are optional
//   		kinesisFirehoseOutput: &kinesisFirehoseOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		kinesisStreamsOutput: &kinesisStreamsOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		lambdaOutput: &lambdaOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnApplicationOutputV2Props struct {
	// The name of the application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Describes a SQL-based Kinesis Data Analytics application's output configuration, in which you identify an in-application stream and a destination where you want the in-application stream data to be written.
	//
	// The destination can be a Kinesis data stream or a Kinesis Data Firehose delivery stream.
	Output interface{} `field:"required" json:"output" yaml:"output"`
}

