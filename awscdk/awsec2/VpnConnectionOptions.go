package awsec2


// Example:
//   // Across all tunnels in the account/region
//   allDataOut := ec2.VpnConnection_MetricAllTunnelDataOut()
//
//   // For a specific vpn connection
//   vpnConnection := vpc.addVpnConnection(jsii.String("Dynamic"), &VpnConnectionOptions{
//   	Ip: jsii.String("1.2.3.4"),
//   })
//   state := vpnConnection.metricTunnelState()
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

