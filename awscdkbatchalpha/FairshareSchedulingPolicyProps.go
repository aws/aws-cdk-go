// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Fairshare SchedulingPolicy configuration.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   fairsharePolicy := batch.NewFairshareSchedulingPolicy(this, jsii.String("myFairsharePolicy"), &FairshareSchedulingPolicyProps{
//   	ShareDecay: cdk.Duration_Minutes(jsii.Number(5)),
//   })
//
// Experimental.
type FairshareSchedulingPolicyProps struct {
	// Used to calculate the percentage of the maximum available vCPU to reserve for share identifiers not present in the Queue.
	//
	// The percentage reserved is defined by the Scheduler as:
	// `(computeReservation/100)^ActiveFairShares` where `ActiveFairShares` is the number of active fair share identifiers.
	//
	// For example, a computeReservation value of 50 indicates that AWS Batch reserves 50% of the
	// maximum available vCPU if there's only one fair share identifier.
	// It reserves 25% if there are two fair share identifiers.
	// It reserves 12.5% if there are three fair share identifiers.
	//
	// A computeReservation value of 25 indicates that AWS Batch should reserve 25% of the
	// maximum available vCPU if there's only one fair share identifier,
	// 6.25% if there are two fair share identifiers,
	// and 1.56% if there are three fair share identifiers.
	// Experimental.
	ComputeReservation *float64 `field:"optional" json:"computeReservation" yaml:"computeReservation"`
	// The name of this SchedulingPolicy.
	// Experimental.
	SchedulingPolicyName *string `field:"optional" json:"schedulingPolicyName" yaml:"schedulingPolicyName"`
	// The amount of time to use to measure the usage of each job.
	//
	// The usage is used to calculate a fair share percentage for each fair share identifier currently in the Queue.
	// A value of zero (0) indicates that only current usage is measured.
	// The decay is linear and gives preference to newer jobs.
	//
	// The maximum supported value is 604800 seconds (1 week).
	// Experimental.
	ShareDecay awscdk.Duration `field:"optional" json:"shareDecay" yaml:"shareDecay"`
	// The shares that this Scheduling Policy applies to.
	//
	// *Note*: It is possible to submit Jobs to the queue with Share Identifiers that
	// are not recognized by the Scheduling Policy.
	// Experimental.
	Shares *[]*Share `field:"optional" json:"shares" yaml:"shares"`
}

