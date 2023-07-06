package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocationNFS`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationNFSProps := &CfnLocationNFSProps{
//   	OnPremConfig: &OnPremConfigProperty{
//   		AgentArns: []*string{
//   			jsii.String("agentArns"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	MountOptions: &MountOptionsProperty{
//   		Version: jsii.String("version"),
//   	},
//   	ServerHostname: jsii.String("serverHostname"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationNFSProps struct {
	// Specifies the Amazon Resource Names (ARNs) of agents that DataSync uses to connect to your NFS file server.
	//
	// If you are copying data to or from your AWS Snowcone device, see [NFS Server on AWS Snowcone](https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone) for more information.
	OnPremConfig interface{} `field:"required" json:"onPremConfig" yaml:"onPremConfig"`
	// Specifies the mount options that DataSync can use to mount your NFS share.
	MountOptions interface{} `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// Specifies the IP address or domain name of your NFS file server.
	//
	// An agent that is installed on-premises uses this hostname to mount the NFS server in a network.
	//
	// If you are copying data to or from your AWS Snowcone device, see [NFS Server on AWS Snowcone](https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone) for more information.
	//
	// > You must specify be an IP version 4 address or Domain Name System (DNS)-compliant name.
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// Specifies the subdirectory in the NFS file server that DataSync transfers to or from.
	//
	// The NFS path should be a path that's exported by the NFS server, or a subdirectory of that path. The path should be such that it can be mounted by other NFS clients in your network.
	//
	// To see all the paths exported by your NFS server, run " `showmount -e nfs-server-name` " from an NFS client that has access to your server. You can specify any directory that appears in the results, and any subdirectory of that directory. Ensure that the NFS export is accessible without Kerberos authentication.
	//
	// To transfer all the data in the folder you specified, DataSync needs to have permissions to read all the data. To ensure this, either configure the NFS export with `no_root_squash,` or ensure that the permissions for all of the files that you want DataSync allow read access for all users. Doing either enables the agent to read the files. For the agent to access directories, you must additionally enable all execute access.
	//
	// If you are copying data to or from your AWS Snowcone device, see [NFS Server on AWS Snowcone](https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone) for more information.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

