package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var as2Config interface{}
//
//   cfnConnectorProps := &CfnConnectorProps{
//   	AccessRole: jsii.String("accessRole"),
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	As2Config: as2Config,
//   	LoggingRole: jsii.String("loggingRole"),
//   	SecurityPolicyName: jsii.String("securityPolicyName"),
//   	SftpConfig: &SftpConfigProperty{
//   		MaxConcurrentConnections: jsii.Number(123),
//   		TrustedHostKeys: []*string{
//   			jsii.String("trustedHostKeys"),
//   		},
//   		UserSecretId: jsii.String("userSecretId"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html
//
type CfnConnectorProps struct {
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
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// The URL of the partner's AS2 or SFTP endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// A structure that contains the parameters for an AS2 connector object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-connector.html#cfn-transfer-connector-as2config
	//
	As2Config interface{} `field:"optional" json:"as2Config" yaml:"as2Config"`
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
}

