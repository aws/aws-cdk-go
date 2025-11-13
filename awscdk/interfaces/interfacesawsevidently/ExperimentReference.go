package interfacesawsevidently


// A reference to a Experiment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentReference := &ExperimentReference{
//   	ExperimentArn: jsii.String("experimentArn"),
//   }
//
type ExperimentReference struct {
	// The Arn of the Experiment resource.
	ExperimentArn *string `field:"required" json:"experimentArn" yaml:"experimentArn"`
}

