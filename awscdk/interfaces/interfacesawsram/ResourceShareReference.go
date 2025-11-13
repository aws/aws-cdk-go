package interfacesawsram


// A reference to a ResourceShare resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceShareReference := &ResourceShareReference{
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   }
//
type ResourceShareReference struct {
	// The Arn of the ResourceShare resource.
	ResourceShareArn *string `field:"required" json:"resourceShareArn" yaml:"resourceShareArn"`
}

