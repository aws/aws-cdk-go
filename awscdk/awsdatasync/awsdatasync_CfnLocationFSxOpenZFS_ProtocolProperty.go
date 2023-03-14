package awsdatasync


// Represents the protocol that AWS DataSync uses to access your Amazon FSx for OpenZFS file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protocolProperty := &protocolProperty{
//   	nfs: &nFSProperty{
//   		mountOptions: &mountOptionsProperty{
//   			version: jsii.String("version"),
//   		},
//   	},
//   }
//
type CfnLocationFSxOpenZFS_ProtocolProperty struct {
	// Represents the Network File System (NFS) protocol that DataSync uses to access your FSx for OpenZFS file system.
	Nfs interface{} `field:"optional" json:"nfs" yaml:"nfs"`
}

