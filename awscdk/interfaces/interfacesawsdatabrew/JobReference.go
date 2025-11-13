package interfacesawsdatabrew


// A reference to a Job resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobReference := &JobReference{
//   	JobName: jsii.String("jobName"),
//   }
//
type JobReference struct {
	// The Name of the Job resource.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
}

