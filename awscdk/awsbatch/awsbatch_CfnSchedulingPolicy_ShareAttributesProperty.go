package awsbatch


// Specifies the weights for the fair share identifiers for the fair share policy.
//
// Fair share identifiers that aren't included have a default weight of `1.0` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shareAttributesProperty := &shareAttributesProperty{
//   	shareIdentifier: jsii.String("shareIdentifier"),
//   	weightFactor: jsii.Number(123),
//   }
//
type CfnSchedulingPolicy_ShareAttributesProperty struct {
	// A fair share identifier or fair share identifier prefix.
	//
	// If the string ends with an asterisk (*), this entry specifies the weight factor to use for fair share identifiers that start with that prefix. The list of fair share identifiers in a fair share policy cannot overlap. For example, you can't have one that specifies a `shareIdentifier` of `UserA*` and another that specifies a `shareIdentifier` of `UserA-1` .
	//
	// There can be no more than 500 fair share identifiers active in a job queue.
	//
	// The string is limited to 255 alphanumeric characters, optionally followed by an asterisk (*).
	ShareIdentifier *string `field:"optional" json:"shareIdentifier" yaml:"shareIdentifier"`
	// The weight factor for the fair share identifier.
	//
	// The default value is 1.0. A lower value has a higher priority for compute resources. For example, jobs that use a share identifier with a weight factor of 0.125 (1/8) get 8 times the compute resources of jobs that use a share identifier with a weight factor of 1.
	//
	// The smallest supported value is 0.0001, and the largest supported value is 999.9999.
	WeightFactor *float64 `field:"optional" json:"weightFactor" yaml:"weightFactor"`
}

