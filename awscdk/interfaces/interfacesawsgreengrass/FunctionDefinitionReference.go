package interfacesawsgreengrass


// A reference to a FunctionDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionDefinitionReference := &FunctionDefinitionReference{
//   	FunctionDefinitionArn: jsii.String("functionDefinitionArn"),
//   	FunctionDefinitionId: jsii.String("functionDefinitionId"),
//   }
//
type FunctionDefinitionReference struct {
	// The ARN of the FunctionDefinition resource.
	FunctionDefinitionArn *string `field:"required" json:"functionDefinitionArn" yaml:"functionDefinitionArn"`
	// The Id of the FunctionDefinition resource.
	FunctionDefinitionId *string `field:"required" json:"functionDefinitionId" yaml:"functionDefinitionId"`
}

