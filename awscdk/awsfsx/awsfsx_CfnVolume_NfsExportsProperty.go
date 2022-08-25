package awsfsx


// The configuration object for mounting a Network File System (NFS) file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nfsExportsProperty := &nfsExportsProperty{
//   	clientConfigurations: []interface{}{
//   		&clientConfigurationsProperty{
//   			clients: jsii.String("clients"),
//   			options: []*string{
//   				jsii.String("options"),
//   			},
//   		},
//   	},
//   }
//
type CfnVolume_NfsExportsProperty struct {
	// A list of configuration objects that contain the client and options for mounting the OpenZFS file system.
	ClientConfigurations interface{} `field:"required" json:"clientConfigurations" yaml:"clientConfigurations"`
}

