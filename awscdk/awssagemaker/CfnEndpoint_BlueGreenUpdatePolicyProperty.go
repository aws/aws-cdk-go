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
//   blueGreenUpdatePolicyProperty := &BlueGreenUpdatePolicyProperty{
//   	TrafficRoutingConfiguration: &TrafficRoutingConfigProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		CanarySize: &CapacitySizeProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   		LinearStepSize: &CapacitySizeProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   		WaitIntervalInSeconds: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	MaximumExecutionTimeoutInSeconds: jsii.Number(123),
//   	TerminationWaitInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-bluegreenupdatepolicy.html
//
type CfnEndpoint_BlueGreenUpdatePolicyProperty struct {
	// Defines the traffic routing strategy to shift traffic from the old fleet to the new fleet during an endpoint deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-bluegreenupdatepolicy.html#cfn-sagemaker-endpoint-bluegreenupdatepolicy-trafficroutingconfiguration
	//
	TrafficRoutingConfiguration interface{} `field:"required" json:"trafficRoutingConfiguration" yaml:"trafficRoutingConfiguration"`
	// Maximum execution timeout for the deployment.
	//
	// Note that the timeout value should be larger than the total waiting time specified in `TerminationWaitInSeconds` and `WaitIntervalInSeconds` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-bluegreenupdatepolicy.html#cfn-sagemaker-endpoint-bluegreenupdatepolicy-maximumexecutiontimeoutinseconds
	//
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// Additional waiting time in seconds after the completion of an endpoint deployment before terminating the old endpoint fleet.
	//
	// Default is 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-bluegreenupdatepolicy.html#cfn-sagemaker-endpoint-bluegreenupdatepolicy-terminationwaitinseconds
	//
	TerminationWaitInSeconds *float64 `field:"optional" json:"terminationWaitInSeconds" yaml:"terminationWaitInSeconds"`
}

