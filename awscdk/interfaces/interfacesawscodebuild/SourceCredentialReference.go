package interfacesawscodebuild


// A reference to a SourceCredential resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceCredentialReference := &SourceCredentialReference{
//   	SourceCredentialArn: jsii.String("sourceCredentialArn"),
//   }
//
type SourceCredentialReference struct {
	// The Arn of the SourceCredential resource.
	SourceCredentialArn *string `field:"required" json:"sourceCredentialArn" yaml:"sourceCredentialArn"`
}

