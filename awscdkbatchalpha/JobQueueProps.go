package awscdkbatchalpha


// Props to configure a JobQueue.
//
// Example:
//   var vpc iVpc
//
//   sharedComputeEnv := batch.NewFargateComputeEnvironment(this, jsii.String("spotEnv"), &FargateComputeEnvironmentProps{
//   	Vpc: Vpc,
//   	Spot: jsii.Boolean(true),
//   })
//   lowPriorityQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	Priority: jsii.Number(1),
//   })
//   highPriorityQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	Priority: jsii.Number(10),
//   })
//   lowPriorityQueue.AddComputeEnvironment(sharedComputeEnv, jsii.Number(1))
//   highPriorityQueue.AddComputeEnvironment(sharedComputeEnv, jsii.Number(1))
//
// Experimental.
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
	// Experimental.
	ComputeEnvironments *[]*OrderedComputeEnvironment `field:"optional" json:"computeEnvironments" yaml:"computeEnvironments"`
	// If the job queue is enabled, it is able to accept jobs.
	//
	// Otherwise, new jobs can't be added to the queue, but jobs already in the queue can finish.
	// Default: true.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The name of the job queue.
	//
	// It can be up to 128 letters long.
	// It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	// Default: - no name.
	//
	// Experimental.
	JobQueueName *string `field:"optional" json:"jobQueueName" yaml:"jobQueueName"`
	// The priority of the job queue.
	//
	// Job queues with a higher priority are evaluated first when associated with the same compute environment.
	// Priority is determined in descending order.
	// For example, a job queue with a priority of 10 is given scheduling preference over a job queue with a priority of 1.
	// Default: 1.
	//
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The SchedulingPolicy for this JobQueue.
	//
	// Instructs the Scheduler how to schedule different jobs.
	// Default: - no scheduling policy.
	//
	// Experimental.
	SchedulingPolicy ISchedulingPolicy `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
}

