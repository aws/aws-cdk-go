package interfacesawsses


// A reference to a ConfigurationSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationSetReference := &ConfigurationSetReference{
//   	ConfigurationSetName: jsii.String("configurationSetName"),
//   }
//
type ConfigurationSetReference struct {
	// The Name of the ConfigurationSet resource.
	ConfigurationSetName *string `field:"required" json:"configurationSetName" yaml:"configurationSetName"`
}

