package awsrds


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentTaskProperty := &DeploymentTaskProperty{
//   	TaskName: jsii.String("taskName"),
//   	TaskStatus: jsii.String("taskStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-deployment-deploymenttask.html
//
type CfnDeployment_DeploymentTaskProperty struct {
	// The name of a Blue/Green Deployment provisioning task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-deployment-deploymenttask.html#cfn-rds-deployment-deploymenttask-taskname
	//
	TaskName *string `field:"optional" json:"taskName" yaml:"taskName"`
	// The status of a Blue/Green Deployment provisioning task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-deployment-deploymenttask.html#cfn-rds-deployment-deploymenttask-taskstatus
	//
	TaskStatus *string `field:"optional" json:"taskStatus" yaml:"taskStatus"`
}

