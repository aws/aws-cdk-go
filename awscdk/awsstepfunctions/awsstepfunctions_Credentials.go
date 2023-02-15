package awsstepfunctions


// Specifies a target role assumed by the State Machine's execution role for invoking the task's resource.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var submitLambda function
//   var iamRole role
//
//
//   // use a fixed role for all task invocations
//   role := sfn.taskRole.fromRole(iamRole)
//   // or use a json expression to resolve the role at runtime based on task inputs
//   //const role = sfn.TaskRole.fromRoleArnJsonPath('$.RoleArn');
//
//   submitJob := tasks.NewLambdaInvoke(this, jsii.String("Submit Job"), &lambdaInvokeProps{
//   	lambdaFunction: submitLambda,
//   	outputPath: jsii.String("$.Payload"),
//   	// use credentials
//   	credentials: &credentials{
//   		role: role,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-task-state.html#task-state-fields
//
type Credentials struct {
	// The role to be assumed for executing the Task.
	Role TaskRole `field:"required" json:"role" yaml:"role"`
}

