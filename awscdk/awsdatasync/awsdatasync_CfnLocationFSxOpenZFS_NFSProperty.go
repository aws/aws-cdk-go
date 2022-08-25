package awsdatasync


// Represents the Network File System (NFS) protocol that AWS DataSync uses to access your Amazon FSx for OpenZFS file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nFSProperty := &nFSProperty{
//   	mountOptions: &mountOptionsProperty{
//   		version: jsii.String("version"),
//   	},
//   }
//
type CfnLocationFSxOpenZFS_NFSProperty struct {
	// Represents the mount options that are available for DataSync to access an NFS location.
	MountOptions interface{} `field:"required" json:"mountOptions" yaml:"mountOptions"`
}

