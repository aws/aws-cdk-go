package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentLifecycleHookProperty := &DeploymentLifecycleHookProperty{
//   	HookTargetArn: jsii.String("hookTargetArn"),
//   	LifecycleStages: []*string{
//   		jsii.String("lifecycleStages"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentlifecyclehook.html
//
type CfnService_DeploymentLifecycleHookProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentlifecyclehook.html#cfn-ecs-service-deploymentlifecyclehook-hooktargetarn
	//
	HookTargetArn *string `field:"required" json:"hookTargetArn" yaml:"hookTargetArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentlifecyclehook.html#cfn-ecs-service-deploymentlifecyclehook-lifecyclestages
	//
	LifecycleStages *[]*string `field:"required" json:"lifecycleStages" yaml:"lifecycleStages"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentlifecyclehook.html#cfn-ecs-service-deploymentlifecyclehook-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

