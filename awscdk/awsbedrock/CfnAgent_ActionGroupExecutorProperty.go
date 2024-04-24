package awsbedrock


// Contains details about the Lambda function containing the business logic that is carried out upon invoking the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionGroupExecutorProperty := &ActionGroupExecutorProperty{
//   	Lambda: jsii.String("lambda"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-actiongroupexecutor.html
//
type CfnAgent_ActionGroupExecutorProperty struct {
	// The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-actiongroupexecutor.html#cfn-bedrock-agent-actiongroupexecutor-lambda
	//
	Lambda *string `field:"required" json:"lambda" yaml:"lambda"`
}

