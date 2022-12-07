package awsdatasync


// Specifies the Network File System (NFS) protocol configuration that AWS DataSync uses to access a storage virtual machine (SVM) on your Amazon FSx for NetApp ONTAP file system.
//
// For more information, see [Accessing FSx for ONTAP file systems](https://docs.aws.amazon.com/datasync/latest/userguide/create-ontap-location.html#create-ontap-location-access) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nFSProperty := &nFSProperty{
//   	mountOptions: &nfsMountOptionsProperty{
//   		version: jsii.String("version"),
//   	},
//   }
//
type CfnLocationFSxONTAP_NFSProperty struct {
	// Specifies how DataSync can access a location using the NFS protocol.
	MountOptions interface{} `field:"required" json:"mountOptions" yaml:"mountOptions"`
}

