package previewawsdatasyncmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLocationNFSPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocationNFSMixinProps := &CfnLocationNFSMixinProps{
//   	MountOptions: &MountOptionsProperty{
//   		Version: jsii.String("version"),
//   	},
//   	OnPremConfig: &OnPremConfigProperty{
//   		AgentArns: []*string{
//   			jsii.String("agentArns"),
//   		},
//   	},
//   	ServerHostname: jsii.String("serverHostname"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationnfs.html
//
type CfnLocationNFSMixinProps struct {
	// Specifies the options that DataSync can use to mount your NFS file server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationnfs.html#cfn-datasync-locationnfs-mountoptions
	//
	MountOptions interface{} `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect to your NFS file server.
	//
	// You can specify more than one agent. For more information, see [Using multiple DataSync agents](https://docs.aws.amazon.com/datasync/latest/userguide/do-i-need-datasync-agent.html#multiple-agents) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationnfs.html#cfn-datasync-locationnfs-onpremconfig
	//
	OnPremConfig interface{} `field:"optional" json:"onPremConfig" yaml:"onPremConfig"`
	// Specifies the DNS name or IP address (IPv4 or IPv6) of the NFS file server that your DataSync agent connects to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationnfs.html#cfn-datasync-locationnfs-serverhostname
	//
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// Specifies the export path in your NFS file server that you want DataSync to mount.
	//
	// This path (or a subdirectory of the path) is where DataSync transfers data to or from. For information on configuring an export for DataSync, see [Accessing NFS file servers](https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#accessing-nfs) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationnfs.html#cfn-datasync-locationnfs-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationnfs.html#cfn-datasync-locationnfs-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

