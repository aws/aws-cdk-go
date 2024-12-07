package awseks


// Todo: add description.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnCluster_ComputeConfigProperty struct {
	// Todo: add description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-computeconfig.html#cfn-eks-cluster-computeconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Todo: add description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-computeconfig.html#cfn-eks-cluster-computeconfig-nodepools
	//
	NodePools *[]*string `field:"optional" json:"nodePools" yaml:"nodePools"`
	// Todo: add description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-computeconfig.html#cfn-eks-cluster-computeconfig-noderolearn
	//
	NodeRoleArn *string `field:"optional" json:"nodeRoleArn" yaml:"nodeRoleArn"`
}

