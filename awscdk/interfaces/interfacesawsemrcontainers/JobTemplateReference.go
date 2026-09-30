package interfacesawsemrcontainers


// A reference to a JobTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobTemplateReference := &JobTemplateReference{
//   	JobTemplateArn: jsii.String("jobTemplateArn"),
//   }
//
type JobTemplateReference struct {
	// The Arn of the JobTemplate resource.
	JobTemplateArn *string `field:"required" json:"jobTemplateArn" yaml:"jobTemplateArn"`
}

