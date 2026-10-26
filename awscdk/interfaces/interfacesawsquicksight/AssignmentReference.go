package interfacesawsquicksight


// A reference to a Assignment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assignmentReference := &AssignmentReference{
//   	AssignmentArn: jsii.String("assignmentArn"),
//   }
//
type AssignmentReference struct {
	// The Arn of the Assignment resource.
	AssignmentArn *string `field:"required" json:"assignmentArn" yaml:"assignmentArn"`
}

