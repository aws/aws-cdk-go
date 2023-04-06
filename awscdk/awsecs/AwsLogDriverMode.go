package awsecs


// awslogs provides two modes for delivering messages from the container to the log driver.
//
// Example:
//   var cluster cluster
//
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromAsset(path.resolve(__dirname, jsii.String(".."), jsii.String("eventhandler-image"))),
//   	MemoryLimitMiB: jsii.Number(256),
//   	Logging: ecs.NewAwsLogDriver(&AwsLogDriverProps{
//   		StreamPrefix: jsii.String("EventDemo"),
//   		Mode: ecs.AwsLogDriverMode_NON_BLOCKING,
//   	}),
//   })
//
//   // An Rule that describes the event trigger (in this case a scheduled run)
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Expression(jsii.String("rate(1 min)")),
//   })
//
//   // Pass an environment variable to the container 'TheContainer' in the task
//   rule.AddTarget(targets.NewEcsTask(&EcsTaskProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	TaskCount: jsii.Number(1),
//   	ContainerOverrides: []containerOverride{
//   		&containerOverride{
//   			ContainerName: jsii.String("TheContainer"),
//   			Environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					Name: jsii.String("I_WAS_TRIGGERED"),
//   					Value: jsii.String("From CloudWatch Events"),
//   				},
//   			},
//   		},
//   	},
//   }))
//
type AwsLogDriverMode string

const (
	// (default) direct, blocking delivery from container to driver.
	AwsLogDriverMode_BLOCKING AwsLogDriverMode = "BLOCKING"
	// The non-blocking message delivery mode prevents applications from blocking due to logging back pressure.
	//
	// Applications are likely to fail in unexpected ways when STDERR or STDOUT streams block.
	AwsLogDriverMode_NON_BLOCKING AwsLogDriverMode = "NON_BLOCKING"
)

