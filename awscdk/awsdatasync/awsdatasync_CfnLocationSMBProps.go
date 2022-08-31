package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocationSMB`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationSMBProps := &cfnLocationSMBProps{
//   	agentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	password: jsii.String("password"),
//   	serverHostname: jsii.String("serverHostname"),
//   	subdirectory: jsii.String("subdirectory"),
//   	user: jsii.String("user"),
//
//   	// the properties below are optional
//   	domain: jsii.String("domain"),
//   	mountOptions: &mountOptionsProperty{
//   		version: jsii.String("version"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationSMBProps struct {
	// The Amazon Resource Names (ARNs) of agents to use for a Server Message Block (SMB) location.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The name of the SMB server.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the SMB server. An agent that is installed on-premises uses this hostname to mount the SMB server in a network.
	//
	// > This name must either be DNS-compliant or must be an IP version 4 (IPv4) address.
	ServerHostname *string `field:"required" json:"serverHostname" yaml:"serverHostname"`
	// The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination.
	//
	// The SMB path should be a path that's exported by the SMB server, or a subdirectory of that path. The path should be such that it can be mounted by other SMB clients in your network.
	//
	// > `Subdirectory` must be specified with forward slashes. For example, `/path/to/folder` .
	//
	// To transfer all the data in the folder you specified, DataSync must have permissions to mount the SMB share, as well as to access all the data in that share. To ensure this, either make sure that the user name and password specified belongs to the user who can mount the share, and who has the appropriate permissions for all of the files and directories that you want DataSync to access, or use credentials of a member of the Backup Operators group to mount the share. Doing either one enables the agent to access the data. For the agent to access directories, you must additionally enable all execute access.
	Subdirectory *string `field:"required" json:"subdirectory" yaml:"subdirectory"`
	// The user who can mount the share and has the permissions to access files and folders in the SMB share.
	//
	// For information about choosing a user name that ensures sufficient permissions to files, folders, and metadata, see [user](https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#SMBuser) .
	User *string `field:"required" json:"user" yaml:"user"`
	// The name of the Windows domain that the SMB server belongs to.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The mount options used by DataSync to access the SMB server.
	MountOptions interface{} `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The key-value pair that represents the tag that you want to add to the location.
	//
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

