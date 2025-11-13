package interfacesawsglue


// A reference to a MLTransform resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mLTransformReference := &MLTransformReference{
//   	MlTransformId: jsii.String("mlTransformId"),
//   }
//
type MLTransformReference struct {
	// The Id of the MLTransform resource.
	MlTransformId *string `field:"required" json:"mlTransformId" yaml:"mlTransformId"`
}

