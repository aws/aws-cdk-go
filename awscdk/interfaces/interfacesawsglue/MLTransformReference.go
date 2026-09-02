package interfacesawsglue


// A reference to a MLTransform resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mLTransformReference := &MLTransformReference{
//   	TransformId: jsii.String("transformId"),
//   }
//
type MLTransformReference struct {
	// The TransformId of the MLTransform resource.
	TransformId *string `field:"required" json:"transformId" yaml:"transformId"`
}

