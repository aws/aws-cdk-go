package mixinsawsbedrock


// Contains details about the Lambda function containing the business logic that is carried out upon invoking the action or the custom control method for handling the information elicited from the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionGroupExecutorProperty := &ActionGroupExecutorProperty{
//   	CustomControl: jsii.String("customControl"),
//   	Lambda: jsii.String("lambda"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-actiongroupexecutor.html
//
type CfnAgentPropsMixin_ActionGroupExecutorProperty struct {
	// To return the action group invocation results directly in the `InvokeInlineAgent` response, specify `RETURN_CONTROL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-actiongroupexecutor.html#cfn-bedrock-agent-actiongroupexecutor-customcontrol
	//
	CustomControl *string `field:"optional" json:"customControl" yaml:"customControl"`
	// The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-actiongroupexecutor.html#cfn-bedrock-agent-actiongroupexecutor-lambda
	//
	Lambda *string `field:"optional" json:"lambda" yaml:"lambda"`
}

