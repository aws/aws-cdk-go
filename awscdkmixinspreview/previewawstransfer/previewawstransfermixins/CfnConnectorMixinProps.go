package previewawstransfermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var as2Config interface{}
//
//   cfnConnectorMixinProps := &CfnConnectorMixinProps{
//   	AccessRole: jsii.String("accessRole"),
//   	As2Config: as2Config,
//   	EgressConfig: &ConnectorEgressConfigProperty{
//   		VpcLattice: &ConnectorVpcLatticeEgressConfigProperty{
//   			PortNumber: jsii.Number(123),
//   			ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//   		},
//   	},
//   	EgressType: jsii.String("egressType"),
//   	LoggingRole: jsii.String("loggingRole"),
//   	SecurityPolicyName: jsii.String("securityPolicyName"),
//   	SftpConfig: &SftpConfigProperty{
//   		MaxConcurrentConnections: jsii.Number(123),
//   		TrustedHostKeys: []*string{
//   			jsii.String("trustedHostKeys"),
//   		},
//   		UserSecretId: jsii.String("userSecretId"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html
//
type CfnConnectorMixinProps struct {
	// Connectors are used to send files using either the AS2 or SFTP protocol.
	//
	// For the access role, provide the Amazon Resource Name (ARN) of the AWS Identity and Access Management role to use.
	//
	// *For AS2 connectors*
	//
	// With AS2, you can send files by calling `StartFileTransfer` and specifying the file paths in the request parameter, `SendFilePaths` . We use the file’s parent directory (for example, for `--send-file-paths /bucket/dir/file.txt` , parent directory is `/bucket/dir/` ) to temporarily store a processed AS2 message file, store the MDN when we receive them from the partner, and write a final JSON file containing relevant metadata of the transmission. So, the `AccessRole` needs to provide read and write access to the parent directory of the file location used in the `StartFileTransfer` request. Additionally, you need to provide read and write access to the parent directory of the files that you intend to send with `StartFileTransfer` .
	//
	// If you are using Basic authentication for your AS2 connector, the access role requires the `secretsmanager:GetSecretValue` permission for the secret. If the secret is encrypted using a customer-managed key instead of the AWS managed key in Secrets Manager, then the role also needs the `kms:Decrypt` permission for that key.
	//
	// *For SFTP connectors*
	//
	// Make sure that the access role provides read and write access to the parent directory of the file location that's used in the `StartFileTransfer` request. Additionally, make sure that the role provides `secretsmanager:GetSecretValue` permission to AWS Secrets Manager .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-accessrole
	//
	AccessRole *string `field:"optional" json:"accessRole" yaml:"accessRole"`
	// A structure that contains the parameters for an AS2 connector object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-as2config
	//
	As2Config interface{} `field:"optional" json:"as2Config" yaml:"as2Config"`
	// Current egress configuration of the connector, showing how traffic is routed to the SFTP server.
	//
	// Contains VPC Lattice settings when using VPC_LATTICE egress type.
	//
	// When using the VPC_LATTICE egress type, AWS Transfer Family uses a managed Service Network to simplify the resource sharing process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-egressconfig
	//
	EgressConfig interface{} `field:"optional" json:"egressConfig" yaml:"egressConfig"`
	// Type of egress configuration for the connector.
	//
	// SERVICE_MANAGED uses Transfer Family managed NAT gateways, while VPC_LATTICE routes traffic through customer VPCs using VPC Lattice.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-egresstype
	//
	EgressType *string `field:"optional" json:"egressType" yaml:"egressType"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that allows a connector to turn on CloudWatch logging for Amazon S3 events.
	//
	// When set, you can view connector activity in your CloudWatch logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-loggingrole
	//
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// The text name of the security policy for the specified connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-securitypolicyname
	//
	SecurityPolicyName *string `field:"optional" json:"securityPolicyName" yaml:"securityPolicyName"`
	// A structure that contains the parameters for an SFTP connector object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-sftpconfig
	//
	SftpConfig interface{} `field:"optional" json:"sftpConfig" yaml:"sftpConfig"`
	// Key-value pairs that can be used to group and search for connectors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The URL of the partner's AS2 or SFTP endpoint.
	//
	// When creating AS2 connectors or service-managed SFTP connectors (connectors without egress configuration), you must provide a URL to specify the remote server endpoint. For VPC Lattice type connectors, the URL must be null.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

