package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionScalingConfigProperty := &FunctionScalingConfigProperty{
//   	MaxExecutionEnvironments: jsii.Number(123),
//   	MinExecutionEnvironments: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-functionscalingconfig.html
//
type CfnFunction_FunctionScalingConfigProperty struct {
	// The maximum number of execution environments that can be provisioned for the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-functionscalingconfig.html#cfn-lambda-function-functionscalingconfig-maxexecutionenvironments
	//
	MaxExecutionEnvironments *float64 `field:"optional" json:"maxExecutionEnvironments" yaml:"maxExecutionEnvironments"`
	// The minimum number of execution environments to maintain for the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-functionscalingconfig.html#cfn-lambda-function-functionscalingconfig-minexecutionenvironments
	//
	MinExecutionEnvironments *float64 `field:"optional" json:"minExecutionEnvironments" yaml:"minExecutionEnvironments"`
}

