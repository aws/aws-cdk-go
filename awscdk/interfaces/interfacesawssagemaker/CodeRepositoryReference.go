package interfacesawssagemaker


// A reference to a CodeRepository resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeRepositoryReference := &CodeRepositoryReference{
//   	CodeRepositoryArn: jsii.String("codeRepositoryArn"),
//   }
//
type CodeRepositoryReference struct {
	// The CodeRepositoryArn of the CodeRepository resource.
	CodeRepositoryArn *string `field:"required" json:"codeRepositoryArn" yaml:"codeRepositoryArn"`
}

