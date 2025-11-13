package interfacesawspinpoint


// A reference to a ApplicationSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationSettingsReference := &ApplicationSettingsReference{
//   	ApplicationSettingsId: jsii.String("applicationSettingsId"),
//   }
//
type ApplicationSettingsReference struct {
	// The Id of the ApplicationSettings resource.
	ApplicationSettingsId *string `field:"required" json:"applicationSettingsId" yaml:"applicationSettingsId"`
}

