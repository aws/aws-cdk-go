package awskinesisanalytics


// Provides a description of a processor that is used to preprocess the records in the stream before being processed by your application code.
//
// Currently, the only input processor available is [AWS Lambda](https://docs.aws.amazon.com/lambda/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProcessingConfigurationProperty := &InputProcessingConfigurationProperty{
//   	InputLambdaProcessor: &InputLambdaProcessorProperty{
//   		ResourceArn: jsii.String("resourceArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html
//
type CfnApplication_InputProcessingConfigurationProperty struct {
	// The [InputLambdaProcessor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputlambdaprocessor.html) that is used to preprocess the records in the stream before being processed by your application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html#cfn-kinesisanalytics-application-inputprocessingconfiguration-inputlambdaprocessor
	//
	InputLambdaProcessor interface{} `field:"optional" json:"inputLambdaProcessor" yaml:"inputLambdaProcessor"`
}

