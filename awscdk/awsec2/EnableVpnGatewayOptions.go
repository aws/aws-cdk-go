package awsec2


// Options for the Vpc.enableVpnGateway() method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet subnet
//   var subnetFilter subnetFilter
//
//   enableVpnGatewayOptions := &EnableVpnGatewayOptions{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AmazonSideAsn: jsii.Number(123),
//   	VpnRoutePropagation: []subnetSelection{
//   		&subnetSelection{
//   			AvailabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			OnePerAz: jsii.Boolean(false),
//   			SubnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			SubnetGroupName: jsii.String("subnetGroupName"),
//   			Subnets: []iSubnet{
//   				subnet,
//   			},
//   			SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   		},
//   	},
//   }
//
type EnableVpnGatewayOptions struct {
	// Default type ipsec.1.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Explicitly specify an Asn or let aws pick an Asn for you.
	// Default: 65000.
	//
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Provide an array of subnets where the route propagation should be added.
	// Default: noPropagation.
	//
	VpnRoutePropagation *[]*SubnetSelection `field:"optional" json:"vpnRoutePropagation" yaml:"vpnRoutePropagation"`
}

