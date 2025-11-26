package previewawseksmixins


// Indicates the current configuration of the compute capability on your EKS Auto Mode cluster.
//
// For example, if the capability is enabled or disabled. If the compute capability is enabled, EKS Auto Mode will create and delete EC2 Managed Instances in your AWS account. For more information, see EKS Auto Mode compute capability in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   computeConfigProperty := &ComputeConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   	NodePools: []*string{
//   		jsii.String("nodePools"),
//   	},
//   	NodeRoleArn: jsii.String("nodeRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-computeconfig.html
//
type CfnClusterPropsMixin_ComputeConfigProperty struct {
	// Request to enable or disable the compute capability on your EKS Auto Mode cluster.
	//
	// If the compute capability is enabled, EKS Auto Mode will create and delete EC2 Managed Instances in your AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-computeconfig.html#cfn-eks-cluster-computeconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Configuration for node pools that defines the compute resources for your EKS Auto Mode cluster.
	//
	// For more information, see EKS Auto Mode Node Pools in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-computeconfig.html#cfn-eks-cluster-computeconfig-nodepools
	//
	NodePools *[]*string `field:"optional" json:"nodePools" yaml:"nodePools"`
	// The ARN of the IAM Role EKS will assign to EC2 Managed Instances in your EKS Auto Mode cluster.
	//
	// This value cannot be changed after the compute capability of EKS Auto Mode is enabled. For more information, see the IAM Reference in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-computeconfig.html#cfn-eks-cluster-computeconfig-noderolearn
	//
	NodeRoleArn *string `field:"optional" json:"nodeRoleArn" yaml:"nodeRoleArn"`
}

