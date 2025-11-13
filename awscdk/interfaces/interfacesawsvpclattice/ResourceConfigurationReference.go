package interfacesawsvpclattice


// A reference to a ResourceConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceConfigurationReference := &ResourceConfigurationReference{
//   	ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//   }
//
type ResourceConfigurationReference struct {
	// The Arn of the ResourceConfiguration resource.
	ResourceConfigurationArn *string `field:"required" json:"resourceConfigurationArn" yaml:"resourceConfigurationArn"`
}

