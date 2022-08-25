package awsec2


// Properties for defining a `CfnClientVpnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClientVpnEndpointProps := &cfnClientVpnEndpointProps{
//   	authenticationOptions: []interface{}{
//   		&clientAuthenticationRequestProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			activeDirectory: &directoryServiceAuthenticationRequestProperty{
//   				directoryId: jsii.String("directoryId"),
//   			},
//   			federatedAuthentication: &federatedAuthenticationRequestProperty{
//   				samlProviderArn: jsii.String("samlProviderArn"),
//
//   				// the properties below are optional
//   				selfServiceSamlProviderArn: jsii.String("selfServiceSamlProviderArn"),
//   			},
//   			mutualAuthentication: &certificateAuthenticationRequestProperty{
//   				clientRootCertificateChainArn: jsii.String("clientRootCertificateChainArn"),
//   			},
//   		},
//   	},
//   	clientCidrBlock: jsii.String("clientCidrBlock"),
//   	connectionLogOptions: &connectionLogOptionsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		cloudwatchLogGroup: jsii.String("cloudwatchLogGroup"),
//   		cloudwatchLogStream: jsii.String("cloudwatchLogStream"),
//   	},
//   	serverCertificateArn: jsii.String("serverCertificateArn"),
//
//   	// the properties below are optional
//   	clientConnectOptions: &clientConnectOptionsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	clientLoginBannerOptions: &clientLoginBannerOptionsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		bannerText: jsii.String("bannerText"),
//   	},
//   	description: jsii.String("description"),
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	selfServicePortal: jsii.String("selfServicePortal"),
//   	sessionTimeoutHours: jsii.Number(123),
//   	splitTunnel: jsii.Boolean(false),
//   	tagSpecifications: []interface{}{
//   		&tagSpecificationProperty{
//   			resourceType: jsii.String("resourceType"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	transportProtocol: jsii.String("transportProtocol"),
//   	vpcId: jsii.String("vpcId"),
//   	vpnPort: jsii.Number(123),
//   }
//
type CfnClientVpnEndpointProps struct {
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions interface{} `field:"required" json:"authenticationOptions" yaml:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses.
	//
	// The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock *string `field:"required" json:"clientCidrBlock" yaml:"clientCidrBlock"`
	// Information about the client connection logging options.
	//
	// If you enable client connection logging, data about client connections is sent to a Cloudwatch Logs log stream. The following information is logged:
	//
	// - Client connection requests
	// - Client connection results (successful and unsuccessful)
	// - Reasons for unsuccessful client connection requests
	// - Client connection termination time.
	ConnectionLogOptions interface{} `field:"required" json:"connectionLogOptions" yaml:"connectionLogOptions"`
	// The ARN of the server certificate.
	//
	// For more information, see the [AWS Certificate Manager User Guide](https://docs.aws.amazon.com/acm/latest/userguide/) .
	ServerCertificateArn *string `field:"required" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// The options for managing connection authorization for new client connections.
	ClientConnectOptions interface{} `field:"optional" json:"clientConnectOptions" yaml:"clientConnectOptions"`
	// Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
	ClientLoginBannerOptions interface{} `field:"optional" json:"clientLoginBannerOptions" yaml:"clientLoginBannerOptions"`
	// A brief description of the Client VPN endpoint.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about the DNS servers to be used for DNS resolution.
	//
	// A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address configured on the device is used for the DNS server.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// The IDs of one or more security groups to apply to the target network.
	//
	// You must also specify the ID of the VPC that contains the security groups.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Specify whether to enable the self-service portal for the Client VPN endpoint.
	//
	// Default Value: `enabled`.
	SelfServicePortal *string `field:"optional" json:"selfServicePortal" yaml:"selfServicePortal"`
	// The maximum VPN session duration time in hours.
	//
	// Valid values: `8 | 10 | 12 | 24`
	//
	// Default value: `24`.
	SessionTimeoutHours *float64 `field:"optional" json:"sessionTimeoutHours" yaml:"sessionTimeoutHours"`
	// Indicates whether split-tunnel is enabled on the AWS Client VPN endpoint.
	//
	// By default, split-tunnel on a VPN endpoint is disabled.
	//
	// For information about split-tunnel VPN endpoints, see [Split-tunnel AWS Client VPN endpoint](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/split-tunnel-vpn.html) in the *AWS Client VPN Administrator Guide* .
	SplitTunnel interface{} `field:"optional" json:"splitTunnel" yaml:"splitTunnel"`
	// The tags to apply to the Client VPN endpoint during creation.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The transport protocol to be used by the VPN session.
	//
	// Default value: `udp`.
	TransportProtocol *string `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
	// The ID of the VPC to associate with the Client VPN endpoint.
	//
	// If no security group IDs are specified in the request, the default security group for the VPC is applied.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The port number to assign to the Client VPN endpoint for TCP and UDP traffic.
	//
	// Valid Values: `443` | `1194`
	//
	// Default Value: `443`.
	VpnPort *float64 `field:"optional" json:"vpnPort" yaml:"vpnPort"`
}

