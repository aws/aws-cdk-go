package interfacesawssagemaker


// A reference to a CodeRepository resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeRepositoryReference := &CodeRepositoryReference{
//   	CodeRepositoryId: jsii.String("codeRepositoryId"),
//   }
//
type CodeRepositoryReference struct {
	// The Id of the CodeRepository resource.
	CodeRepositoryId *string `field:"required" json:"codeRepositoryId" yaml:"codeRepositoryId"`
}

