package interfacesawsroute53recoveryreadiness


// A reference to a ResourceSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceSetReference := &ResourceSetReference{
//   	ResourceSetArn: jsii.String("resourceSetArn"),
//   	ResourceSetName: jsii.String("resourceSetName"),
//   }
//
type ResourceSetReference struct {
	// The ARN of the ResourceSet resource.
	ResourceSetArn *string `field:"required" json:"resourceSetArn" yaml:"resourceSetArn"`
	// The ResourceSetName of the ResourceSet resource.
	ResourceSetName *string `field:"required" json:"resourceSetName" yaml:"resourceSetName"`
}

