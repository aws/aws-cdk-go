package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocationFSxONTAP`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationFSxONTAPProps := &cfnLocationFSxONTAPProps{
//   	securityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	storageVirtualMachineArn: jsii.String("storageVirtualMachineArn"),
//
//   	// the properties below are optional
//   	protocol: &protocolProperty{
//   		nfs: &nFSProperty{
//   			mountOptions: &nfsMountOptionsProperty{
//   				version: jsii.String("version"),
//   			},
//   		},
//   		smb: &sMBProperty{
//   			mountOptions: &smbMountOptionsProperty{
//   				version: jsii.String("version"),
//   			},
//   			password: jsii.String("password"),
//   			user: jsii.String("user"),
//
//   			// the properties below are optional
//   			domain: jsii.String("domain"),
//   		},
//   	},
//   	subdirectory: jsii.String("subdirectory"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationFSxONTAPProps struct {
	// Specifies the Amazon Resource Names (ARNs) of the security groups that DataSync can use to access your FSx for ONTAP file system.
	//
	// You must configure the security groups to allow outbound traffic on the following ports (depending on the protocol that you're using):
	//
	// - *Network File System (NFS)* : TCP ports 111, 635, and 2049
	// - *Server Message Block (SMB)* : TCP port 445
	//
	// Your file system's security groups must also allow inbound traffic on the same port.
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// Specifies the ARN of the storage virtual machine (SVM) in your file system where you want to copy data to or from.
	StorageVirtualMachineArn *string `field:"required" json:"storageVirtualMachineArn" yaml:"storageVirtualMachineArn"`
	// Specifies the data transfer protocol that DataSync uses to access your Amazon FSx file system.
	Protocol interface{} `field:"optional" json:"protocol" yaml:"protocol"`
	// Specifies a path to the file share in the SVM where you'll copy your data.
	//
	// You can specify a junction path (also known as a mount point), qtree path (for NFS file shares), or share name (for SMB file shares). For example, your mount path might be `/vol1` , `/vol1/tree1` , or `/share1` .
	//
	// > Don't specify a junction path in the SVM's root volume. For more information, see [Managing FSx for ONTAP storage virtual machines](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) in the *Amazon FSx for NetApp ONTAP User Guide* .
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

