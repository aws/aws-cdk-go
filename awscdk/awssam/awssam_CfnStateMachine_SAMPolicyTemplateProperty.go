package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAMPolicyTemplateProperty := &sAMPolicyTemplateProperty{
//   	lambdaInvokePolicy: &functionSAMPTProperty{
//   		functionName: jsii.String("functionName"),
//   	},
//   	stepFunctionsExecutionPolicy: &stateMachineSAMPTProperty{
//   		stateMachineName: jsii.String("stateMachineName"),
//   	},
//   }
//
type CfnStateMachine_SAMPolicyTemplateProperty struct {
	// `CfnStateMachine.SAMPolicyTemplateProperty.LambdaInvokePolicy`.
	LambdaInvokePolicy interface{} `field:"optional" json:"lambdaInvokePolicy" yaml:"lambdaInvokePolicy"`
	// `CfnStateMachine.SAMPolicyTemplateProperty.StepFunctionsExecutionPolicy`.
	StepFunctionsExecutionPolicy interface{} `field:"optional" json:"stepFunctionsExecutionPolicy" yaml:"stepFunctionsExecutionPolicy"`
}

