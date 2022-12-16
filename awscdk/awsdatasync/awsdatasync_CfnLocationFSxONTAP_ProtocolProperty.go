package awsdatasync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protocolProperty := &protocolProperty{
//   	nfs: &nFSProperty{
//   		mountOptions: &nfsMountOptionsProperty{
//   			version: jsii.String("version"),
//   		},
//   	},
//   	smb: &sMBProperty{
//   		mountOptions: &smbMountOptionsProperty{
//   			version: jsii.String("version"),
//   		},
//   		password: jsii.String("password"),
//   		user: jsii.String("user"),
//
//   		// the properties below are optional
//   		domain: jsii.String("domain"),
//   	},
//   }
//
type CfnLocationFSxONTAP_ProtocolProperty struct {
	// `CfnLocationFSxONTAP.ProtocolProperty.NFS`.
	Nfs interface{} `field:"optional" json:"nfs" yaml:"nfs"`
	// `CfnLocationFSxONTAP.ProtocolProperty.SMB`.
	Smb interface{} `field:"optional" json:"smb" yaml:"smb"`
}

