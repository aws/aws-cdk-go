package previewawsstepfunctionsmixins


// Properties for CfnStateMachineVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineVersionMixinProps := &CfnStateMachineVersionMixinProps{
//   	Description: jsii.String("description"),
//   	StateMachineArn: jsii.String("stateMachineArn"),
//   	StateMachineRevisionId: jsii.String("stateMachineRevisionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html
//
type CfnStateMachineVersionMixinProps struct {
	// An optional description of the state machine version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html#cfn-stepfunctions-statemachineversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the state machine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html#cfn-stepfunctions-statemachineversion-statemachinearn
	//
	StateMachineArn *string `field:"optional" json:"stateMachineArn" yaml:"stateMachineArn"`
	// Identifier for a state machine revision, which is an immutable, read-only snapshot of a state machine’s definition and configuration.
	//
	// Only publish the state machine version if the current state machine's revision ID matches the specified ID. Use this option to avoid publishing a version if the state machine has changed since you last updated it.
	//
	// To specify the initial state machine revision, set the value as `INITIAL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html#cfn-stepfunctions-statemachineversion-statemachinerevisionid
	//
	StateMachineRevisionId *string `field:"optional" json:"stateMachineRevisionId" yaml:"stateMachineRevisionId"`
}

