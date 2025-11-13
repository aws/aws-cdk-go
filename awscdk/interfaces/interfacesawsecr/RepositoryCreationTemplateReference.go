package interfacesawsecr


// A reference to a RepositoryCreationTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryCreationTemplateReference := &RepositoryCreationTemplateReference{
//   	Prefix: jsii.String("prefix"),
//   }
//
type RepositoryCreationTemplateReference struct {
	// The Prefix of the RepositoryCreationTemplate resource.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

