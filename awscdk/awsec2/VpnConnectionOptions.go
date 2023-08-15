package awsec2


// Example:
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &VpcProps{
//   	VpnConnections: map[string]vpnConnectionOptions{
//   		"dynamic": &vpnConnectionOptions{
//   			 // Dynamic routing (BGP)
//   			"ip": jsii.String("1.2.3.4"),
//   		},
//   		"static": &vpnConnectionOptions{
//   			 // Static routing
//   			"ip": jsii.String("4.5.6.7"),
//   			"staticRoutes": []*string{
//   				jsii.String("192.168.10.0/24"),
//   				jsii.String("192.168.20.0/24"),
//   			},
//   		},
//   	},
//   })
//
type VpnConnectionOptions struct {
	// The ip address of the customer gateway.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// The ASN of the customer gateway.
	// Default: 65000.
	//
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// The static routes to be routed from the VPN gateway to the customer gateway.
	// Default: Dynamic routing (BGP).
	//
	StaticRoutes *[]*string `field:"optional" json:"staticRoutes" yaml:"staticRoutes"`
	// The tunnel options for the VPN connection.
	//
	// At most two elements (one per tunnel).
	// Duplicates not allowed.
	// Default: Amazon generated tunnel options.
	//
	TunnelOptions *[]*VpnTunnelOption `field:"optional" json:"tunnelOptions" yaml:"tunnelOptions"`
}

