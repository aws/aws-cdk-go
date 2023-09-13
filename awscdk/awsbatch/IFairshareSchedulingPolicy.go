package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents a Fairshare Scheduling Policy. Instructs the scheduler to allocate ComputeEnvironment vCPUs based on Job shareIdentifiers.
//
// The Faireshare Scheduling Policy ensures that each share gets a certain amount of vCPUs.
// It does this by deciding how many Jobs of each share to schedule *relative to how many jobs of
// each share are currently being executed by the ComputeEnvironment*. The weight factors associated with
// each share determine the ratio of vCPUs allocated; see the readme for a more in-depth discussion of
// fairshare policies.
type IFairshareSchedulingPolicy interface {
	ISchedulingPolicy
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
	// Default: - no vCPU is reserved.
	//
	ComputeReservation() *float64
	// The amount of time to use to measure the usage of each job.
	//
	// The usage is used to calculate a fair share percentage for each fair share identifier currently in the Queue.
	// A value of zero (0) indicates that only current usage is measured.
	// The decay is linear and gives preference to newer jobs.
	//
	// The maximum supported value is 604800 seconds (1 week).
	// Default: - 0: only the current job usage is considered.
	//
	ShareDecay() awscdk.Duration
	// The shares that this Scheduling Policy applies to.
	//
	// *Note*: It is possible to submit Jobs to the queue with Share Identifiers that
	// are not recognized by the Scheduling Policy.
	Shares() *[]*Share
}

// The jsii proxy for IFairshareSchedulingPolicy
type jsiiProxy_IFairshareSchedulingPolicy struct {
	jsiiProxy_ISchedulingPolicy
}

func (j *jsiiProxy_IFairshareSchedulingPolicy) ComputeReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFairshareSchedulingPolicy) ShareDecay() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"shareDecay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFairshareSchedulingPolicy) Shares() *[]*Share {
	var returns *[]*Share
	_jsii_.Get(
		j,
		"shares",
		&returns,
	)
	return returns
}

