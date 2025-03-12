package awsiot


// Starts execution of a Step Functions state machine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stepFunctionsActionProperty := &StepFunctionsActionProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	StateMachineName: jsii.String("stateMachineName"),
//
//   	// the properties below are optional
//   	ExecutionNamePrefix: jsii.String("executionNamePrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-stepfunctionsaction.html
//
type CfnTopicRule_StepFunctionsActionProperty struct {
	// The ARN of the role that grants IoT permission to start execution of a state machine ("Action":"states:StartExecution").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-stepfunctionsaction.html#cfn-iot-topicrule-stepfunctionsaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the Step Functions state machine whose execution will be started.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-stepfunctionsaction.html#cfn-iot-topicrule-stepfunctionsaction-statemachinename
	//
	StateMachineName *string `field:"required" json:"stateMachineName" yaml:"stateMachineName"`
	// (Optional) A name will be given to the state machine execution consisting of this prefix followed by a UUID.
	//
	// Step Functions automatically creates a unique name for each state machine execution if one is not provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-stepfunctionsaction.html#cfn-iot-topicrule-stepfunctionsaction-executionnameprefix
	//
	ExecutionNamePrefix *string `field:"optional" json:"executionNamePrefix" yaml:"executionNamePrefix"`
}

