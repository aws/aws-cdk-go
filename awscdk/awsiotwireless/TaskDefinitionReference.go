package awsiotwireless


// A reference to a TaskDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskDefinitionReference := &TaskDefinitionReference{
//   	TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//   	TaskDefinitionId: jsii.String("taskDefinitionId"),
//   }
//
type TaskDefinitionReference struct {
	// The ARN of the TaskDefinition resource.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The Id of the TaskDefinition resource.
	TaskDefinitionId *string `field:"required" json:"taskDefinitionId" yaml:"taskDefinitionId"`
}

