package interfacesawspinpointemail


// A reference to a ConfigurationSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationSetReference := &ConfigurationSetReference{
//   	ConfigurationSetId: jsii.String("configurationSetId"),
//   }
//
type ConfigurationSetReference struct {
	// The Id of the ConfigurationSet resource.
	ConfigurationSetId *string `field:"required" json:"configurationSetId" yaml:"configurationSetId"`
}

