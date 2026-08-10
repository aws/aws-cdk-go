package interfacesawsbedrock


// A reference to a ModelInvocationJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelInvocationJobReference := &ModelInvocationJobReference{
//   	JobArn: jsii.String("jobArn"),
//   }
//
type ModelInvocationJobReference struct {
	// The JobArn of the ModelInvocationJob resource.
	JobArn *string `field:"required" json:"jobArn" yaml:"jobArn"`
}

