package interfacesawsbedrockagentcore


// A reference to a ConfigurationBundle resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationBundleReference := &ConfigurationBundleReference{
//   	BundleArn: jsii.String("bundleArn"),
//   }
//
type ConfigurationBundleReference struct {
	// The BundleArn of the ConfigurationBundle resource.
	BundleArn *string `field:"required" json:"bundleArn" yaml:"bundleArn"`
}

