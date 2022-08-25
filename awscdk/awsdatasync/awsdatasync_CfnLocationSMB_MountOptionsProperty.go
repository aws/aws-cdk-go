package awsdatasync


// The mount options used by DataSync to access the SMB server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mountOptionsProperty := &mountOptionsProperty{
//   	version: jsii.String("version"),
//   }
//
type CfnLocationSMB_MountOptionsProperty struct {
	// The specific SMB version that you want DataSync to use to mount your SMB share.
	//
	// If you don't specify a version, DataSync defaults to `AUTOMATIC` . That is, DataSync automatically selects a version based on negotiation with the SMB server.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

