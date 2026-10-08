package interfacesawsmediatailor


// A reference to a Program resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   programReference := &ProgramReference{
//   	ProgramArn: jsii.String("programArn"),
//   }
//
type ProgramReference struct {
	// The Arn of the Program resource.
	ProgramArn *string `field:"required" json:"programArn" yaml:"programArn"`
}

