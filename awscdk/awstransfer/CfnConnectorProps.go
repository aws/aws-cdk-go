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
//   	As2Config: as2Config,
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	LoggingRole: jsii.String("loggingRole"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConnectorProps struct {
	// With AS2, you can send files by calling `StartFileTransfer` and specifying the file paths in the request parameter, `SendFilePaths` .
	//
	// We use the file’s parent directory (for example, for `--send-file-paths /bucket/dir/file.txt` , parent directory is `/bucket/dir/` ) to temporarily store a processed AS2 message file, store the MDN when we receive them from the partner, and write a final JSON file containing relevant metadata of the transmission. So, the `AccessRole` needs to provide read and write access to the parent directory of the file location used in the `StartFileTransfer` request. Additionally, you need to provide read and write access to the parent directory of the files that you intend to send with `StartFileTransfer` .
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// A structure that contains the parameters for a connector object.
	As2Config interface{} `field:"required" json:"as2Config" yaml:"as2Config"`
	// The URL of the partner's AS2 endpoint.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that allows a connector to turn on CloudWatch logging for Amazon S3 events.
	//
	// When set, you can view connector activity in your CloudWatch logs.
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// Key-value pairs that can be used to group and search for connectors.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

