package awsconnect


// A reference to a TaskTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskTemplateReference := &TaskTemplateReference{
//   	TaskTemplateArn: jsii.String("taskTemplateArn"),
//   }
//
type TaskTemplateReference struct {
	// The Arn of the TaskTemplate resource.
	TaskTemplateArn *string `field:"required" json:"taskTemplateArn" yaml:"taskTemplateArn"`
}

