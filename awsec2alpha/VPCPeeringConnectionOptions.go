package awsec2alpha


// Options to define a VPC peering connection.
//
// Example:
//   stack := awscdk.Newstack()
//
//   acceptorVpc := awsec2alpha.NewVpcV2(this, jsii.String("VpcA"), &VpcV2Props{
//   	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.0.0.0/16")),
//   })
//
//   requestorVpc := awsec2alpha.NewVpcV2(this, jsii.String("VpcB"), &VpcV2Props{
//   	PrimaryAddressBlock: awsec2alpha.IpAddresses_*Ipv4(jsii.String("10.1.0.0/16")),
//   })
//
//   peeringConnection := requestorVpc.CreatePeeringConnection(jsii.String("peeringConnection"), &VPCPeeringConnectionOptions{
//   	AcceptorVpc: acceptorVpc,
//   })
//
//   routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
//   	Vpc: requestorVpc,
//   })
//
//   routeTable.AddRoute(jsii.String("vpcPeeringRoute"), jsii.String("10.0.0.0/16"), map[string]iRouteTarget{
//   	"gateway": peeringConnection,
//   })
//
// Experimental.
type VPCPeeringConnectionOptions struct {
	// The VPC that is accepting the peering connection.
	// Experimental.
	AcceptorVpc IVpcV2 `field:"required" json:"acceptorVpc" yaml:"acceptorVpc"`
	// The role arn created in the acceptor account.
	// Default: - no peerRoleArn needed if not cross account connection.
	//
	// Experimental.
	PeerRoleArn *string `field:"optional" json:"peerRoleArn" yaml:"peerRoleArn"`
	// The resource name of the peering connection.
	// Default: - peering connection provisioned without any name.
	//
	// Experimental.
	VpcPeeringConnectionName *string `field:"optional" json:"vpcPeeringConnectionName" yaml:"vpcPeeringConnectionName"`
}

