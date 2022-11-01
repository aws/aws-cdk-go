package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnServer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerProps := &cfnServerProps{
//   	certificate: jsii.String("certificate"),
//   	domain: jsii.String("domain"),
//   	endpointDetails: &endpointDetailsProperty{
//   		addressAllocationIds: []*string{
//   			jsii.String("addressAllocationIds"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		vpcEndpointId: jsii.String("vpcEndpointId"),
//   		vpcId: jsii.String("vpcId"),
//   	},
//   	endpointType: jsii.String("endpointType"),
//   	identityProviderDetails: &identityProviderDetailsProperty{
//   		directoryId: jsii.String("directoryId"),
//   		function: jsii.String("function"),
//   		invocationRole: jsii.String("invocationRole"),
//   		url: jsii.String("url"),
//   	},
//   	identityProviderType: jsii.String("identityProviderType"),
//   	loggingRole: jsii.String("loggingRole"),
//   	postAuthenticationLoginBanner: jsii.String("postAuthenticationLoginBanner"),
//   	preAuthenticationLoginBanner: jsii.String("preAuthenticationLoginBanner"),
//   	protocolDetails: &protocolDetailsProperty{
//   		as2Transports: []*string{
//   			jsii.String("as2Transports"),
//   		},
//   		passiveIp: jsii.String("passiveIp"),
//   		setStatOption: jsii.String("setStatOption"),
//   		tlsSessionResumptionMode: jsii.String("tlsSessionResumptionMode"),
//   	},
//   	protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	securityPolicyName: jsii.String("securityPolicyName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	workflowDetails: &workflowDetailsProperty{
//   		onPartialUpload: []interface{}{
//   			&workflowDetailProperty{
//   				executionRole: jsii.String("executionRole"),
//   				workflowId: jsii.String("workflowId"),
//   			},
//   		},
//   		onUpload: []interface{}{
//   			&workflowDetailProperty{
//   				executionRole: jsii.String("executionRole"),
//   				workflowId: jsii.String("workflowId"),
//   			},
//   		},
//   	},
//   }
//
type CfnServerProps struct {
	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate.
	//
	// Required when `Protocols` is set to `FTPS` .
	//
	// To request a new public certificate, see [Request a public certificate](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-public.html) in the *AWS Certificate Manager User Guide* .
	//
	// To import an existing certificate into ACM, see [Importing certificates into ACM](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in the *AWS Certificate Manager User Guide* .
	//
	// To request a private certificate to use FTPS through private IP addresses, see [Request a private certificate](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-private.html) in the *AWS Certificate Manager User Guide* .
	//
	// Certificates with the following cryptographic algorithms and key sizes are supported:
	//
	// - 2048-bit RSA (RSA_2048)
	// - 4096-bit RSA (RSA_4096)
	// - Elliptic Prime Curve 256 bit (EC_prime256v1)
	// - Elliptic Prime Curve 384 bit (EC_secp384r1)
	// - Elliptic Prime Curve 521 bit (EC_secp521r1)
	//
	// > The certificate must be a valid SSL/TLS X.509 version 3 certificate with FQDN or IP address specified and information about the issuer.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Specifies the domain of the storage system that is used for file transfers.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The virtual private cloud (VPC) endpoint settings that are configured for your server.
	//
	// When you host your endpoint within your VPC, you can make it accessible only to resources within your VPC, or you can attach Elastic IP addresses and make it accessible to clients over the internet. Your VPC's default security groups are automatically assigned to your endpoint.
	EndpointDetails interface{} `field:"optional" json:"endpointDetails" yaml:"endpointDetails"`
	// The type of endpoint that you want your server to use.
	//
	// You can choose to make your server's endpoint publicly accessible (PUBLIC) or host it inside your VPC. With an endpoint that is hosted in a VPC, you can restrict access to your server and resources only within your VPC or choose to make it internet facing by attaching Elastic IP addresses directly to it.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// Required when `IdentityProviderType` is set to `AWS_DIRECTORY_SERVICE` or `API_GATEWAY` .
	//
	// Accepts an array containing all of the information required to use a directory in `AWS_DIRECTORY_SERVICE` or invoke a customer-supplied authentication API, including the API Gateway URL. Not required when `IdentityProviderType` is set to `SERVICE_MANAGED` .
	IdentityProviderDetails interface{} `field:"optional" json:"identityProviderDetails" yaml:"identityProviderDetails"`
	// Specifies the mode of authentication for a server.
	//
	// The default value is `SERVICE_MANAGED` , which allows you to store and access user credentials within the AWS Transfer Family service.
	//
	// Use `AWS_DIRECTORY_SERVICE` to provide access to Active Directory groups in AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment or in AWS using AD Connectors. This option also requires you to provide a Directory ID using the `IdentityProviderDetails` parameter.
	//
	// Use the `API_GATEWAY` value to integrate with an identity provider of your choosing. The `API_GATEWAY` setting requires you to provide an API Gateway endpoint URL to call for authentication using the `IdentityProviderDetails` parameter.
	//
	// Use the `AWS_LAMBDA` value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the `Function` parameter for the `IdentityProviderDetails` data type.
	IdentityProviderType *string `field:"optional" json:"identityProviderType" yaml:"identityProviderType"`
	// Specifies the Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFS events.
	//
	// When set, user activity can be viewed in your CloudWatch logs.
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// Specify a string to display when users connect to a server. This string is displayed after the user authenticates.
	//
	// > The SFTP protocol does not support post-authentication display banners.
	PostAuthenticationLoginBanner *string `field:"optional" json:"postAuthenticationLoginBanner" yaml:"postAuthenticationLoginBanner"`
	// Specify a string to display when users connect to a server.
	//
	// This string is displayed before the user authenticates. For example, the following banner displays details about using the system.
	//
	// `This system is for the use of authorized users only. Individuals using this computer system without authority, or in excess of their authority, are subject to having all of their activities on this system monitored and recorded by system personnel.`
	PreAuthenticationLoginBanner *string `field:"optional" json:"preAuthenticationLoginBanner" yaml:"preAuthenticationLoginBanner"`
	// The protocol settings that are configured for your server.
	//
	// - Use the `PassiveIp` parameter to indicate passive mode (for FTP and FTPS protocols). Enter a single dotted-quad IPv4 address, such as the external IP address of a firewall, router, or load balancer.
	// - Use the `SetStatOption` to ignore the error that is generated when the client attempts to use SETSTAT on a file you are uploading to an S3 bucket. Set the value to `ENABLE_NO_OP` to have the Transfer Family server ignore the SETSTAT command, and upload files without needing to make any changes to your SFTP client. Note that with `SetStatOption` set to `ENABLE_NO_OP` , Transfer generates a log entry to CloudWatch Logs, so you can determine when the client is making a SETSTAT call.
	// - Use the `TlsSessionResumptionMode` parameter to determine whether or not your Transfer server resumes recent, negotiated sessions through a unique session ID.
	ProtocolDetails interface{} `field:"optional" json:"protocolDetails" yaml:"protocolDetails"`
	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint.
	//
	// The available protocols are:
	//
	// - `SFTP` (Secure Shell (SSH) File Transfer Protocol): File transfer over SSH
	// - `FTPS` (File Transfer Protocol Secure): File transfer with TLS encryption
	// - `FTP` (File Transfer Protocol): Unencrypted file transfer
	//
	// > If you select `FTPS` , you must choose a certificate stored in AWS Certificate Manager (ACM) which is used to identify your server when clients connect to it over FTPS.
	// >
	// > If `Protocol` includes either `FTP` or `FTPS` , then the `EndpointType` must be `VPC` and the `IdentityProviderType` must be `AWS_DIRECTORY_SERVICE` or `API_GATEWAY` .
	// >
	// > If `Protocol` includes `FTP` , then `AddressAllocationIds` cannot be associated.
	// >
	// > If `Protocol` is set only to `SFTP` , the `EndpointType` can be set to `PUBLIC` and the `IdentityProviderType` can be set to `SERVICE_MANAGED` .
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// Specifies the name of the security policy that is attached to the server.
	SecurityPolicyName *string `field:"optional" json:"securityPolicyName" yaml:"securityPolicyName"`
	// Key-value pairs that can be used to group and search for servers.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the workflow ID for the workflow to assign and the execution role used for executing the workflow.
	WorkflowDetails interface{} `field:"optional" json:"workflowDetails" yaml:"workflowDetails"`
}

