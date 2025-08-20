package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAgreement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgreementProps := &CfnAgreementProps{
//   	AccessRole: jsii.String("accessRole"),
//   	LocalProfileId: jsii.String("localProfileId"),
//   	PartnerProfileId: jsii.String("partnerProfileId"),
//   	ServerId: jsii.String("serverId"),
//
//   	// the properties below are optional
//   	BaseDirectory: jsii.String("baseDirectory"),
//   	CustomDirectories: &CustomDirectoriesProperty{
//   		FailedFilesDirectory: jsii.String("failedFilesDirectory"),
//   		MdnFilesDirectory: jsii.String("mdnFilesDirectory"),
//   		PayloadFilesDirectory: jsii.String("payloadFilesDirectory"),
//   		StatusFilesDirectory: jsii.String("statusFilesDirectory"),
//   		TemporaryFilesDirectory: jsii.String("temporaryFilesDirectory"),
//   	},
//   	Description: jsii.String("description"),
//   	EnforceMessageSigning: jsii.String("enforceMessageSigning"),
//   	PreserveFilename: jsii.String("preserveFilename"),
//   	Status: jsii.String("status"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html
//
type CfnAgreementProps struct {
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-accessrole
	//
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// A unique identifier for the AS2 local profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-localprofileid
	//
	LocalProfileId *string `field:"required" json:"localProfileId" yaml:"localProfileId"`
	// A unique identifier for the partner profile used in the agreement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-partnerprofileid
	//
	PartnerProfileId *string `field:"required" json:"partnerProfileId" yaml:"partnerProfileId"`
	// A system-assigned unique identifier for a server instance.
	//
	// This identifier indicates the specific server that the agreement uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-serverid
	//
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// The landing directory (folder) for files that are transferred by using the AS2 protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-basedirectory
	//
	BaseDirectory *string `field:"optional" json:"baseDirectory" yaml:"baseDirectory"`
	// A `CustomDirectoriesType` structure.
	//
	// This structure specifies custom directories for storing various AS2 message files. You can specify directories for the following types of files.
	//
	// - Failed files
	// - MDN files
	// - Payload files
	// - Status files
	// - Temporary files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-customdirectories
	//
	CustomDirectories interface{} `field:"optional" json:"customDirectories" yaml:"customDirectories"`
	// The name or short description that's used to identify the agreement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines whether or not unsigned messages from your trading partners will be accepted.
	//
	// - `ENABLED` : Transfer Family rejects unsigned messages from your trading partner.
	// - `DISABLED` (default value): Transfer Family accepts unsigned messages from your trading partner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-enforcemessagesigning
	//
	EnforceMessageSigning *string `field:"optional" json:"enforceMessageSigning" yaml:"enforceMessageSigning"`
	// Determines whether or not Transfer Family appends a unique string of characters to the end of the AS2 message payload filename when saving it.
	//
	// - `ENABLED` : the filename provided by your trading parter is preserved when the file is saved.
	// - `DISABLED` (default value): when Transfer Family saves the file, the filename is adjusted, as described in [File names and locations](https://docs.aws.amazon.com/transfer/latest/userguide/send-as2-messages.html#file-names-as2) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-preservefilename
	//
	PreserveFilename *string `field:"optional" json:"preserveFilename" yaml:"preserveFilename"`
	// The current status of the agreement, either `ACTIVE` or `INACTIVE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Key-value pairs that can be used to group and search for agreements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html#cfn-transfer-agreement-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

