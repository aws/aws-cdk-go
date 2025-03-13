package awsec2alpha


// Properties to define an egress-only internet gateway.
//
// Example:
//   stack := awscdk.Newstack()
//   myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
//   	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.1.0.0/16")),
//   	SecondaryAddressBlocks: []iIpAddresses{
//   		awsec2alpha.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
//   			CidrBlockName: jsii.String("AmazonProvided"),
//   		}),
//   	},
//   })
//
//   eigw := awsec2alpha.NewEgressOnlyInternetGateway(this, jsii.String("EIGW"), &EgressOnlyInternetGatewayProps{
//   	Vpc: myVpc,
//   })
//
//   routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
//   	Vpc: myVpc,
//   })
//
//   routeTable.AddRoute(jsii.String("EIGW"), jsii.String("::/0"), map[string]iRouteTarget{
//   	"gateway": eigw,
//   })
//
// Experimental.
type EgressOnlyInternetGatewayProps struct {
	// The ID of the VPC for which to create the egress-only internet gateway.
	// Experimental.
	Vpc IVpcV2 `field:"required" json:"vpc" yaml:"vpc"`
	// The resource name of the egress-only internet gateway.
	// Default: - provisioned without a resource name.
	//
	// Experimental.
	EgressOnlyInternetGatewayName *string `field:"optional" json:"egressOnlyInternetGatewayName" yaml:"egressOnlyInternetGatewayName"`
}

