package interfacesawsdatasync


// A reference to a Task resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskReference := &TaskReference{
//   	TaskArn: jsii.String("taskArn"),
//   }
//
type TaskReference struct {
	// The TaskArn of the Task resource.
	TaskArn *string `field:"required" json:"taskArn" yaml:"taskArn"`
}

