package awsiotanalytics


// An activity that runs a Lambda function to modify the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaProperty := &LambdaProperty{
//   	BatchSize: jsii.Number(123),
//   	LambdaName: jsii.String("lambdaName"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Next: jsii.String("next"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-lambda.html
//
type CfnPipeline_LambdaProperty struct {
	// The number of messages passed to the Lambda function for processing.
	//
	// The AWS Lambda function must be able to process all of these messages within five minutes, which is the maximum timeout duration for Lambda functions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-lambda.html#cfn-iotanalytics-pipeline-lambda-batchsize
	//
	BatchSize *float64 `field:"required" json:"batchSize" yaml:"batchSize"`
	// The name of the Lambda function that is run on the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-lambda.html#cfn-iotanalytics-pipeline-lambda-lambdaname
	//
	LambdaName *string `field:"required" json:"lambdaName" yaml:"lambdaName"`
	// The name of the 'lambda' activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-lambda.html#cfn-iotanalytics-pipeline-lambda-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-lambda.html#cfn-iotanalytics-pipeline-lambda-next
	//
	Next *string `field:"optional" json:"next" yaml:"next"`
}

