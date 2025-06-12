package awsec2


// Options for InitPackage.yum/apt/rubyGem/python.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var initServiceRestartHandle initServiceRestartHandle
//
//   namedPackageOptions := &NamedPackageOptions{
//   	ServiceRestartHandles: []*initServiceRestartHandle{
//   		initServiceRestartHandle,
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
type NamedPackageOptions struct {
	// Restart the given services after this command has run.
	// Default: - Do not restart any service.
	//
	ServiceRestartHandles *[]InitServiceRestartHandle `field:"optional" json:"serviceRestartHandles" yaml:"serviceRestartHandles"`
	// Specify the versions to install.
	// Default: - Install the latest version.
	//
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

