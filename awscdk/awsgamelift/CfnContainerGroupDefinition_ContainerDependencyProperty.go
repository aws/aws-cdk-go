package awsgamelift


// A container's dependency on another container in the same container group.
//
// The dependency impacts how the dependent container is able to start or shut down based the status of the other container.
//
// For example, *ContainerA* is configured with the following dependency: a `START` dependency on *ContainerB* . This means that *ContainerA* can't start until *ContainerB* has started. It also means that *ContainerA* must shut down before *ContainerB* .
//
// *Part of:* [GameServerContainerDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_GameServerContainerDefinition.html) , [GameServerContainerDefinitionInput](https://docs.aws.amazon.com/gamelift/latest/apireference/API_GameServerContainerDefinitionInput.html) , [SupportContainerDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_SupportContainerDefinition.html) , [SupportContainerDefinitionInput](https://docs.aws.amazon.com/gamelift/latest/apireference/API_SupportContainerDefinitionInput.html)
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
	// The condition that the dependency container must reach before the dependent container can start. Valid conditions include:.
	//
	// - START - The dependency container must have started.
	// - COMPLETE - The dependency container has run to completion (exits). Use this condition with nonessential containers, such as those that run a script and then exit. The dependency container can't be an essential container.
	// - SUCCESS - The dependency container has run to completion and exited with a zero status. The dependency container can't be an essential container.
	// - HEALTHY - The dependency container has passed its Docker health check. Use this condition with dependency containers that have health checks configured. This condition is confirmed at container group startup only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdependency.html#cfn-gamelift-containergroupdefinition-containerdependency-condition
	//
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// A descriptive label for the container definition that this container depends on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdependency.html#cfn-gamelift-containergroupdefinition-containerdependency-containername
	//
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
}

