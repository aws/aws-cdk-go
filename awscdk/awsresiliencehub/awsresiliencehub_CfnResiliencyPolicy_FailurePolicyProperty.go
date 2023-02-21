package awsresiliencehub


// Defines a failure policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failurePolicyProperty := &failurePolicyProperty{
//   	rpoInSecs: jsii.Number(123),
//   	rtoInSecs: jsii.Number(123),
//   }
//
type CfnResiliencyPolicy_FailurePolicyProperty struct {
	// The Recovery Point Objective (RPO), in seconds.
	RpoInSecs *float64 `field:"required" json:"rpoInSecs" yaml:"rpoInSecs"`
	// The Recovery Time Objective (RTO), in seconds.
	RtoInSecs *float64 `field:"required" json:"rtoInSecs" yaml:"rtoInSecs"`
}

