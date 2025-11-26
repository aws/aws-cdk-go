package previewawstransfermixins


// The protocol settings that are configured for your server.
//
// > Avoid placing Network Load Balancers (NLBs) or NAT gateways in front of AWS Transfer Family servers, as this increases costs and can cause performance issues, including reduced connection limits for FTPS. For more details, see [Avoid placing NLBs and NATs in front of AWS Transfer Family](https://docs.aws.amazon.com/transfer/latest/userguide/infrastructure-security.html#nlb-considerations) .
//
// - To indicate passive mode (for FTP and FTPS protocols), use the `PassiveIp` parameter. Enter a single dotted-quad IPv4 address, such as the external IP address of a firewall, router, or load balancer.
// - To ignore the error that is generated when the client attempts to use the `SETSTAT` command on a file that you are uploading to an Amazon S3 bucket, use the `SetStatOption` parameter. To have the AWS Transfer Family server ignore the `SETSTAT` command and upload files without needing to make any changes to your SFTP client, set the value to `ENABLE_NO_OP` . If you set the `SetStatOption` parameter to `ENABLE_NO_OP` , Transfer Family generates a log entry to Amazon CloudWatch Logs, so that you can determine when the client is making a `SETSTAT` call.
// - To determine whether your AWS Transfer Family server resumes recent, negotiated sessions through a unique session ID, use the `TlsSessionResumptionMode` parameter.
// - `As2Transports` indicates the transport method for the AS2 messages. Currently, only HTTP is supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   protocolDetailsProperty := &ProtocolDetailsProperty{
//   	As2Transports: []*string{
//   		jsii.String("as2Transports"),
//   	},
//   	PassiveIp: jsii.String("passiveIp"),
//   	SetStatOption: jsii.String("setStatOption"),
//   	TlsSessionResumptionMode: jsii.String("tlsSessionResumptionMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html
//
type CfnServerPropsMixin_ProtocolDetailsProperty struct {
	// List of `As2Transport` objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html#cfn-transfer-server-protocoldetails-as2transports
	//
	As2Transports *[]*string `field:"optional" json:"as2Transports" yaml:"as2Transports"`
	// Indicates passive mode, for FTP and FTPS protocols.
	//
	// Enter a single IPv4 address, such as the public IP address of a firewall, router, or load balancer. For example:
	//
	// `aws transfer update-server --protocol-details PassiveIp=0.0.0.0`
	//
	// Replace `0.0.0.0` in the example above with the actual IP address you want to use.
	//
	// > If you change the `PassiveIp` value, you must stop and then restart your Transfer Family server for the change to take effect. For details on using passive mode (PASV) in a NAT environment, see [Configuring your FTPS server behind a firewall or NAT with AWS Transfer Family](https://docs.aws.amazon.com/storage/configuring-your-ftps-server-behind-a-firewall-or-nat-with-aws-transfer-family/) .
	// >
	// > Additionally, avoid placing Network Load Balancers (NLBs) or NAT gateways in front of AWS Transfer Family servers. This configuration increases costs and can cause performance issues. When NLBs or NATs are in the communication path, Transfer Family cannot accurately recognize client IP addresses, which impacts connection sharding and limits FTPS servers to only 300 simultaneous connections instead of 10,000. If you must use an NLB, use port 21 for health checks and enable TLS session resumption by setting `TlsSessionResumptionMode = ENFORCED` . For optimal performance, migrate to VPC endpoints with Elastic IP addresses instead of using NLBs. For more details, see [Avoid placing NLBs and NATs in front of AWS Transfer Family](https://docs.aws.amazon.com/transfer/latest/userguide/infrastructure-security.html#nlb-considerations) .
	//
	// *Special values*
	//
	// The `AUTO` and `0.0.0.0` are special values for the `PassiveIp` parameter. The value `PassiveIp=AUTO` is assigned by default to FTP and FTPS type servers. In this case, the server automatically responds with one of the endpoint IPs within the PASV response. `PassiveIp=0.0.0.0` has a more unique application for its usage. For example, if you have a High Availability (HA) Network Load Balancer (NLB) environment, where you have 3 subnets, you can only specify a single IP address using the `PassiveIp` parameter. This reduces the effectiveness of having High Availability. In this case, you can specify `PassiveIp=0.0.0.0` . This tells the client to use the same IP address as the Control connection and utilize all AZs for their connections. Note, however, that not all FTP clients support the `PassiveIp=0.0.0.0` response. FileZilla and WinSCP do support it. If you are using other clients, check to see if your client supports the `PassiveIp=0.0.0.0` response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html#cfn-transfer-server-protocoldetails-passiveip
	//
	PassiveIp *string `field:"optional" json:"passiveIp" yaml:"passiveIp"`
	// Use the `SetStatOption` to ignore the error that is generated when the client attempts to use `SETSTAT` on a file you are uploading to an S3 bucket.
	//
	// Some SFTP file transfer clients can attempt to change the attributes of remote files, including timestamp and permissions, using commands, such as `SETSTAT` when uploading the file. However, these commands are not compatible with object storage systems, such as Amazon S3. Due to this incompatibility, file uploads from these clients can result in errors even when the file is otherwise successfully uploaded.
	//
	// Set the value to `ENABLE_NO_OP` to have the Transfer Family server ignore the `SETSTAT` command, and upload files without needing to make any changes to your SFTP client. While the `SetStatOption` `ENABLE_NO_OP` setting ignores the error, it does generate a log entry in Amazon CloudWatch Logs, so you can determine when the client is making a `SETSTAT` call.
	//
	// > If you want to preserve the original timestamp for your file, and modify other file attributes using `SETSTAT` , you can use Amazon EFS as backend storage with Transfer Family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html#cfn-transfer-server-protocoldetails-setstatoption
	//
	SetStatOption *string `field:"optional" json:"setStatOption" yaml:"setStatOption"`
	// A property used with Transfer Family servers that use the FTPS protocol.
	//
	// TLS Session Resumption provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. `TlsSessionResumptionMode` determines whether or not the server resumes recent, negotiated sessions through a unique session ID. This property is available during `CreateServer` and `UpdateServer` calls. If a `TlsSessionResumptionMode` value is not specified during `CreateServer` , it is set to `ENFORCED` by default.
	//
	// - `DISABLED` : the server does not process TLS session resumption client requests and creates a new TLS session for each request.
	// - `ENABLED` : the server processes and accepts clients that are performing TLS session resumption. The server doesn't reject client data connections that do not perform the TLS session resumption client processing.
	// - `ENFORCED` : the server processes and accepts clients that are performing TLS session resumption. The server rejects client data connections that do not perform the TLS session resumption client processing. Before you set the value to `ENFORCED` , test your clients.
	//
	// > Not all FTPS clients perform TLS session resumption. So, if you choose to enforce TLS session resumption, you prevent any connections from FTPS clients that don't perform the protocol negotiation. To determine whether or not you can use the `ENFORCED` value, you need to test your clients.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html#cfn-transfer-server-protocoldetails-tlssessionresumptionmode
	//
	TlsSessionResumptionMode *string `field:"optional" json:"tlsSessionResumptionMode" yaml:"tlsSessionResumptionMode"`
}

