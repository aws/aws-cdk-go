package awsec2alpha


// Additional props needed for secondary Address.
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
type SecondaryAddressProps struct {
	// Required to set Secondary cidr block resource name in order to generate unique logical id for the resource.
	// Experimental.
	CidrBlockName *string `field:"required" json:"cidrBlockName" yaml:"cidrBlockName"`
}

