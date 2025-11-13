package interfacesawsarczonalshift


// A reference to a ZonalAutoshiftConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zonalAutoshiftConfigurationReference := &ZonalAutoshiftConfigurationReference{
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   }
//
type ZonalAutoshiftConfigurationReference struct {
	// The ResourceIdentifier of the ZonalAutoshiftConfiguration resource.
	ResourceIdentifier *string `field:"required" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

