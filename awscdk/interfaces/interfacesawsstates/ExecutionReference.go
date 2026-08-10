package interfacesawsstates


// A reference to a Execution resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionReference := &ExecutionReference{
//   	ExecutionArn: jsii.String("executionArn"),
//   }
//
type ExecutionReference struct {
	// The ExecutionArn of the Execution resource.
	ExecutionArn *string `field:"required" json:"executionArn" yaml:"executionArn"`
}

