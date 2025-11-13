package interfacesawsbedrockagentcore


// A reference to a CodeInterpreterCustom resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeInterpreterCustomReference := &CodeInterpreterCustomReference{
//   	CodeInterpreterId: jsii.String("codeInterpreterId"),
//   }
//
type CodeInterpreterCustomReference struct {
	// The CodeInterpreterId of the CodeInterpreterCustom resource.
	CodeInterpreterId *string `field:"required" json:"codeInterpreterId" yaml:"codeInterpreterId"`
}

