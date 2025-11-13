package interfacesawsapigateway


// A reference to a DocumentationPart resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentationPartReference := &DocumentationPartReference{
//   	DocumentationPartId: jsii.String("documentationPartId"),
//   	RestApiId: jsii.String("restApiId"),
//   }
//
type DocumentationPartReference struct {
	// The DocumentationPartId of the DocumentationPart resource.
	DocumentationPartId *string `field:"required" json:"documentationPartId" yaml:"documentationPartId"`
	// The RestApiId of the DocumentationPart resource.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

