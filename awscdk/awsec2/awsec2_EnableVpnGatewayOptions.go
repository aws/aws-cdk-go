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
//   enableVpnGatewayOptions := &enableVpnGatewayOptions{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	amazonSideAsn: jsii.Number(123),
//   	vpnRoutePropagation: []subnetSelection{
//   		&subnetSelection{
//   			availabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			onePerAz: jsii.Boolean(false),
//   			subnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			subnetGroupName: jsii.String("subnetGroupName"),
//   			subnets: []iSubnet{
//   				subnet,
//   			},
//   			subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   		},
//   	},
//   }
//
type EnableVpnGatewayOptions struct {
	// Default type ipsec.1.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Explicitly specify an Asn or let aws pick an Asn for you.
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Provide an array of subnets where the route propagation should be added.
	VpnRoutePropagation *[]*SubnetSelection `field:"optional" json:"vpnRoutePropagation" yaml:"vpnRoutePropagation"`
}

