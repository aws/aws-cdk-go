package awsbedrockagentcore


// The Lambda function configuration for code-based evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaEvaluatorConfigProperty := &LambdaEvaluatorConfigProperty{
//   	LambdaArn: jsii.String("lambdaArn"),
//
//   	// the properties below are optional
//   	LambdaTimeoutInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-lambdaevaluatorconfig.html
//
type CfnEvaluator_LambdaEvaluatorConfigProperty struct {
	// The ARN of the Lambda function used for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-lambdaevaluatorconfig.html#cfn-bedrockagentcore-evaluator-lambdaevaluatorconfig-lambdaarn
	//
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// The timeout in seconds for the Lambda function invocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-lambdaevaluatorconfig.html#cfn-bedrockagentcore-evaluator-lambdaevaluatorconfig-lambdatimeoutinseconds
	//
	LambdaTimeoutInSeconds *float64 `field:"optional" json:"lambdaTimeoutInSeconds" yaml:"lambdaTimeoutInSeconds"`
}

