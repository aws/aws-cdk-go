package previewawslogs


// Properties for XRay delivery destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   xRayDeliveryDestinationProps := &XRayDeliveryDestinationProps{
//   	SourceResource: jsii.String("sourceResource"),
//   }
//
// Experimental.
type XRayDeliveryDestinationProps struct {
	// Arn of the source resource.
	// Experimental.
	SourceResource *string `field:"required" json:"sourceResource" yaml:"sourceResource"`
}

