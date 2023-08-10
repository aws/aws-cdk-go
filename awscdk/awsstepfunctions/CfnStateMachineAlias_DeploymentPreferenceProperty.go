package awsstepfunctions


// Enables gradual state machine deployments.
//
// CloudFormation automatically shifts traffic from the version the alias currently points to, to a new state machine version that you specify.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentPreferenceProperty := &DeploymentPreferenceProperty{
//   	StateMachineVersionArn: jsii.String("stateMachineVersionArn"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Alarms: []*string{
//   		jsii.String("alarms"),
//   	},
//   	Interval: jsii.Number(123),
//   	Percentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html
//
type CfnStateMachineAlias_DeploymentPreferenceProperty struct {
	// The Amazon Resource Name (ARN) of the [`AWS::StepFunctions::StateMachineVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html) resource that will be the final version to which the alias points to when the traffic shifting is complete.
	//
	// While performing gradual deployments, you can only provide a single state machine version ARN. To explicitly set version weights in a CloudFormation template, use `RoutingConfiguration` instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-statemachineversionarn
	//
	StateMachineVersionArn *string `field:"required" json:"stateMachineVersionArn" yaml:"stateMachineVersionArn"`
	// The type of deployment you want to perform. You can specify one of the following types:.
	//
	// - `LINEAR` - Shifts traffic to the new version in equal increments with an equal number of minutes between each increment.
	//
	// For example, if you specify the increment percent as `20` with an interval of `600` minutes, this deployment increases traffic by 20 percent every 600 minutes until the new version receives 100 percent of the traffic. This deployment immediately rolls back the new version if any CloudWatch alarms are triggered.
	// - `ALL_AT_ONCE` - Shifts 100 percent of traffic to the new version immediately. CloudFormation monitors the new version and rolls it back automatically to the previous version if any CloudWatch alarms are triggered.
	// - `CANARY` - Shifts traffic in two increments.
	//
	// In the first increment, a small percentage of traffic, for example, 10 percent is shifted to the new version. In the second increment, before a specified time interval in seconds gets over, the remaining traffic is shifted to the new version. The shift to the new version for the remaining traffic takes place only if no CloudWatch alarms are triggered during the specified time interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A list of Amazon CloudWatch alarms to be monitored during the deployment.
	//
	// The deployment fails and rolls back if any of these alarms go into the `ALARM` state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-alarms
	//
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// The time in minutes between each traffic shifting increment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The percentage of traffic to shift to the new version in each increment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-percentage
	//
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

