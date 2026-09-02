package interfacesawsiot


// A reference to a Job resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobReference := &JobReference{
//   	JobArn: jsii.String("jobArn"),
//   }
//
type JobReference struct {
	// The Arn of the Job resource.
	JobArn *string `field:"required" json:"jobArn" yaml:"jobArn"`
}

