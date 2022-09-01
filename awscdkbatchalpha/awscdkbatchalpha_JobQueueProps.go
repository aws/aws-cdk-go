// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// Properties of a batch job queue.
//
// Example:
//   var computeEnvironment computeEnvironment
//
//   jobQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &jobQueueProps{
//   	computeEnvironments: []jobQueueComputeEnvironment{
//   		&jobQueueComputeEnvironment{
//   			// Defines a collection of compute resources to handle assigned batch jobs
//   			computeEnvironment: computeEnvironment,
//   			// Order determines the allocation order for jobs (i.e. Lower means higher preference for job assignment)
//   			order: jsii.Number(1),
//   		},
//   	},
//   })
//
// Experimental.
type JobQueueProps struct {
	// The set of compute environments mapped to a job queue and their order relative to each other.
	//
	// The job scheduler uses this parameter to
	// determine which compute environment should execute a given job. Compute environments must be in the VALID state before you can associate them
	// with a job queue. You can associate up to three compute environments with a job queue.
	// Experimental.
	ComputeEnvironments *[]*JobQueueComputeEnvironment `field:"required" json:"computeEnvironments" yaml:"computeEnvironments"`
	// The state of the job queue.
	//
	// If set to true, it is able to accept jobs.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// A name for the job queue.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	JobQueueName *string `field:"optional" json:"jobQueueName" yaml:"jobQueueName"`
	// The priority of the job queue.
	//
	// Job queues with a higher priority (or a higher integer value for the priority parameter) are evaluated first
	// when associated with the same compute environment. Priority is determined in descending order, for example, a job queue with a priority value
	// of 10 is given scheduling preference over a job queue with a priority value of 1.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

