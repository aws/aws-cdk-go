package awseks


// The configuration of your local Amazon EKS cluster on an AWS Outpost.
//
// Before creating a cluster on an Outpost, review [Creating a local cluster on an Outpost](https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-local-cluster-create.html) in the *Amazon EKS User Guide* . This API isn't available for Amazon EKS clusters on the AWS cloud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outpostConfigProperty := &outpostConfigProperty{
//   	controlPlaneInstanceType: jsii.String("controlPlaneInstanceType"),
//   	outpostArns: []*string{
//   		jsii.String("outpostArns"),
//   	},
//
//   	// the properties below are optional
//   	controlPlanePlacement: &controlPlanePlacementProperty{
//   		groupName: jsii.String("groupName"),
//   	},
//   }
//
type CfnCluster_OutpostConfigProperty struct {
	// The Amazon EC2 instance type that you want to use for your local Amazon EKS cluster on Outposts.
	//
	// Choose an instance type based on the number of nodes that your cluster will have. For more information, see [Capacity considerations](https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-capacity-considerations.html) in the *Amazon EKS User Guide* .
	//
	// The instance type that you specify is used for all Kubernetes control plane instances. The instance type can't be changed after cluster creation. The control plane is not automatically scaled by Amazon EKS.
	ControlPlaneInstanceType *string `field:"required" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// The ARN of the Outpost that you want to use for your local Amazon EKS cluster on Outposts.
	//
	// Only a single Outpost ARN is supported.
	OutpostArns *[]*string `field:"required" json:"outpostArns" yaml:"outpostArns"`
	// `CfnCluster.OutpostConfigProperty.ControlPlanePlacement`.
	ControlPlanePlacement interface{} `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
}

