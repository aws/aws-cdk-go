package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   durableConfigProperty := &DurableConfigProperty{
//   	ExecutionTimeout: jsii.Number(123),
//
//   	// the properties below are optional
//   	RetentionPeriodInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-durableconfig.html
//
type CfnFunction_DurableConfigProperty struct {
	// The amount of time (in seconds) that Lambda allows a durable function to run before stopping it.
	//
	// The maximum is one 366-day year or 31,622,400 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-durableconfig.html#cfn-lambda-function-durableconfig-executiontimeout
	//
	ExecutionTimeout *float64 `field:"required" json:"executionTimeout" yaml:"executionTimeout"`
	// The number of days after a durable execution is closed that Lambda retains its history, from one to 90 days.
	//
	// The default is 14 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-durableconfig.html#cfn-lambda-function-durableconfig-retentionperiodindays
	//
	// Default: - 14.
	//
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
}

