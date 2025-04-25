package awsecs


// Indicates whether to use Availability Zone rebalancing for an ECS service.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	AvailabilityZoneRebalancing: ecs.AvailabilityZoneRebalancing_ENABLED,
//   })
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-rebalancing.html
//
type AvailabilityZoneRebalancing string

const (
	// Availability zone rebalancing is enabled.
	AvailabilityZoneRebalancing_ENABLED AvailabilityZoneRebalancing = "ENABLED"
	// Availability zone rebalancing is disabled.
	AvailabilityZoneRebalancing_DISABLED AvailabilityZoneRebalancing = "DISABLED"
)

