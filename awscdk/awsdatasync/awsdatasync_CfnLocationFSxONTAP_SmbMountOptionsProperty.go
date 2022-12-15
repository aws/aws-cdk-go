package awsdatasync


// Specifies how DataSync can access a location using the SMB protocol.
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
	// Specifies the SMB version that you want DataSync to use when mounting your SMB share.
	//
	// If you don't specify a version, DataSync defaults to `AUTOMATIC` and chooses a version based on negotiation with the SMB server.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

