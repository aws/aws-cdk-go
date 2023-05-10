package awsdatasync


// Specifies the data transfer protocol that AWS DataSync uses to access your Amazon FSx file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protocolProperty := &ProtocolProperty{
//   	Nfs: &NFSProperty{
//   		MountOptions: &NfsMountOptionsProperty{
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	Smb: &SMBProperty{
//   		MountOptions: &SmbMountOptionsProperty{
//   			Version: jsii.String("version"),
//   		},
//   		Password: jsii.String("password"),
//   		User: jsii.String("user"),
//
//   		// the properties below are optional
//   		Domain: jsii.String("domain"),
//   	},
//   }
//
type CfnLocationFSxONTAP_ProtocolProperty struct {
	// Specifies the Network File System (NFS) protocol configuration that DataSync uses to access your FSx for ONTAP file system's storage virtual machine (SVM).
	Nfs interface{} `field:"optional" json:"nfs" yaml:"nfs"`
	// Specifies the Server Message Block (SMB) protocol configuration that DataSync uses to access your FSx for ONTAP file system's SVM.
	Smb interface{} `field:"optional" json:"smb" yaml:"smb"`
}

