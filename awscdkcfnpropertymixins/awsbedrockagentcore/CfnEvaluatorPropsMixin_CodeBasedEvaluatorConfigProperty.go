package awsbedrockagentcore


// The configuration for code-based evaluation using a Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   codeBasedEvaluatorConfigProperty := &CodeBasedEvaluatorConfigProperty{
//   	LambdaConfig: &LambdaEvaluatorConfigProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   		LambdaTimeoutInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-codebasedevaluatorconfig.html
//
type CfnEvaluatorPropsMixin_CodeBasedEvaluatorConfigProperty struct {
	// The Lambda function configuration for code-based evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-codebasedevaluatorconfig.html#cfn-bedrockagentcore-evaluator-codebasedevaluatorconfig-lambdaconfig
	//
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
}

