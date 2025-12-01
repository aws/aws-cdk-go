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
//   	CodeRepositoryName: jsii.String("codeRepositoryName"),
//   }
//
type CodeRepositoryReference struct {
	// The Id of the CodeRepository resource.
	CodeRepositoryId *string `field:"required" json:"codeRepositoryId" yaml:"codeRepositoryId"`
	// The CodeRepositoryName of the CodeRepository resource.
	CodeRepositoryName *string `field:"required" json:"codeRepositoryName" yaml:"codeRepositoryName"`
}

