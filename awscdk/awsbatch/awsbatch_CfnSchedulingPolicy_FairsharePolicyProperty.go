package awsbatch


// The fair share policy for a scheduling policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fairsharePolicyProperty := &fairsharePolicyProperty{
//   	computeReservation: jsii.Number(123),
//   	shareDecaySeconds: jsii.Number(123),
//   	shareDistribution: []interface{}{
//   		&shareAttributesProperty{
//   			shareIdentifier: jsii.String("shareIdentifier"),
//   			weightFactor: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnSchedulingPolicy_FairsharePolicyProperty struct {
	// A value used to reserve some of the available maximum vCPU for fair share identifiers that have not yet been used.
	//
	// The reserved ratio is `( *computeReservation* /100)^ *ActiveFairShares*` where `*ActiveFairShares*` is the number of active fair share identifiers.
	//
	// For example, a `computeReservation` value of 50 indicates that AWS Batch should reserve 50% of the maximum available vCPU if there is only one fair share identifier, 25% if there are two fair share identifiers, and 12.5% if there are three fair share identifiers. A `computeReservation` value of 25 indicates that AWS Batch should reserve 25% of the maximum available vCPU if there is only one fair share identifier, 6.25% if there are two fair share identifiers, and 1.56% if there are three fair share identifiers.
	//
	// The minimum value is 0 and the maximum value is 99.
	ComputeReservation *float64 `field:"optional" json:"computeReservation" yaml:"computeReservation"`
	// The time period to use to calculate a fair share percentage for each fair share identifier in use, in seconds.
	//
	// A value of zero (0) indicates that only current usage should be measured. The decay allows for more recently run jobs to have more weight than jobs that ran earlier. The maximum supported value is 604800 (1 week).
	ShareDecaySeconds *float64 `field:"optional" json:"shareDecaySeconds" yaml:"shareDecaySeconds"`
	// An array of `SharedIdentifier` objects that contain the weights for the fair share identifiers for the fair share policy.
	//
	// Fair share identifiers that aren't included have a default weight of `1.0` .
	ShareDistribution interface{} `field:"optional" json:"shareDistribution" yaml:"shareDistribution"`
}

