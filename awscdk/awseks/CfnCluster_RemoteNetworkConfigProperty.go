package awseks


// The configuration in the cluster for EKS Hybrid Nodes.
//
// You can't change or update this configuration after the cluster is created.
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
	// The list of network CIDRs that can contain hybrid nodes.
	//
	// These CIDR blocks define the expected IP address range of the hybrid nodes that join the cluster. These blocks are typically determined by your network administrator.
	//
	// Enter one or more IPv4 CIDR blocks in decimal dotted-quad notation (for example, `10.2.0.0/16` ).
	//
	// It must satisfy the following requirements:
	//
	// - Each block must be within an `IPv4` RFC-1918 network range. Minimum allowed size is /24, maximum allowed size is /8. Publicly-routable addresses aren't supported.
	// - Each block cannot overlap with the range of the VPC CIDR blocks for your EKS resources, or the block of the Kubernetes service IP range.
	// - Each block must have a route to the VPC that uses the VPC CIDR blocks, not public IPs or Elastic IPs. There are many options including AWS Transit Gateway , AWS Site-to-Site VPN , or AWS Direct Connect .
	// - Each host must allow outbound connection to the EKS cluster control plane on TCP ports `443` and `10250` .
	// - Each host must allow inbound connection from the EKS cluster control plane on TCP port 10250 for logs, exec and port-forward operations.
	// - Each host must allow TCP and UDP network connectivity to and from other hosts that are running `CoreDNS` on UDP port `53` for service and pod DNS names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-remotenetworkconfig.html#cfn-eks-cluster-remotenetworkconfig-remotenodenetworks
	//
	RemoteNodeNetworks interface{} `field:"required" json:"remoteNodeNetworks" yaml:"remoteNodeNetworks"`
	// The list of network CIDRs that can contain pods that run Kubernetes webhooks on hybrid nodes.
	//
	// These CIDR blocks are determined by configuring your Container Network Interface (CNI) plugin. We recommend the Calico CNI or Cilium CNI. Note that the Amazon VPC CNI plugin for Kubernetes isn't available for on-premises and edge locations.
	//
	// Enter one or more IPv4 CIDR blocks in decimal dotted-quad notation (for example, `10.2.0.0/16` ).
	//
	// It must satisfy the following requirements:
	//
	// - Each block must be within an `IPv4` RFC-1918 network range. Minimum allowed size is /24, maximum allowed size is /8. Publicly-routable addresses aren't supported.
	// - Each block cannot overlap with the range of the VPC CIDR blocks for your EKS resources, or the block of the Kubernetes service IP range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-remotenetworkconfig.html#cfn-eks-cluster-remotenetworkconfig-remotepodnetworks
	//
	RemotePodNetworks interface{} `field:"optional" json:"remotePodNetworks" yaml:"remotePodNetworks"`
}

