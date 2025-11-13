package interfacesawspersonalize


// A reference to a Solution resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   solutionReference := &SolutionReference{
//   	SolutionArn: jsii.String("solutionArn"),
//   }
//
type SolutionReference struct {
	// The SolutionArn of the Solution resource.
	SolutionArn *string `field:"required" json:"solutionArn" yaml:"solutionArn"`
}

