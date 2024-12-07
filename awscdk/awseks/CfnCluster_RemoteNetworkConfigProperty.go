package awseks


// Configuration fields for specifying on-premises node and pod CIDRs that are external to the VPC passed during cluster creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remoteNetworkConfigProperty := &RemoteNetworkConfigProperty{
//   	RemoteNodeNetworks: []interface{}{
//   		&RemoteNodeNetworkProperty{
//   			Cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	RemotePodNetworks: []interface{}{
//   		&RemotePodNetworkProperty{
//   			Cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-remotenetworkconfig.html
//
type CfnCluster_RemoteNetworkConfigProperty struct {
	// Network configuration of nodes run on-premises with EKS Hybrid Nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-remotenetworkconfig.html#cfn-eks-cluster-remotenetworkconfig-remotenodenetworks
	//
	RemoteNodeNetworks interface{} `field:"required" json:"remoteNodeNetworks" yaml:"remoteNodeNetworks"`
	// Network configuration of pods run on-premises with EKS Hybrid Nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-remotenetworkconfig.html#cfn-eks-cluster-remotenetworkconfig-remotepodnetworks
	//
	RemotePodNetworks interface{} `field:"optional" json:"remotePodNetworks" yaml:"remotePodNetworks"`
}

