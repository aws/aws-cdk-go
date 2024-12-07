package awseks


// Network configuration of nodes run on-premises with EKS Hybrid Nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remoteNodeNetworkProperty := &RemoteNodeNetworkProperty{
//   	Cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-remotenodenetwork.html
//
type CfnCluster_RemoteNodeNetworkProperty struct {
	// Specifies the list of remote node CIDRs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-remotenodenetwork.html#cfn-eks-cluster-remotenodenetwork-cidrs
	//
	Cidrs *[]*string `field:"required" json:"cidrs" yaml:"cidrs"`
}

