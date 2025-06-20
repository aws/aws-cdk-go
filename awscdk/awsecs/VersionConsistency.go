package awsecs


// State of the container version consistency feature.
//
// Example:
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
//   	VersionConsistency: ecs.VersionConsistency_DISABLED,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinition.html#cfn-ecs-taskdefinition-containerdefinition-versionconsistency
//
type VersionConsistency string

const (
	// The version consistency feature is enabled for this container.
	VersionConsistency_ENABLED VersionConsistency = "ENABLED"
	// The version consistency feature is disabled for this container.
	VersionConsistency_DISABLED VersionConsistency = "DISABLED"
)

