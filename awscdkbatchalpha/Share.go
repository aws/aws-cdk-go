package awscdkbatchalpha


// Represents a group of Job Definitions.
//
// All Job Definitions that
// declare a share identifier will be considered members of the Share
// defined by that share identifier.
//
// The Scheduler divides the maximum available vCPUs of the ComputeEnvironment
// among Jobs in the Queue based on their shareIdentifier and the weightFactor
// associated with that shareIdentifier.
//
// Example:
//   fairsharePolicy := batch.NewFairshareSchedulingPolicy(this, jsii.String("myFairsharePolicy"))
//
//   fairsharePolicy.AddShare(&Share{
//   	ShareIdentifier: jsii.String("A"),
//   	WeightFactor: jsii.Number(1),
//   })
//   fairsharePolicy.AddShare(&Share{
//   	ShareIdentifier: jsii.String("B"),
//   	WeightFactor: jsii.Number(1),
//   })
//   batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	SchedulingPolicy: fairsharePolicy,
//   })
//
// Experimental.
type Share struct {
	// The identifier of this Share.
	//
	// All jobs that specify this share identifier
	// when submitted to the queue will be considered as part of this Share.
	// Experimental.
	ShareIdentifier *string `field:"required" json:"shareIdentifier" yaml:"shareIdentifier"`
	// The weight factor given to this Share.
	//
	// The Scheduler decides which jobs to put in the Compute Environment
	// such that the following ratio is equal for each job:
	//
	// `sharevCpu / weightFactor`,
	//
	// where `sharevCpu` is the total amount of vCPU given to that particular share; that is,
	// the sum of the vCPU of each job currently in the Compute Environment for that share.
	//
	// See the readme of this module for a detailed example that shows how these are used,
	// how it relates to `computeReservation`, and how `shareDecay` affects these calculations.
	// Experimental.
	WeightFactor *float64 `field:"required" json:"weightFactor" yaml:"weightFactor"`
}

