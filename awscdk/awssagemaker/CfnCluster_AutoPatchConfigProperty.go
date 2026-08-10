package awssagemaker


// The configuration for automatic patching of the instance group.
//
// Enables workload-aware, patch-level AMI updates.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoPatchConfigProperty := &AutoPatchConfigProperty{
//   	PatchingStrategy: jsii.String("patchingStrategy"),
//
//   	// the properties below are optional
//   	DeploymentConfig: &DeploymentConfigProperty{
//   		AutoRollbackConfiguration: []interface{}{
//   			&AlarmDetailsProperty{
//   				AlarmName: jsii.String("alarmName"),
//   			},
//   		},
//   		RollingUpdatePolicy: &RollingUpdatePolicyProperty{
//   			MaximumBatchSize: &CapacitySizeConfigProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			RollbackMaximumBatchSize: &CapacitySizeConfigProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   		WaitIntervalInSeconds: jsii.Number(123),
//   	},
//   	PatchSchedule: &PatchScheduleProperty{
//   		NextPatchDate: jsii.String("nextPatchDate"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-autopatchconfig.html
//
type CfnCluster_AutoPatchConfigProperty struct {
	// The patching strategy that determines when and how instances are patched.
	//
	// WhenIdle patches instances as they become idle. WhenAllIdle patches all instances when they are all idle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-autopatchconfig.html#cfn-sagemaker-cluster-autopatchconfig-patchingstrategy
	//
	PatchingStrategy *string `field:"required" json:"patchingStrategy" yaml:"patchingStrategy"`
	// The configuration to use when updating the AMI versions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-autopatchconfig.html#cfn-sagemaker-cluster-autopatchconfig-deploymentconfig
	//
	DeploymentConfig interface{} `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The schedule configuration for automatic patching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-autopatchconfig.html#cfn-sagemaker-cluster-autopatchconfig-patchschedule
	//
	PatchSchedule interface{} `field:"optional" json:"patchSchedule" yaml:"patchSchedule"`
}

