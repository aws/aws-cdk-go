package awssagemaker


// Update policy for a blue/green deployment.
//
// If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blueGreenUpdatePolicyProperty := &blueGreenUpdatePolicyProperty{
//   	trafficRoutingConfiguration: &trafficRoutingConfigProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		canarySize: &capacitySizeProperty{
//   			type: jsii.String("type"),
//   			value: jsii.Number(123),
//   		},
//   		linearStepSize: &capacitySizeProperty{
//   			type: jsii.String("type"),
//   			value: jsii.Number(123),
//   		},
//   		waitIntervalInSeconds: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	maximumExecutionTimeoutInSeconds: jsii.Number(123),
//   	terminationWaitInSeconds: jsii.Number(123),
//   }
//
type CfnEndpoint_BlueGreenUpdatePolicyProperty struct {
	// Defines the traffic routing strategy to shift traffic from the old fleet to the new fleet during an endpoint deployment.
	TrafficRoutingConfiguration interface{} `field:"required" json:"trafficRoutingConfiguration" yaml:"trafficRoutingConfiguration"`
	// Maximum execution timeout for the deployment.
	//
	// Note that the timeout value should be larger than the total waiting time specified in `TerminationWaitInSeconds` and `WaitIntervalInSeconds` .
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// Additional waiting time in seconds after the completion of an endpoint deployment before terminating the old endpoint fleet.
	//
	// Default is 0.
	TerminationWaitInSeconds *float64 `field:"optional" json:"terminationWaitInSeconds" yaml:"terminationWaitInSeconds"`
}

