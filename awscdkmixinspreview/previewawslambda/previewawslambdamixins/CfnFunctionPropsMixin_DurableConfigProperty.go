package previewawslambdamixins


// Configuration settings for [durable functions](https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html) , including execution timeout and retention period for execution history.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   durableConfigProperty := &DurableConfigProperty{
//   	ExecutionTimeout: jsii.Number(123),
//   	RetentionPeriodInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-durableconfig.html
//
type CfnFunctionPropsMixin_DurableConfigProperty struct {
	// The maximum time (in seconds) that a durable execution can run before timing out.
	//
	// This timeout applies to the entire durable execution, not individual function invocations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-durableconfig.html#cfn-lambda-function-durableconfig-executiontimeout
	//
	ExecutionTimeout *float64 `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// The number of days to retain execution history after a durable execution completes.
	//
	// After this period, execution history is no longer available through the GetDurableExecutionHistory API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-durableconfig.html#cfn-lambda-function-durableconfig-retentionperiodindays
	//
	// Default: - 14.
	//
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
}

