package awseks


// The placement configuration for all the control plane instances of your local Amazon EKS cluster on an AWS Outpost.
//
// For more information, see [Capacity considerations](https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-capacity-considerations.html) in the *Amazon EKS User Guide*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   controlPlanePlacementProperty := &ControlPlanePlacementProperty{
//   	GroupName: jsii.String("groupName"),
//   }
//
type CfnCluster_ControlPlanePlacementProperty struct {
	// The name of the placement group for the Kubernetes control plane instances.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
}

