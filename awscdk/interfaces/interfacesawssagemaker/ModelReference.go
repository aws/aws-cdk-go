package interfacesawssagemaker


// A reference to a Model resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelReference := &ModelReference{
//   	ModelId: jsii.String("modelId"),
//   }
//
type ModelReference struct {
	// The Id of the Model resource.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
}

