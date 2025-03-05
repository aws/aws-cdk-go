package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for a client VPN endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var clientVpnConnectionHandler iClientVpnConnectionHandler
//   var clientVpnUserBasedAuthentication clientVpnUserBasedAuthentication
//   var logGroup logGroup
//   var logStream logStream
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   clientVpnEndpointProps := &ClientVpnEndpointProps{
//   	Cidr: jsii.String("cidr"),
//   	ServerCertificateArn: jsii.String("serverCertificateArn"),
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	AuthorizeAllUsersToVpcCidr: jsii.Boolean(false),
//   	ClientCertificateArn: jsii.String("clientCertificateArn"),
//   	ClientConnectionHandler: clientVpnConnectionHandler,
//   	ClientLoginBanner: jsii.String("clientLoginBanner"),
//   	Description: jsii.String("description"),
//   	DnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	Logging: jsii.Boolean(false),
//   	LogGroup: logGroup,
//   	LogStream: logStream,
//   	Port: awscdk.Aws_ec2.VpnPort_HTTPS,
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	SelfServicePortal: jsii.Boolean(false),
//   	SessionTimeout: awscdk.*Aws_ec2.ClientVpnSessionTimeout_EIGHT_HOURS,
//   	SplitTunnel: jsii.Boolean(false),
//   	TransportProtocol: awscdk.*Aws_ec2.TransportProtocol_TCP,
//   	UserBasedAuthentication: clientVpnUserBasedAuthentication,
//   	VpcSubnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.*Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type ClientVpnEndpointProps struct {
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses.
	//
	// The address range cannot overlap with the local CIDR of the VPC
	// in which the associated subnet is located, or the routes that you add manually.
	//
	// Changing the address range will replace the Client VPN endpoint.
	//
	// The CIDR block should be /22 or greater.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The ARN of the server certificate.
	ServerCertificateArn *string `field:"required" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// Whether to authorize all users to the VPC CIDR.
	//
	// This automatically creates an authorization rule. Set this to `false` and
	// use `addAuthorizationRule()` to create your own rules instead.
	// Default: true.
	//
	AuthorizeAllUsersToVpcCidr *bool `field:"optional" json:"authorizeAllUsersToVpcCidr" yaml:"authorizeAllUsersToVpcCidr"`
	// The ARN of the client certificate for mutual authentication.
	//
	// The certificate must be signed by a certificate authority (CA) and it must
	// be provisioned in AWS Certificate Manager (ACM).
	// Default: - use user-based authentication.
	//
	ClientCertificateArn *string `field:"optional" json:"clientCertificateArn" yaml:"clientCertificateArn"`
	// The AWS Lambda function used for connection authorization.
	//
	// The name of the Lambda function must begin with the `AWSClientVPN-` prefix.
	// Default: - no connection handler.
	//
	ClientConnectionHandler IClientVpnConnectionHandler `field:"optional" json:"clientConnectionHandler" yaml:"clientConnectionHandler"`
	// Customizable text that will be displayed in a banner on AWS provided clients when a VPN session is established.
	//
	// UTF-8 encoded characters only. Maximum of 1400 characters.
	// Default: - no banner is presented to the client.
	//
	ClientLoginBanner *string `field:"optional" json:"clientLoginBanner" yaml:"clientLoginBanner"`
	// A brief description of the Client VPN endpoint.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about the DNS servers to be used for DNS resolution.
	//
	// A Client VPN endpoint can have up to two DNS servers.
	// Default: - use the DNS address configured on the device.
	//
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Whether to enable connections logging.
	// Default: true.
	//
	Logging *bool `field:"optional" json:"logging" yaml:"logging"`
	// A CloudWatch Logs log group for connection logging.
	// Default: - a new group is created.
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// A CloudWatch Logs log stream for connection logging.
	// Default: - a new stream is created.
	//
	LogStream awslogs.ILogStream `field:"optional" json:"logStream" yaml:"logStream"`
	// The port number to assign to the Client VPN endpoint for TCP and UDP traffic.
	// Default: VpnPort.HTTPS
	//
	Port VpnPort `field:"optional" json:"port" yaml:"port"`
	// The security groups to apply to the target network.
	// Default: - a new security group is created.
	//
	SecurityGroups *[]ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specify whether to enable the self-service portal for the Client VPN endpoint.
	// Default: true.
	//
	SelfServicePortal *bool `field:"optional" json:"selfServicePortal" yaml:"selfServicePortal"`
	// The maximum VPN session duration time.
	// Default: ClientVpnSessionTimeout.TWENTY_FOUR_HOURS
	//
	SessionTimeout ClientVpnSessionTimeout `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// Indicates whether split-tunnel is enabled on the AWS Client VPN endpoint.
	// See: https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/split-tunnel-vpn.html
	//
	// Default: false.
	//
	SplitTunnel *bool `field:"optional" json:"splitTunnel" yaml:"splitTunnel"`
	// The transport protocol to be used by the VPN session.
	// Default: TransportProtocol.UDP
	//
	TransportProtocol TransportProtocol `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
	// The type of user-based authentication to use.
	// See: https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/client-authentication.html
	//
	// Default: - use mutual authentication.
	//
	UserBasedAuthentication ClientVpnUserBasedAuthentication `field:"optional" json:"userBasedAuthentication" yaml:"userBasedAuthentication"`
	// Subnets to associate to the client VPN endpoint.
	// Default: - the VPC default strategy.
	//
	VpcSubnets *SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The VPC to connect to.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

