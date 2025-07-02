package awssagemaker


// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigProperty := &DeploymentConfigProperty{
//   	AutoRollbackConfiguration: &AutoRollbackConfigProperty{
//   		Alarms: []interface{}{
//   			&AlarmProperty{
//   				AlarmName: jsii.String("alarmName"),
//   			},
//   		},
//   	},
//   	BlueGreenUpdatePolicy: &BlueGreenUpdatePolicyProperty{
//   		TrafficRoutingConfiguration: &TrafficRoutingConfigProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			CanarySize: &CapacitySizeProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//   			LinearStepSize: &CapacitySizeProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//   			WaitIntervalInSeconds: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		MaximumExecutionTimeoutInSeconds: jsii.Number(123),
//   		TerminationWaitInSeconds: jsii.Number(123),
//   	},
//   	RollingUpdatePolicy: &RollingUpdatePolicyProperty{
//   		MaximumBatchSize: &CapacitySizeProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   		WaitIntervalInSeconds: jsii.Number(123),
//
//   		// the properties below are optional
//   		MaximumExecutionTimeoutInSeconds: jsii.Number(123),
//   		RollbackMaximumBatchSize: &CapacitySizeProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-deploymentconfig.html
//
type CfnEndpoint_DeploymentConfigProperty struct {
	// Automatic rollback configuration for handling endpoint deployment failures and recovery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-deploymentconfig.html#cfn-sagemaker-endpoint-deploymentconfig-autorollbackconfiguration
	//
	AutoRollbackConfiguration interface{} `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// Update policy for a blue/green deployment.
	//
	// If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-deploymentconfig.html#cfn-sagemaker-endpoint-deploymentconfig-bluegreenupdatepolicy
	//
	BlueGreenUpdatePolicy interface{} `field:"optional" json:"blueGreenUpdatePolicy" yaml:"blueGreenUpdatePolicy"`
	// Specifies a rolling deployment strategy for updating a SageMaker endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-deploymentconfig.html#cfn-sagemaker-endpoint-deploymentconfig-rollingupdatepolicy
	//
	RollingUpdatePolicy interface{} `field:"optional" json:"rollingUpdatePolicy" yaml:"rollingUpdatePolicy"`
}

