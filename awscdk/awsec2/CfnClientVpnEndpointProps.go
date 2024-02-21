package awsec2


// Properties for defining a `CfnClientVpnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClientVpnEndpointProps := &CfnClientVpnEndpointProps{
//   	AuthenticationOptions: []interface{}{
//   		&ClientAuthenticationRequestProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			ActiveDirectory: &DirectoryServiceAuthenticationRequestProperty{
//   				DirectoryId: jsii.String("directoryId"),
//   			},
//   			FederatedAuthentication: &FederatedAuthenticationRequestProperty{
//   				SamlProviderArn: jsii.String("samlProviderArn"),
//
//   				// the properties below are optional
//   				SelfServiceSamlProviderArn: jsii.String("selfServiceSamlProviderArn"),
//   			},
//   			MutualAuthentication: &CertificateAuthenticationRequestProperty{
//   				ClientRootCertificateChainArn: jsii.String("clientRootCertificateChainArn"),
//   			},
//   		},
//   	},
//   	ClientCidrBlock: jsii.String("clientCidrBlock"),
//   	ConnectionLogOptions: &ConnectionLogOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		CloudwatchLogGroup: jsii.String("cloudwatchLogGroup"),
//   		CloudwatchLogStream: jsii.String("cloudwatchLogStream"),
//   	},
//   	ServerCertificateArn: jsii.String("serverCertificateArn"),
//
//   	// the properties below are optional
//   	ClientConnectOptions: &ClientConnectOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	ClientLoginBannerOptions: &ClientLoginBannerOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		BannerText: jsii.String("bannerText"),
//   	},
//   	Description: jsii.String("description"),
//   	DnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SelfServicePortal: jsii.String("selfServicePortal"),
//   	SessionTimeoutHours: jsii.Number(123),
//   	SplitTunnel: jsii.Boolean(false),
//   	TagSpecifications: []interface{}{
//   		&TagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []cfnTag{
//   				&cfnTag{
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
type CfnClientVpnEndpointProps struct {
	// Information about the authentication method to be used to authenticate clients.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-authenticationoptions
	//
	AuthenticationOptions interface{} `field:"required" json:"authenticationOptions" yaml:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses.
	//
	// The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. Client CIDR range must have a size of at least /22 and must not be greater than /12.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-clientcidrblock
	//
	ClientCidrBlock *string `field:"required" json:"clientCidrBlock" yaml:"clientCidrBlock"`
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
	ConnectionLogOptions interface{} `field:"required" json:"connectionLogOptions" yaml:"connectionLogOptions"`
	// The ARN of the server certificate.
	//
	// For more information, see the [AWS Certificate Manager User Guide](https://docs.aws.amazon.com/acm/latest/userguide/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-servercertificatearn
	//
	ServerCertificateArn *string `field:"required" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// The options for managing connection authorization for new client connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-clientconnectoptions
	//
	ClientConnectOptions interface{} `field:"optional" json:"clientConnectOptions" yaml:"clientConnectOptions"`
	// Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-clientloginbanneroptions
	//
	ClientLoginBannerOptions interface{} `field:"optional" json:"clientLoginBannerOptions" yaml:"clientLoginBannerOptions"`
	// A brief description of the Client VPN endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
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

