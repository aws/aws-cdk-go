package awsec2


// Options for InitPackage.rpm/InitPackage.msi.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var initServiceRestartHandle initServiceRestartHandle
//
//   locationPackageOptions := &locationPackageOptions{
//   	key: jsii.String("key"),
//   	serviceRestartHandles: []*initServiceRestartHandle{
//   		initServiceRestartHandle,
//   	},
//   }
//
type LocationPackageOptions struct {
	// Identifier key for this package.
	//
	// You can use this to order package installs.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Restart the given service after this command has run.
	ServiceRestartHandles *[]InitServiceRestartHandle `field:"optional" json:"serviceRestartHandles" yaml:"serviceRestartHandles"`
}

