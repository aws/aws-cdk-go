package awsresiliencehub


// Defines a failure policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failurePolicyProperty := &FailurePolicyProperty{
//   	RpoInSecs: jsii.Number(123),
//   	RtoInSecs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html
//
type CfnResiliencyPolicy_FailurePolicyProperty struct {
	// The Recovery Point Objective (RPO), in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html#cfn-resiliencehub-resiliencypolicy-failurepolicy-rpoinsecs
	//
	RpoInSecs *float64 `field:"required" json:"rpoInSecs" yaml:"rpoInSecs"`
	// The Recovery Time Objective (RTO), in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html#cfn-resiliencehub-resiliencypolicy-failurepolicy-rtoinsecs
	//
	RtoInSecs *float64 `field:"required" json:"rtoInSecs" yaml:"rtoInSecs"`
}

