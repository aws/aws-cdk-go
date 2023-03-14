package awskinesisanalyticsv2


// For an SQL-based Amazon Kinesis Data Analytics application, describes a processor that is used to preprocess the records in the stream before being processed by your application code.
//
// Currently, the only input processor available is [Amazon Lambda](https://docs.aws.amazon.com/lambda/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProcessingConfigurationProperty := &inputProcessingConfigurationProperty{
//   	inputLambdaProcessor: &inputLambdaProcessorProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   }
//
type CfnApplication_InputProcessingConfigurationProperty struct {
	// The [InputLambdaProcessor](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_InputLambdaProcessor.html) that is used to preprocess the records in the stream before being processed by your application code.
	InputLambdaProcessor interface{} `field:"optional" json:"inputLambdaProcessor" yaml:"inputLambdaProcessor"`
}

