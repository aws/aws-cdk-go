package awsecs


// awslogs provides two modes for delivering messages from the container to the log driver.
//
// Example:
//   var cluster cluster
//
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromAsset(path.resolve(__dirname, jsii.String(".."), jsii.String("eventhandler-image"))),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.NewAwsLogDriver(&awsLogDriverProps{
//   		streamPrefix: jsii.String("EventDemo"),
//   		mode: ecs.awsLogDriverMode_NON_BLOCKING,
//   	}),
//   })
//
//   // An Rule that describes the event trigger (in this case a scheduled run)
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.expression(jsii.String("rate(1 min)")),
//   })
//
//   // Pass an environment variable to the container 'TheContainer' in the task
//   rule.addTarget(targets.NewEcsTask(&ecsTaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	taskCount: jsii.Number(1),
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerName: jsii.String("TheContainer"),
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("I_WAS_TRIGGERED"),
//   					value: jsii.String("From CloudWatch Events"),
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

