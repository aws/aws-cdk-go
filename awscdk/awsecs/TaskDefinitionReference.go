package awsecs


// A reference to a TaskDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskDefinitionReference := &TaskDefinitionReference{
//   	TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//   }
//
type TaskDefinitionReference struct {
	// The TaskDefinitionArn of the TaskDefinition resource.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
}

