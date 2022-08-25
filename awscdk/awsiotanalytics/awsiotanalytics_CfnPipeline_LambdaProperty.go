package awsiotanalytics


// An activity that runs a Lambda function to modify the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaProperty := &lambdaProperty{
//   	batchSize: jsii.Number(123),
//   	lambdaName: jsii.String("lambdaName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_LambdaProperty struct {
	// The number of messages passed to the Lambda function for processing.
	//
	// The AWS Lambda function must be able to process all of these messages within five minutes, which is the maximum timeout duration for Lambda functions.
	BatchSize *float64 `field:"required" json:"batchSize" yaml:"batchSize"`
	// The name of the Lambda function that is run on the message.
	LambdaName *string `field:"required" json:"lambdaName" yaml:"lambdaName"`
	// The name of the 'lambda' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

