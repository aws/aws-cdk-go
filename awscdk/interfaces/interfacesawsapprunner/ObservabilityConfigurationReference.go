package interfacesawsapprunner


// A reference to a ObservabilityConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   observabilityConfigurationReference := &ObservabilityConfigurationReference{
//   	ObservabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   }
//
type ObservabilityConfigurationReference struct {
	// The ObservabilityConfigurationArn of the ObservabilityConfiguration resource.
	ObservabilityConfigurationArn *string `field:"required" json:"observabilityConfigurationArn" yaml:"observabilityConfigurationArn"`
}

