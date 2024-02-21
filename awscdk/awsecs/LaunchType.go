package awsecs


// The launch type of an ECS service.
//
// Example:
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster iCluster
//   var taskDefinition taskDefinition
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   rule.AddTarget(targets.NewEcsTask(&EcsTaskProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	LaunchType: ecs.LaunchType_FARGATE,
//   }))
//
type LaunchType string

const (
	// The service will be launched using the EC2 launch type.
	LaunchType_EC2 LaunchType = "EC2"
	// The service will be launched using the FARGATE launch type.
	LaunchType_FARGATE LaunchType = "FARGATE"
	// The service will be launched using the EXTERNAL launch type.
	LaunchType_EXTERNAL LaunchType = "EXTERNAL"
)

