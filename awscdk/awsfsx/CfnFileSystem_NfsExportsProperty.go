package awsfsx


// The configuration object for mounting a file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nfsExportsProperty := &NfsExportsProperty{
//   	ClientConfigurations: []interface{}{
//   		&ClientConfigurationsProperty{
//   			Clients: jsii.String("clients"),
//   			Options: []*string{
//   				jsii.String("options"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-nfsexports.html
//
type CfnFileSystem_NfsExportsProperty struct {
	// A list of configuration objects that contain the client and options for mounting the OpenZFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-nfsexports.html#cfn-fsx-filesystem-nfsexports-clientconfigurations
	//
	ClientConfigurations interface{} `field:"optional" json:"clientConfigurations" yaml:"clientConfigurations"`
}

