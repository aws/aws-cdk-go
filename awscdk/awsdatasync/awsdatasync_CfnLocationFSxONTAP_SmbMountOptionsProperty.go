package awsdatasync


// Specifies the version of the Server Message Block (SMB) protocol that AWS DataSync uses to access an SMB file server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smbMountOptionsProperty := &smbMountOptionsProperty{
//   	version: jsii.String("version"),
//   }
//
type CfnLocationFSxONTAP_SmbMountOptionsProperty struct {
	// By default, DataSync automatically chooses an SMB protocol version based on negotiation with your SMB file server.
	//
	// You also can configure DataSync to use a specific SMB version, but we recommend doing this only if DataSync has trouble negotiating with the SMB file server automatically.
	//
	// These are the following options for configuring the SMB version:
	//
	// - `AUTOMATIC` (default): DataSync and the SMB file server negotiate a protocol version that they mutually support. (DataSync supports SMB versions 1.0 and later.)
	//
	// This is the recommended option. If you instead choose a specific version that your file server doesn't support, you may get an `Operation Not Supported` error.
	// - `SMB3` : Restricts the protocol negotiation to only SMB version 3.0.2.
	// - `SMB2` : Restricts the protocol negotiation to only SMB version 2.1.
	// - `SMB2_0` : Restricts the protocol negotiation to only SMB version 2.0.
	// - `SMB1` : Restricts the protocol negotiation to only SMB version 1.0.
	//
	// > The `SMB1` option isn't available when [creating an Amazon FSx for NetApp ONTAP location](https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateLocationFsxOntap.html) .
	Version *string `field:"optional" json:"version" yaml:"version"`
}

