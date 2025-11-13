package interfacesawsglue


// A reference to a IntegrationResourceProperty resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationResourcePropertyReference := &IntegrationResourcePropertyReference{
//   	ResourceArn: jsii.String("resourceArn"),
//   	ResourcePropertyArn: jsii.String("resourcePropertyArn"),
//   }
//
type IntegrationResourcePropertyReference struct {
	// The ResourceArn of the IntegrationResourceProperty resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The ResourcePropertyArn of the IntegrationResourceProperty resource.
	ResourcePropertyArn *string `field:"required" json:"resourcePropertyArn" yaml:"resourcePropertyArn"`
}

