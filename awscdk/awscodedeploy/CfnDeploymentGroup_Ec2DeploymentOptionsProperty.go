package awscodedeploy


// Options that control how CodeDeploy manages Amazon EC2 Auto Scaling groups during deployments to the deployment group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ec2DeploymentOptionsProperty := &Ec2DeploymentOptionsProperty{
//   	AutoScalingGroupProcessesToSuspend: []*string{
//   		jsii.String("autoScalingGroupProcessesToSuspend"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2deploymentoptions.html
//
type CfnDeploymentGroup_Ec2DeploymentOptionsProperty struct {
	// The Amazon EC2 Auto Scaling processes to suspend during an in-place deployment.
	//
	// CodeDeploy suspends these processes before it deploys to instances in an Auto Scaling group, and resumes them when the deployment succeeds, fails, or is stopped. Valid values are Terminate, AZRebalance, InstanceRefresh, and ScheduledActions. Process names are case sensitive. To stop suspending processes, remove the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2deploymentoptions.html#cfn-codedeploy-deploymentgroup-ec2deploymentoptions-autoscalinggroupprocessestosuspend
	//
	AutoScalingGroupProcessesToSuspend *[]*string `field:"optional" json:"autoScalingGroupProcessesToSuspend" yaml:"autoScalingGroupProcessesToSuspend"`
}

