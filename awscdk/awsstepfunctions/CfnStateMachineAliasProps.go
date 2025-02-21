package awsstepfunctions


// Properties for defining a `CfnStateMachineAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStateMachineAliasProps := &CfnStateMachineAliasProps{
//   	DeploymentPreference: &DeploymentPreferenceProperty{
//   		StateMachineVersionArn: jsii.String("stateMachineVersionArn"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		Interval: jsii.Number(123),
//   		Percentage: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RoutingConfiguration: []interface{}{
//   		&RoutingConfigurationVersionProperty{
//   			StateMachineVersionArn: jsii.String("stateMachineVersionArn"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html
//
type CfnStateMachineAliasProps struct {
	// The settings that enable gradual state machine deployments.
	//
	// These settings include [Alarms](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-alarms) , [Interval](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-interval) , [Percentage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-percentage) , [StateMachineVersionArn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-statemachineversionarn) , and [Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-type) .
	//
	// CloudFormation automatically shifts traffic from the version an alias currently points to, to a new state machine version that you specify.
	//
	// > `RoutingConfiguration` and `DeploymentPreference` are mutually exclusive properties. You must define only one of these properties.
	//
	// Based on the type of deployment you want to perform, you can specify one of the following settings:
	//
	// - `LINEAR` - Shifts traffic to the new version in equal increments with an equal number of minutes between each increment.
	//
	// For example, if you specify the increment percent as `20` with an interval of `600` minutes, this deployment increases traffic by 20 percent every 600 minutes until the new version receives 100 percent of the traffic. This deployment immediately rolls back the new version if any Amazon CloudWatch alarms are triggered.
	// - `ALL_AT_ONCE` - Shifts 100 percent of traffic to the new version immediately. CloudFormation monitors the new version and rolls it back automatically to the previous version if any CloudWatch alarms are triggered.
	// - `CANARY` - Shifts traffic in two increments.
	//
	// In the first increment, a small percentage of traffic, for example, 10 percent is shifted to the new version. In the second increment, before a specified time interval in seconds gets over, the remaining traffic is shifted to the new version. The shift to the new version for the remaining traffic takes place only if no CloudWatch alarms are triggered during the specified time interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html#cfn-stepfunctions-statemachinealias-deploymentpreference
	//
	DeploymentPreference interface{} `field:"optional" json:"deploymentPreference" yaml:"deploymentPreference"`
	// An optional description of the state machine alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html#cfn-stepfunctions-statemachinealias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the state machine alias.
	//
	// If you don't provide a name, it uses an automatically generated name based on the logical ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html#cfn-stepfunctions-statemachinealias-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The routing configuration of an alias.
	//
	// Routing configuration splits [StartExecution](https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html) requests between one or two versions of the same state machine.
	//
	// Use `RoutingConfiguration` if you want to explicitly set the alias [weights](https://docs.aws.amazon.com/step-functions/latest/apireference/API_RoutingConfigurationListItem.html#StepFunctions-Type-RoutingConfigurationListItem-weight) . Weight is the percentage of traffic you want to route to a state machine version.
	//
	// > `RoutingConfiguration` and `DeploymentPreference` are mutually exclusive properties. You must define only one of these properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html#cfn-stepfunctions-statemachinealias-routingconfiguration
	//
	RoutingConfiguration interface{} `field:"optional" json:"routingConfiguration" yaml:"routingConfiguration"`
}

