package awseks


// An object representing the Upgrade Policy to use for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   upgradePolicyProperty := &UpgradePolicyProperty{
//   	SupportType: jsii.String("supportType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-upgradepolicy.html
//
type CfnCluster_UpgradePolicyProperty struct {
	// Specify the support type for your cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-upgradepolicy.html#cfn-eks-cluster-upgradepolicy-supporttype
	//
	SupportType *string `field:"optional" json:"supportType" yaml:"supportType"`
}

