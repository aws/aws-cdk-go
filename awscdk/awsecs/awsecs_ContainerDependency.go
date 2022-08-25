package awsecs


// The details of a dependency on another container in the task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerDefinition containerDefinition
//
//   containerDependency := &containerDependency{
//   	container: containerDefinition,
//
//   	// the properties below are optional
//   	condition: awscdk.Aws_ecs.containerDependencyCondition_START,
//   }
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDependency.html
//
type ContainerDependency struct {
	// The container to depend on.
	Container ContainerDefinition `field:"required" json:"container" yaml:"container"`
	// The state the container needs to be in to satisfy the dependency and proceed with startup.
	//
	// Valid values are ContainerDependencyCondition.START, ContainerDependencyCondition.COMPLETE,
	// ContainerDependencyCondition.SUCCESS and ContainerDependencyCondition.HEALTHY.
	Condition ContainerDependencyCondition `field:"optional" json:"condition" yaml:"condition"`
}

