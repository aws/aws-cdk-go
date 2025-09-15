package awseks


// The update configuration for the node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updateConfigProperty := &UpdateConfigProperty{
//   	MaxUnavailable: jsii.Number(123),
//   	MaxUnavailablePercentage: jsii.Number(123),
//   	UpdateStrategy: jsii.String("updateStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-updateconfig.html
//
type CfnNodegroup_UpdateConfigProperty struct {
	// The maximum number of nodes unavailable at once during a version update.
	//
	// Nodes are updated in parallel. This value or `maxUnavailablePercentage` is required to have a value.The maximum number is 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-updateconfig.html#cfn-eks-nodegroup-updateconfig-maxunavailable
	//
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// The maximum percentage of nodes unavailable during a version update.
	//
	// This percentage of nodes are updated in parallel, up to 100 nodes at once. This value or `maxUnavailable` is required to have a value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-updateconfig.html#cfn-eks-nodegroup-updateconfig-maxunavailablepercentage
	//
	MaxUnavailablePercentage *float64 `field:"optional" json:"maxUnavailablePercentage" yaml:"maxUnavailablePercentage"`
	// The configuration for the behavior to follow during a node group version update of this managed node group.
	//
	// You choose between two possible strategies for replacing nodes during an [`UpdateNodegroupVersion`](https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateNodegroupVersion.html) action.
	//
	// An Amazon EKS managed node group updates by replacing nodes with new nodes of newer AMI versions in parallel. The *update strategy* changes the managed node update behavior of the managed node group for each quantity. The *default* strategy has guardrails to protect you from misconfiguration and launches the new instances first, before terminating the old instances. The *minimal* strategy removes the guardrails and terminates the old instances before launching the new instances. This minimal strategy is useful in scenarios where you are constrained to resources or costs (for example, with hardware accelerators such as GPUs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-updateconfig.html#cfn-eks-nodegroup-updateconfig-updatestrategy
	//
	UpdateStrategy *string `field:"optional" json:"updateStrategy" yaml:"updateStrategy"`
}

