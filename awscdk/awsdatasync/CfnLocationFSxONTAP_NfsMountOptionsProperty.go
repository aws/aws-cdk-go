package awsdatasync


// Specifies how DataSync can access a location using the NFS protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nfsMountOptionsProperty := &NfsMountOptionsProperty{
//   	Version: jsii.String("version"),
//   }
//
type CfnLocationFSxONTAP_NfsMountOptionsProperty struct {
	// Specifies the NFS version that you want DataSync to use when mounting your NFS share.
	//
	// If the server refuses to use the version specified, the task fails.
	//
	// You can specify the following options:
	//
	// - `AUTOMATIC` (default): DataSync chooses NFS version 4.1.
	// - `NFS3` : Stateless protocol version that allows for asynchronous writes on the server.
	// - `NFSv4_0` : Stateful, firewall-friendly protocol version that supports delegations and pseudo file systems.
	// - `NFSv4_1` : Stateful protocol version that supports sessions, directory delegations, and parallel data processing. NFS version 4.1 also includes all features available in version 4.0.
	//
	// > DataSync currently only supports NFS version 3 with Amazon FSx for NetApp ONTAP locations.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

