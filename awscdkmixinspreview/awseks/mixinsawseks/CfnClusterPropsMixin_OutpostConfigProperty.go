package mixinsawseks


// The configuration of your local Amazon EKS cluster on an AWS Outpost.
//
// Before creating a cluster on an Outpost, review [Creating a local cluster on an Outpost](https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-local-cluster-create.html) in the *Amazon EKS User Guide* . This API isn't available for Amazon EKS clusters on the AWS cloud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outpostConfigProperty := &OutpostConfigProperty{
//   	ControlPlaneInstanceType: jsii.String("controlPlaneInstanceType"),
//   	ControlPlanePlacement: &ControlPlanePlacementProperty{
//   		GroupName: jsii.String("groupName"),
//   	},
//   	OutpostArns: []*string{
//   		jsii.String("outpostArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-outpostconfig.html
//
type CfnClusterPropsMixin_OutpostConfigProperty struct {
	// The Amazon EC2 instance type that you want to use for your local Amazon EKS cluster on Outposts.
	//
	// Choose an instance type based on the number of nodes that your cluster will have. For more information, see [Capacity considerations](https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-capacity-considerations.html) in the *Amazon EKS User Guide* .
	//
	// The instance type that you specify is used for all Kubernetes control plane instances. The instance type can't be changed after cluster creation. The control plane is not automatically scaled by Amazon EKS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-outpostconfig.html#cfn-eks-cluster-outpostconfig-controlplaneinstancetype
	//
	ControlPlaneInstanceType *string `field:"optional" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// An object representing the placement configuration for all the control plane instances of your local Amazon EKS cluster on an AWS Outpost.
	//
	// For more information, see [Capacity considerations](https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-capacity-considerations.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-outpostconfig.html#cfn-eks-cluster-outpostconfig-controlplaneplacement
	//
	ControlPlanePlacement interface{} `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
	// The ARN of the Outpost that you want to use for your local Amazon EKS cluster on Outposts.
	//
	// Only a single Outpost ARN is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-outpostconfig.html#cfn-eks-cluster-outpostconfig-outpostarns
	//
	OutpostArns *[]*string `field:"optional" json:"outpostArns" yaml:"outpostArns"`
}

