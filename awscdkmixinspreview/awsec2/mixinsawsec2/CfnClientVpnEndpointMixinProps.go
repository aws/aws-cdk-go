package mixinsawsec2


// Properties for CfnClientVpnEndpointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClientVpnEndpointMixinProps := &CfnClientVpnEndpointMixinProps{
//   	AuthenticationOptions: []interface{}{
//   		&ClientAuthenticationRequestProperty{
//   			ActiveDirectory: &DirectoryServiceAuthenticationRequestProperty{
//   				DirectoryId: jsii.String("directoryId"),
//   			},
//   			FederatedAuthentication: &FederatedAuthenticationRequestProperty{
//   				SamlProviderArn: jsii.String("samlProviderArn"),
//   				SelfServiceSamlProviderArn: jsii.String("selfServiceSamlProviderArn"),
//   			},
//   			MutualAuthentication: &CertificateAuthenticationRequestProperty{
//   				ClientRootCertificateChainArn: jsii.String("clientRootCertificateChainArn"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ClientCidrBlock: jsii.String("clientCidrBlock"),
//   	ClientConnectOptions: &ClientConnectOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	ClientLoginBannerOptions: &ClientLoginBannerOptionsProperty{
//   		BannerText: jsii.String("bannerText"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	ClientRouteEnforcementOptions: &ClientRouteEnforcementOptionsProperty{
//   		Enforced: jsii.Boolean(false),
//   	},
//   	ConnectionLogOptions: &ConnectionLogOptionsProperty{
//   		CloudwatchLogGroup: jsii.String("cloudwatchLogGroup"),
//   		CloudwatchLogStream: jsii.String("cloudwatchLogStream"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	DisconnectOnSessionTimeout: jsii.Boolean(false),
//   	DnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SelfServicePortal: jsii.String("selfServicePortal"),
//   	ServerCertificateArn: jsii.String("serverCertificateArn"),
//   	SessionTimeoutHours: jsii.Number(123),
//   	SplitTunnel: jsii.Boolean(false),
//   	TagSpecifications: []interface{}{
//   		&TagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	TransportProtocol: jsii.String("transportProtocol"),
//   	VpcId: jsii.String("vpcId"),
//   	VpnPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html
//
type CfnClientVpnEndpointMixinProps struct {
	// Information about the authentication method to be used to authenticate clients.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-authenticationoptions
	//
	AuthenticationOptions interface{} `field:"optional" json:"authenticationOptions" yaml:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses.
	//
	// The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. Client CIDR range must have a size of at least /22 and must not be greater than /12.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-clientcidrblock
	//
	ClientCidrBlock *string `field:"optional" json:"clientCidrBlock" yaml:"clientCidrBlock"`
	// The options for managing connection authorization for new client connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-clientconnectoptions
	//
	ClientConnectOptions interface{} `field:"optional" json:"clientConnectOptions" yaml:"clientConnectOptions"`
	// Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-clientloginbanneroptions
	//
	ClientLoginBannerOptions interface{} `field:"optional" json:"clientLoginBannerOptions" yaml:"clientLoginBannerOptions"`
	// Client route enforcement is a feature of the Client VPN service that helps enforce administrator defined routes on devices connected through the VPN.
	//
	// T his feature helps improve your security posture by ensuring that network traffic originating from a connected client is not inadvertently sent outside the VPN tunnel.
	//
	// Client route enforcement works by monitoring the route table of a connected device for routing policy changes to the VPN connection. If the feature detects any VPN routing policy modifications, it will automatically force an update to the route table, reverting it back to the expected route configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-clientrouteenforcementoptions
	//
	ClientRouteEnforcementOptions interface{} `field:"optional" json:"clientRouteEnforcementOptions" yaml:"clientRouteEnforcementOptions"`
	// Information about the client connection logging options.
	//
	// If you enable client connection logging, data about client connections is sent to a Cloudwatch Logs log stream. The following information is logged:
	//
	// - Client connection requests
	// - Client connection results (successful and unsuccessful)
	// - Reasons for unsuccessful client connection requests
	// - Client connection termination time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-connectionlogoptions
	//
	ConnectionLogOptions interface{} `field:"optional" json:"connectionLogOptions" yaml:"connectionLogOptions"`
	// A brief description of the Client VPN endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the client VPN session is disconnected after the maximum `sessionTimeoutHours` is reached.
	//
	// If `true` , users are prompted to reconnect client VPN. If `false` , client VPN attempts to reconnect automatically. The default value is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-disconnectonsessiontimeout
	//
	DisconnectOnSessionTimeout interface{} `field:"optional" json:"disconnectOnSessionTimeout" yaml:"disconnectOnSessionTimeout"`
	// Information about the DNS servers to be used for DNS resolution.
	//
	// A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address configured on the device is used for the DNS server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-dnsservers
	//
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// The IDs of one or more security groups to apply to the target network.
	//
	// You must also specify the ID of the VPC that contains the security groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Specify whether to enable the self-service portal for the Client VPN endpoint.
	//
	// Default Value: `enabled`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-selfserviceportal
	//
	SelfServicePortal *string `field:"optional" json:"selfServicePortal" yaml:"selfServicePortal"`
	// The ARN of the server certificate.
	//
	// For more information, see the [Certificate Manager User Guide](https://docs.aws.amazon.com/acm/latest/userguide/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-servercertificatearn
	//
	ServerCertificateArn *string `field:"optional" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// The maximum VPN session duration time in hours.
	//
	// Valid values: `8 | 10 | 12 | 24`
	//
	// Default value: `24`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-sessiontimeouthours
	//
	SessionTimeoutHours *float64 `field:"optional" json:"sessionTimeoutHours" yaml:"sessionTimeoutHours"`
	// Indicates whether split-tunnel is enabled on the AWS Client VPN endpoint.
	//
	// By default, split-tunnel on a VPN endpoint is disabled.
	//
	// For information about split-tunnel VPN endpoints, see [Split-tunnel AWS Client VPN endpoint](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/split-tunnel-vpn.html) in the *AWS Client VPN Administrator Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-splittunnel
	//
	SplitTunnel interface{} `field:"optional" json:"splitTunnel" yaml:"splitTunnel"`
	// The tags to apply to the Client VPN endpoint during creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-tagspecifications
	//
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The transport protocol to be used by the VPN session.
	//
	// Default value: `udp`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-transportprotocol
	//
	TransportProtocol *string `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
	// The ID of the VPC to associate with the Client VPN endpoint.
	//
	// If no security group IDs are specified in the request, the default security group for the VPC is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The port number to assign to the Client VPN endpoint for TCP and UDP traffic.
	//
	// Valid Values: `443` | `1194`
	//
	// Default Value: `443`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-vpnport
	//
	VpnPort *float64 `field:"optional" json:"vpnPort" yaml:"vpnPort"`
}

