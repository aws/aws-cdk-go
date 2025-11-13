package interfacesawsapigateway


// A reference to a Model resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelReference := &ModelReference{
//   	ModelName: jsii.String("modelName"),
//   }
//
type ModelReference struct {
	// The Name of the Model resource.
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
}

