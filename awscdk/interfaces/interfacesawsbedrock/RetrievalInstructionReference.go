package interfacesawsbedrock


// A reference to a RetrievalInstruction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retrievalInstructionReference := &RetrievalInstructionReference{
//   	RetrievalInstructionArn: jsii.String("retrievalInstructionArn"),
//   }
//
type RetrievalInstructionReference struct {
	// The Arn of the RetrievalInstruction resource.
	RetrievalInstructionArn *string `field:"required" json:"retrievalInstructionArn" yaml:"retrievalInstructionArn"`
}

