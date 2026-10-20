package interfacesawssmsvoice


// A reference to a NotifyConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notifyConfigurationReference := &NotifyConfigurationReference{
//   	NotifyConfigurationArn: jsii.String("notifyConfigurationArn"),
//   }
//
type NotifyConfigurationReference struct {
	// The NotifyConfigurationArn of the NotifyConfiguration resource.
	NotifyConfigurationArn *string `field:"required" json:"notifyConfigurationArn" yaml:"notifyConfigurationArn"`
}

