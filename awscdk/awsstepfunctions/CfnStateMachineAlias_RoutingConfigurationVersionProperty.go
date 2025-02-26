package awsstepfunctions


// The state machine version to which you want to route the execution traffic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingConfigurationVersionProperty := &RoutingConfigurationVersionProperty{
//   	StateMachineVersionArn: jsii.String("stateMachineVersionArn"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-routingconfigurationversion.html
//
type CfnStateMachineAlias_RoutingConfigurationVersionProperty struct {
	// The Amazon Resource Name (ARN) that identifies one or two state machine versions defined in the routing configuration.
	//
	// If you specify the ARN of a second version, it must belong to the same state machine as the first version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-routingconfigurationversion.html#cfn-stepfunctions-statemachinealias-routingconfigurationversion-statemachineversionarn
	//
	StateMachineVersionArn *string `field:"required" json:"stateMachineVersionArn" yaml:"stateMachineVersionArn"`
	// The percentage of traffic you want to route to the state machine version.
	//
	// The sum of the weights in the routing configuration must be equal to 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-routingconfigurationversion.html#cfn-stepfunctions-statemachinealias-routingconfigurationversion-weight
	//
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

