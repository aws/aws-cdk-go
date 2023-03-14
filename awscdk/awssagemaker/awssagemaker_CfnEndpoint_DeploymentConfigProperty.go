package awssagemaker


// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigProperty := &deploymentConfigProperty{
//   	blueGreenUpdatePolicy: &blueGreenUpdatePolicyProperty{
//   		trafficRoutingConfiguration: &trafficRoutingConfigProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			canarySize: &capacitySizeProperty{
//   				type: jsii.String("type"),
//   				value: jsii.Number(123),
//   			},
//   			linearStepSize: &capacitySizeProperty{
//   				type: jsii.String("type"),
//   				value: jsii.Number(123),
//   			},
//   			waitIntervalInSeconds: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		maximumExecutionTimeoutInSeconds: jsii.Number(123),
//   		terminationWaitInSeconds: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	autoRollbackConfiguration: &autoRollbackConfigProperty{
//   		alarms: []interface{}{
//   			&alarmProperty{
//   				alarmName: jsii.String("alarmName"),
//   			},
//   		},
//   	},
//   }
//
type CfnEndpoint_DeploymentConfigProperty struct {
	// Update policy for a blue/green deployment.
	//
	// If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default.
	BlueGreenUpdatePolicy interface{} `field:"required" json:"blueGreenUpdatePolicy" yaml:"blueGreenUpdatePolicy"`
	// Automatic rollback configuration for handling endpoint deployment failures and recovery.
	AutoRollbackConfiguration interface{} `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
}

