package awscodebuild


// The compute fleet overflow behavior.
//
// Example:
//   fleet := codebuild.NewFleet(this, jsii.String("Fleet"), &FleetProps{
//   	ComputeType: codebuild.FleetComputeType_MEDIUM,
//   	EnvironmentType: codebuild.EnvironmentType_LINUX_CONTAINER,
//   	BaseCapacity: jsii.Number(1),
//   	OverflowBehavior: codebuild.FleetOverflowBehavior_ON_DEMAND,
//   })
//
type FleetOverflowBehavior string

const (
	// Overflow builds wait for existing fleet instances to become available.
	FleetOverflowBehavior_QUEUE FleetOverflowBehavior = "QUEUE"
	// Overflow builds run on CodeBuild on-demand instances.
	FleetOverflowBehavior_ON_DEMAND FleetOverflowBehavior = "ON_DEMAND"
)

