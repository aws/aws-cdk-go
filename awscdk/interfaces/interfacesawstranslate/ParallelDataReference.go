package interfacesawstranslate


// A reference to a ParallelData resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parallelDataReference := &ParallelDataReference{
//   	ParallelDataArn: jsii.String("parallelDataArn"),
//   }
//
type ParallelDataReference struct {
	// The Arn of the ParallelData resource.
	ParallelDataArn *string `field:"required" json:"parallelDataArn" yaml:"parallelDataArn"`
}

