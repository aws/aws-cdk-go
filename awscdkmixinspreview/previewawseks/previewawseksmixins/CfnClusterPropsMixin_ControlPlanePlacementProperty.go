package previewawseksmixins


// The placement configuration for all the control plane instances of your local Amazon EKS cluster on an AWS Outpost.
//
// For more information, see [Capacity considerations](https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-capacity-considerations.html) in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   controlPlanePlacementProperty := &ControlPlanePlacementProperty{
//   	GroupName: jsii.String("groupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-controlplaneplacement.html
//
type CfnClusterPropsMixin_ControlPlanePlacementProperty struct {
	// The name of the placement group for the Kubernetes control plane instances.
	//
	// This property is only used for a local cluster on an AWS Outpost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-controlplaneplacement.html#cfn-eks-cluster-controlplaneplacement-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
}

