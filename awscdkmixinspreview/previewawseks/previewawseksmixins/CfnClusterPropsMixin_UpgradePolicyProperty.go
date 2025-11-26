package previewawseksmixins


// The support policy to use for the cluster.
//
// Extended support allows you to remain on specific Kubernetes versions for longer. Clusters in extended support have higher costs. The default value is `EXTENDED` . Use `STANDARD` to disable extended support.
//
// [Learn more about EKS Extended Support in the *Amazon EKS User Guide* .](https://docs.aws.amazon.com/eks/latest/userguide/extended-support-control.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   upgradePolicyProperty := &UpgradePolicyProperty{
//   	SupportType: jsii.String("supportType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-upgradepolicy.html
//
type CfnClusterPropsMixin_UpgradePolicyProperty struct {
	// If the cluster is set to `EXTENDED` , it will enter extended support at the end of standard support.
	//
	// If the cluster is set to `STANDARD` , it will be automatically upgraded at the end of standard support.
	//
	// [Learn more about EKS Extended Support in the *Amazon EKS User Guide* .](https://docs.aws.amazon.com/eks/latest/userguide/extended-support-control.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-upgradepolicy.html#cfn-eks-cluster-upgradepolicy-supporttype
	//
	SupportType *string `field:"optional" json:"supportType" yaml:"supportType"`
}

