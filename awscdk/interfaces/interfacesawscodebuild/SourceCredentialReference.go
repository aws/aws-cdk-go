package interfacesawscodebuild


// A reference to a SourceCredential resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceCredentialReference := &SourceCredentialReference{
//   	SourceCredentialId: jsii.String("sourceCredentialId"),
//   }
//
type SourceCredentialReference struct {
	// The Id of the SourceCredential resource.
	SourceCredentialId *string `field:"required" json:"sourceCredentialId" yaml:"sourceCredentialId"`
}

