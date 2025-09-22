package awsgreengrass


// A reference to a FunctionDefinitionVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionDefinitionVersionReference := &FunctionDefinitionVersionReference{
//   	FunctionDefinitionVersionId: jsii.String("functionDefinitionVersionId"),
//   }
//
type FunctionDefinitionVersionReference struct {
	// The Id of the FunctionDefinitionVersion resource.
	FunctionDefinitionVersionId *string `field:"required" json:"functionDefinitionVersionId" yaml:"functionDefinitionVersionId"`
}

