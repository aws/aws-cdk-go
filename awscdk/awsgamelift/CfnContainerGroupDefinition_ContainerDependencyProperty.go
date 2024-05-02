package awsgamelift


// A dependency that impacts a container's startup and shutdown.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerDependencyProperty := &ContainerDependencyProperty{
//   	Condition: jsii.String("condition"),
//   	ContainerName: jsii.String("containerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdependency.html
//
type CfnContainerGroupDefinition_ContainerDependencyProperty struct {
	// The type of dependency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdependency.html#cfn-gamelift-containergroupdefinition-containerdependency-condition
	//
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// A descriptive label for the container definition.
	//
	// The container being defined depends on this container's condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdependency.html#cfn-gamelift-containergroupdefinition-containerdependency-containername
	//
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
}

