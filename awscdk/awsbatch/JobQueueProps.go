package awsbatch


// Props to configure a JobQueue.
//
// Example:
//   var vpc iVpc
//
//
//   ecsJob := batch.NewEcsJobDefinition(this, jsii.String("JobDefn"), &EcsJobDefinitionProps{
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("containerDefn"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/amazonlinux/amazonlinux:latest")),
//   		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   		Cpu: jsii.Number(256),
//   	}),
//   })
//
//   queue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	ComputeEnvironments: []orderedComputeEnvironment{
//   		&orderedComputeEnvironment{
//   			ComputeEnvironment: batch.NewManagedEc2EcsComputeEnvironment(this, jsii.String("managedEc2CE"), &ManagedEc2EcsComputeEnvironmentProps{
//   				Vpc: *Vpc,
//   			}),
//   			Order: jsii.Number(1),
//   		},
//   	},
//   	Priority: jsii.Number(10),
//   })
//
//   user := iam.NewUser(this, jsii.String("MyUser"))
//   ecsJob.GrantSubmitJob(user, queue)
//
type JobQueueProps struct {
	// The set of compute environments mapped to a job queue and their order relative to each other.
	//
	// The job scheduler uses this parameter to determine which compute environment runs a specific job.
	// Compute environments must be in the VALID state before you can associate them with a job queue.
	// You can associate up to three compute environments with a job queue.
	// All of the compute environments must be either EC2 (EC2 or SPOT) or Fargate (FARGATE or FARGATE_SPOT);
	// EC2 and Fargate compute environments can't be mixed.
	//
	// *Note*: All compute environments that are associated with a job queue must share the same architecture.
	// AWS Batch doesn't support mixing compute environment architecture types in a single job queue.
	// Default: none.
	//
	ComputeEnvironments *[]*OrderedComputeEnvironment `field:"optional" json:"computeEnvironments" yaml:"computeEnvironments"`
	// If the job queue is enabled, it is able to accept jobs.
	//
	// Otherwise, new jobs can't be added to the queue, but jobs already in the queue can finish.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The name of the job queue.
	//
	// It can be up to 128 letters long.
	// It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	// Default: - no name.
	//
	JobQueueName *string `field:"optional" json:"jobQueueName" yaml:"jobQueueName"`
	// The priority of the job queue.
	//
	// Job queues with a higher priority are evaluated first when associated with the same compute environment.
	// Priority is determined in descending order.
	// For example, a job queue with a priority of 10 is given scheduling preference over a job queue with a priority of 1.
	// Default: 1.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The SchedulingPolicy for this JobQueue.
	//
	// Instructs the Scheduler how to schedule different jobs.
	// Default: - no scheduling policy.
	//
	SchedulingPolicy ISchedulingPolicy `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
}

