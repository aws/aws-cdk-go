package awsbedrock


// The structure of the executor invoking the actions in custom orchestration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   orchestrationExecutorProperty := &OrchestrationExecutorProperty{
//   	Lambda: jsii.String("lambda"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-orchestrationexecutor.html
//
type CfnAgent_OrchestrationExecutorProperty struct {
	// The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-orchestrationexecutor.html#cfn-bedrock-agent-orchestrationexecutor-lambda
	//
	Lambda *string `field:"required" json:"lambda" yaml:"lambda"`
}

