package interfacesawsapigateway


// A reference to a DocumentationVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentationVersionReference := &DocumentationVersionReference{
//   	DocumentationVersion: jsii.String("documentationVersion"),
//   	RestApiId: jsii.String("restApiId"),
//   }
//
type DocumentationVersionReference struct {
	// The DocumentationVersion of the DocumentationVersion resource.
	DocumentationVersion *string `field:"required" json:"documentationVersion" yaml:"documentationVersion"`
	// The RestApiId of the DocumentationVersion resource.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

