package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAMPolicyTemplateProperty := &SAMPolicyTemplateProperty{
//   	LambdaInvokePolicy: &FunctionSAMPTProperty{
//   		FunctionName: jsii.String("functionName"),
//   	},
//   	StepFunctionsExecutionPolicy: &StateMachineSAMPTProperty{
//   		StateMachineName: jsii.String("stateMachineName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-sampolicytemplate.html
//
type CfnStateMachine_SAMPolicyTemplateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-sampolicytemplate.html#cfn-serverless-statemachine-sampolicytemplate-lambdainvokepolicy
	//
	LambdaInvokePolicy interface{} `field:"optional" json:"lambdaInvokePolicy" yaml:"lambdaInvokePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-sampolicytemplate.html#cfn-serverless-statemachine-sampolicytemplate-stepfunctionsexecutionpolicy
	//
	StepFunctionsExecutionPolicy interface{} `field:"optional" json:"stepFunctionsExecutionPolicy" yaml:"stepFunctionsExecutionPolicy"`
}

