package interfacesawsdeadline


// A reference to a Worker resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerReference := &WorkerReference{
//   	WorkerArn: jsii.String("workerArn"),
//   }
//
type WorkerReference struct {
	// The Arn of the Worker resource.
	WorkerArn *string `field:"required" json:"workerArn" yaml:"workerArn"`
}

