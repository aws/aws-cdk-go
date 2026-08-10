package interfacesawsbedrockagentcore


// A reference to a CodeInterpreter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeInterpreterReference := &CodeInterpreterReference{
//   	CodeInterpreterArn: jsii.String("codeInterpreterArn"),
//   }
//
type CodeInterpreterReference struct {
	// The CodeInterpreterArn of the CodeInterpreter resource.
	CodeInterpreterArn *string `field:"required" json:"codeInterpreterArn" yaml:"codeInterpreterArn"`
}

