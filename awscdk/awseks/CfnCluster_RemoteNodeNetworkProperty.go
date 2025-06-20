package awseks


// A network CIDR that can contain hybrid nodes.
//
// These CIDR blocks define the expected IP address range of the hybrid nodes that join the cluster. These blocks are typically determined by your network administrator.
//
// Enter one or more IPv4 CIDR blocks in decimal dotted-quad notation (for example, `10.2.0.0/16` ).
//
// It must satisfy the following requirements:
//
// - Each block must be within an `IPv4` RFC-1918 network range. Minimum allowed size is /32, maximum allowed size is /8. Publicly-routable addresses aren't supported.
// - Each block cannot overlap with the range of the VPC CIDR blocks for your EKS resources, or the block of the Kubernetes service IP range.
// - Each block must have a route to the VPC that uses the VPC CIDR blocks, not public IPs or Elastic IPs. There are many options including AWS Transit Gateway , AWS Site-to-Site VPN , or AWS Direct Connect .
// - Each host must allow outbound connection to the EKS cluster control plane on TCP ports `443` and `10250` .
// - Each host must allow inbound connection from the EKS cluster control plane on TCP port 10250 for logs, exec and port-forward operations.
// - Each host must allow TCP and UDP network connectivity to and from other hosts that are running `CoreDNS` on UDP port `53` for service and pod DNS names.
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
	// A network CIDR that can contain hybrid nodes.
	//
	// These CIDR blocks define the expected IP address range of the hybrid nodes that join the cluster. These blocks are typically determined by your network administrator.
	//
	// Enter one or more IPv4 CIDR blocks in decimal dotted-quad notation (for example, `10.2.0.0/16` ).
	//
	// It must satisfy the following requirements:
	//
	// - Each block must be within an `IPv4` RFC-1918 network range. Minimum allowed size is /32, maximum allowed size is /8. Publicly-routable addresses aren't supported.
	// - Each block cannot overlap with the range of the VPC CIDR blocks for your EKS resources, or the block of the Kubernetes service IP range.
	// - Each block must have a route to the VPC that uses the VPC CIDR blocks, not public IPs or Elastic IPs. There are many options including AWS Transit Gateway , AWS Site-to-Site VPN , or AWS Direct Connect .
	// - Each host must allow outbound connection to the EKS cluster control plane on TCP ports `443` and `10250` .
	// - Each host must allow inbound connection from the EKS cluster control plane on TCP port 10250 for logs, exec and port-forward operations.
	// - Each host must allow TCP and UDP network connectivity to and from other hosts that are running `CoreDNS` on UDP port `53` for service and pod DNS names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-remotenodenetwork.html#cfn-eks-cluster-remotenodenetwork-cidrs
	//
	Cidrs *[]*string `field:"required" json:"cidrs" yaml:"cidrs"`
}

