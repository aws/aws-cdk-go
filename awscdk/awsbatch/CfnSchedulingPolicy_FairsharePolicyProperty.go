package awsbatch


// The fair share policy for a scheduling policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fairsharePolicyProperty := &FairsharePolicyProperty{
//   	ComputeReservation: jsii.Number(123),
//   	ShareDecaySeconds: jsii.Number(123),
//   	ShareDistribution: []interface{}{
//   		&ShareAttributesProperty{
//   			ShareIdentifier: jsii.String("shareIdentifier"),
//   			WeightFactor: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-schedulingpolicy-fairsharepolicy.html
//
type CfnSchedulingPolicy_FairsharePolicyProperty struct {
	// A value used to reserve some of the available maximum vCPU for fair share identifiers that aren't already used.
	//
	// The reserved ratio is `( *computeReservation* /100)^ *ActiveFairShares*` where `*ActiveFairShares*` is the number of active fair share identifiers.
	//
	// For example, a `computeReservation` value of 50 indicates that AWS Batch reserves 50% of the maximum available vCPU if there's only one fair share identifier. It reserves 25% if there are two fair share identifiers. It reserves 12.5% if there are three fair share identifiers. A `computeReservation` value of 25 indicates that AWS Batch should reserve 25% of the maximum available vCPU if there's only one fair share identifier, 6.25% if there are two fair share identifiers, and 1.56% if there are three fair share identifiers.
	//
	// The minimum value is 0 and the maximum value is 99.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-schedulingpolicy-fairsharepolicy.html#cfn-batch-schedulingpolicy-fairsharepolicy-computereservation
	//
	ComputeReservation *float64 `field:"optional" json:"computeReservation" yaml:"computeReservation"`
	// The amount of time (in seconds) to use to calculate a fair share percentage for each fair share identifier in use.
	//
	// A value of zero (0) indicates the default minimum time window (600 seconds). The maximum supported value is 604800 (1 week).
	//
	// The decay allows for more recently run jobs to have more weight than jobs that ran earlier. Consider adjusting this number if you have jobs that (on average) run longer than ten minutes, or a large difference in job count or job run times between share identifiers, and the allocation of resources doesn’t meet your needs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-schedulingpolicy-fairsharepolicy.html#cfn-batch-schedulingpolicy-fairsharepolicy-sharedecayseconds
	//
	ShareDecaySeconds *float64 `field:"optional" json:"shareDecaySeconds" yaml:"shareDecaySeconds"`
	// An array of `SharedIdentifier` objects that contain the weights for the fair share identifiers for the fair share policy.
	//
	// Fair share identifiers that aren't included have a default weight of `1.0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-schedulingpolicy-fairsharepolicy.html#cfn-batch-schedulingpolicy-fairsharepolicy-sharedistribution
	//
	ShareDistribution interface{} `field:"optional" json:"shareDistribution" yaml:"shareDistribution"`
}

