package mixinsawsdatasync


// Specifies the data transfer protocol that AWS DataSync uses to access your Amazon FSx file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   protocolProperty := &ProtocolProperty{
//   	Nfs: &NFSProperty{
//   		MountOptions: &NfsMountOptionsProperty{
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	Smb: &SMBProperty{
//   		Domain: jsii.String("domain"),
//   		MountOptions: &SmbMountOptionsProperty{
//   			Version: jsii.String("version"),
//   		},
//   		Password: jsii.String("password"),
//   		User: jsii.String("user"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-protocol.html
//
type CfnLocationFSxONTAPPropsMixin_ProtocolProperty struct {
	// Specifies the Network File System (NFS) protocol configuration that DataSync uses to access your FSx for ONTAP file system's storage virtual machine (SVM).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-protocol.html#cfn-datasync-locationfsxontap-protocol-nfs
	//
	Nfs interface{} `field:"optional" json:"nfs" yaml:"nfs"`
	// Specifies the Server Message Block (SMB) protocol configuration that DataSync uses to access your FSx for ONTAP file system's SVM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-protocol.html#cfn-datasync-locationfsxontap-protocol-smb
	//
	Smb interface{} `field:"optional" json:"smb" yaml:"smb"`
}

