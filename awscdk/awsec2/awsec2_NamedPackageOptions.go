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
//   namedPackageOptions := &namedPackageOptions{
//   	serviceRestartHandles: []*initServiceRestartHandle{
//   		initServiceRestartHandle,
//   	},
//   	version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
type NamedPackageOptions struct {
	// Restart the given services after this command has run.
	ServiceRestartHandles *[]InitServiceRestartHandle `field:"optional" json:"serviceRestartHandles" yaml:"serviceRestartHandles"`
	// Specify the versions to install.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

