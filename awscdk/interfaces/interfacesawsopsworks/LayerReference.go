package interfacesawsopsworks


// A reference to a Layer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   layerReference := &LayerReference{
//   	LayerId: jsii.String("layerId"),
//   }
//
type LayerReference struct {
	// The Id of the Layer resource.
	LayerId *string `field:"required" json:"layerId" yaml:"layerId"`
}

