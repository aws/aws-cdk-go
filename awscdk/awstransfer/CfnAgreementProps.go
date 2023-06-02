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
//   	BaseDirectory: jsii.String("baseDirectory"),
//   	LocalProfileId: jsii.String("localProfileId"),
//   	PartnerProfileId: jsii.String("partnerProfileId"),
//   	ServerId: jsii.String("serverId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Status: jsii.String("status"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAgreementProps struct {
	// With AS2, you can send files by calling `StartFileTransfer` and specifying the file paths in the request parameter, `SendFilePaths` .
	//
	// We use the file’s parent directory (for example, for `--send-file-paths /bucket/dir/file.txt` , parent directory is `/bucket/dir/` ) to temporarily store a processed AS2 message file, store the MDN when we receive them from the partner, and write a final JSON file containing relevant metadata of the transmission. So, the `AccessRole` needs to provide read and write access to the parent directory of the file location used in the `StartFileTransfer` request. Additionally, you need to provide read and write access to the parent directory of the files that you intend to send with `StartFileTransfer` .
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// The landing directory (folder) for files that are transferred by using the AS2 protocol.
	BaseDirectory *string `field:"required" json:"baseDirectory" yaml:"baseDirectory"`
	// A unique identifier for the AS2 local profile.
	LocalProfileId *string `field:"required" json:"localProfileId" yaml:"localProfileId"`
	// A unique identifier for the partner profile used in the agreement.
	PartnerProfileId *string `field:"required" json:"partnerProfileId" yaml:"partnerProfileId"`
	// A system-assigned unique identifier for a server instance.
	//
	// This identifier indicates the specific server that the agreement uses.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// The name or short description that's used to identify the agreement.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The current status of the agreement, either `ACTIVE` or `INACTIVE` .
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Key-value pairs that can be used to group and search for agreements.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

