package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnServer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerProps := &CfnServerProps{
//   	Certificate: jsii.String("certificate"),
//   	Domain: jsii.String("domain"),
//   	EndpointDetails: &EndpointDetailsProperty{
//   		AddressAllocationIds: []*string{
//   			jsii.String("addressAllocationIds"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcEndpointId: jsii.String("vpcEndpointId"),
//   		VpcId: jsii.String("vpcId"),
//   	},
//   	EndpointType: jsii.String("endpointType"),
//   	IdentityProviderDetails: &IdentityProviderDetailsProperty{
//   		DirectoryId: jsii.String("directoryId"),
//   		Function: jsii.String("function"),
//   		InvocationRole: jsii.String("invocationRole"),
//   		SftpAuthenticationMethods: jsii.String("sftpAuthenticationMethods"),
//   		Url: jsii.String("url"),
//   	},
//   	IdentityProviderType: jsii.String("identityProviderType"),
//   	LoggingRole: jsii.String("loggingRole"),
//   	PostAuthenticationLoginBanner: jsii.String("postAuthenticationLoginBanner"),
//   	PreAuthenticationLoginBanner: jsii.String("preAuthenticationLoginBanner"),
//   	ProtocolDetails: &ProtocolDetailsProperty{
//   		As2Transports: []*string{
//   			jsii.String("as2Transports"),
//   		},
//   		PassiveIp: jsii.String("passiveIp"),
//   		SetStatOption: jsii.String("setStatOption"),
//   		TlsSessionResumptionMode: jsii.String("tlsSessionResumptionMode"),
//   	},
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	S3StorageOptions: &S3StorageOptionsProperty{
//   		DirectoryListingOptimization: jsii.String("directoryListingOptimization"),
//   	},
//   	SecurityPolicyName: jsii.String("securityPolicyName"),
//   	StructuredLogDestinations: []*string{
//   		jsii.String("structuredLogDestinations"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkflowDetails: &WorkflowDetailsProperty{
//   		OnPartialUpload: []interface{}{
//   			&WorkflowDetailProperty{
//   				ExecutionRole: jsii.String("executionRole"),
//   				WorkflowId: jsii.String("workflowId"),
//   			},
//   		},
//   		OnUpload: []interface{}{
//   			&WorkflowDetailProperty{
//   				ExecutionRole: jsii.String("executionRole"),
//   				WorkflowId: jsii.String("workflowId"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-certificate
	//
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Specifies the domain of the storage system that is used for file transfers.
	//
	// There are two domains available: Amazon Simple Storage Service (Amazon S3) and Amazon Elastic File System (Amazon EFS). The default value is S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The virtual private cloud (VPC) endpoint settings that are configured for your server.
	//
	// When you host your endpoint within your VPC, you can make your endpoint accessible only to resources within your VPC, or you can attach Elastic IP addresses and make your endpoint accessible to clients over the internet. Your VPC's default security groups are automatically assigned to your endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-endpointdetails
	//
	EndpointDetails interface{} `field:"optional" json:"endpointDetails" yaml:"endpointDetails"`
	// The type of endpoint that you want your server to use.
	//
	// You can choose to make your server's endpoint publicly accessible (PUBLIC) or host it inside your VPC. With an endpoint that is hosted in a VPC, you can restrict access to your server and resources only within your VPC or choose to make it internet facing by attaching Elastic IP addresses directly to it.
	//
	// > After May 19, 2021, you won't be able to create a server using `EndpointType=VPC_ENDPOINT` in your AWS account if your account hasn't already done so before May 19, 2021. If you have already created servers with `EndpointType=VPC_ENDPOINT` in your AWS account on or before May 19, 2021, you will not be affected. After this date, use `EndpointType` = `VPC` .
	// >
	// > For more information, see [Discontinuing the use of VPC_ENDPOINT](https://docs.aws.amazon.com//transfer/latest/userguide/create-server-in-vpc.html#deprecate-vpc-endpoint) .
	// >
	// > It is recommended that you use `VPC` as the `EndpointType` . With this endpoint type, you have the option to directly associate up to three Elastic IPv4 addresses (BYO IP included) with your server's endpoint and use VPC security groups to restrict traffic by the client's public IP address. This is not possible with `EndpointType` set to `VPC_ENDPOINT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-endpointtype
	//
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// Required when `IdentityProviderType` is set to `AWS_DIRECTORY_SERVICE` , `AWS _LAMBDA` or `API_GATEWAY` .
	//
	// Accepts an array containing all of the information required to use a directory in `AWS_DIRECTORY_SERVICE` or invoke a customer-supplied authentication API, including the API Gateway URL. Not required when `IdentityProviderType` is set to `SERVICE_MANAGED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-identityproviderdetails
	//
	IdentityProviderDetails interface{} `field:"optional" json:"identityProviderDetails" yaml:"identityProviderDetails"`
	// The mode of authentication for a server.
	//
	// The default value is `SERVICE_MANAGED` , which allows you to store and access user credentials within the AWS Transfer Family service.
	//
	// Use `AWS_DIRECTORY_SERVICE` to provide access to Active Directory groups in AWS Directory Service for Microsoft Active Directory or Microsoft Active Directory in your on-premises environment or in AWS using AD Connector. This option also requires you to provide a Directory ID by using the `IdentityProviderDetails` parameter.
	//
	// Use the `API_GATEWAY` value to integrate with an identity provider of your choosing. The `API_GATEWAY` setting requires you to provide an Amazon API Gateway endpoint URL to call for authentication by using the `IdentityProviderDetails` parameter.
	//
	// Use the `AWS_LAMBDA` value to directly use an AWS Lambda function as your identity provider. If you choose this value, you must specify the ARN for the Lambda function in the `Function` parameter for the `IdentityProviderDetails` data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-identityprovidertype
	//
	IdentityProviderType *string `field:"optional" json:"identityProviderType" yaml:"identityProviderType"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFSevents.
	//
	// When set, you can view user activity in your CloudWatch logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-loggingrole
	//
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// Specifies a string to display when users connect to a server. This string is displayed after the user authenticates.
	//
	// > The SFTP protocol does not support post-authentication display banners.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-postauthenticationloginbanner
	//
	PostAuthenticationLoginBanner *string `field:"optional" json:"postAuthenticationLoginBanner" yaml:"postAuthenticationLoginBanner"`
	// Specifies a string to display when users connect to a server.
	//
	// This string is displayed before the user authenticates. For example, the following banner displays details about using the system:
	//
	// `This system is for the use of authorized users only. Individuals using this computer system without authority, or in excess of their authority, are subject to having all of their activities on this system monitored and recorded by system personnel.`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-preauthenticationloginbanner
	//
	PreAuthenticationLoginBanner *string `field:"optional" json:"preAuthenticationLoginBanner" yaml:"preAuthenticationLoginBanner"`
	// The protocol settings that are configured for your server.
	//
	// - To indicate passive mode (for FTP and FTPS protocols), use the `PassiveIp` parameter. Enter a single dotted-quad IPv4 address, such as the external IP address of a firewall, router, or load balancer.
	// - To ignore the error that is generated when the client attempts to use the `SETSTAT` command on a file that you are uploading to an Amazon S3 bucket, use the `SetStatOption` parameter. To have the AWS Transfer Family server ignore the `SETSTAT` command and upload files without needing to make any changes to your SFTP client, set the value to `ENABLE_NO_OP` . If you set the `SetStatOption` parameter to `ENABLE_NO_OP` , Transfer Family generates a log entry to Amazon CloudWatch Logs, so that you can determine when the client is making a `SETSTAT` call.
	// - To determine whether your AWS Transfer Family server resumes recent, negotiated sessions through a unique session ID, use the `TlsSessionResumptionMode` parameter.
	// - `As2Transports` indicates the transport method for the AS2 messages. Currently, only HTTP is supported.
	//
	// The `Protocols` parameter is an array of strings.
	//
	// *Allowed values* : One or more of `SFTP` , `FTPS` , `FTP` , `AS2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-protocoldetails
	//
	ProtocolDetails interface{} `field:"optional" json:"protocolDetails" yaml:"protocolDetails"`
	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint.
	//
	// The available protocols are:
	//
	// - `SFTP` (Secure Shell (SSH) File Transfer Protocol): File transfer over SSH
	// - `FTPS` (File Transfer Protocol Secure): File transfer with TLS encryption
	// - `FTP` (File Transfer Protocol): Unencrypted file transfer
	// - `AS2` (Applicability Statement 2): used for transporting structured business-to-business data
	//
	// > - If you select `FTPS` , you must choose a certificate stored in AWS Certificate Manager (ACM) which is used to identify your server when clients connect to it over FTPS.
	// > - If `Protocol` includes either `FTP` or `FTPS` , then the `EndpointType` must be `VPC` and the `IdentityProviderType` must be either `AWS_DIRECTORY_SERVICE` , `AWS_LAMBDA` , or `API_GATEWAY` .
	// > - If `Protocol` includes `FTP` , then `AddressAllocationIds` cannot be associated.
	// > - If `Protocol` is set only to `SFTP` , the `EndpointType` can be set to `PUBLIC` and the `IdentityProviderType` can be set any of the supported identity types: `SERVICE_MANAGED` , `AWS_DIRECTORY_SERVICE` , `AWS_LAMBDA` , or `API_GATEWAY` .
	// > - If `Protocol` includes `AS2` , then the `EndpointType` must be `VPC` , and domain must be Amazon S3.
	//
	// The `Protocols` parameter is an array of strings.
	//
	// *Allowed values* : One or more of `SFTP` , `FTPS` , `FTP` , `AS2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-protocols
	//
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// Specifies whether or not performance for your Amazon S3 directories is optimized. This is disabled by default.
	//
	// By default, home directory mappings have a `TYPE` of `DIRECTORY` . If you enable this option, you would then need to explicitly set the `HomeDirectoryMapEntry` `Type` to `FILE` if you want a mapping to have a file target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-s3storageoptions
	//
	S3StorageOptions interface{} `field:"optional" json:"s3StorageOptions" yaml:"s3StorageOptions"`
	// Specifies the name of the security policy for the server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-securitypolicyname
	//
	SecurityPolicyName *string `field:"optional" json:"securityPolicyName" yaml:"securityPolicyName"`
	// Specifies the log groups to which your server logs are sent.
	//
	// To specify a log group, you must provide the ARN for an existing log group. In this case, the format of the log group is as follows:
	//
	// `arn:aws:logs:region-name:amazon-account-id:log-group:log-group-name:*`
	//
	// For example, `arn:aws:logs:us-east-1:111122223333:log-group:mytestgroup:*`
	//
	// If you have previously specified a log group for a server, you can clear it, and in effect turn off structured logging, by providing an empty value for this parameter in an `update-server` call. For example:
	//
	// `update-server --server-id s-1234567890abcdef0 --structured-log-destinations`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-structuredlogdestinations
	//
	StructuredLogDestinations *[]*string `field:"optional" json:"structuredLogDestinations" yaml:"structuredLogDestinations"`
	// Key-value pairs that can be used to group and search for servers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the workflow ID for the workflow to assign and the execution role that's used for executing the workflow.
	//
	// In addition to a workflow to execute when a file is uploaded completely, `WorkflowDetails` can also contain a workflow ID (and execution role) for a workflow to execute on partial upload. A partial upload occurs when a file is open when the session disconnects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-workflowdetails
	//
	WorkflowDetails interface{} `field:"optional" json:"workflowDetails" yaml:"workflowDetails"`
}

