package mixinsawsdatasync


// Specifies the options that DataSync can use to mount your NFS file server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mountOptionsProperty := &MountOptionsProperty{
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationnfs-mountoptions.html
//
type CfnLocationNFSPropsMixin_MountOptionsProperty struct {
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationnfs-mountoptions.html#cfn-datasync-locationnfs-mountoptions-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

