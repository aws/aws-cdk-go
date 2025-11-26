package previewawsresiliencehubmixins


// Defines a failure policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   failurePolicyProperty := &FailurePolicyProperty{
//   	RpoInSecs: jsii.Number(123),
//   	RtoInSecs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html
//
type CfnResiliencyPolicyPropsMixin_FailurePolicyProperty struct {
	// Recovery Point Objective (RPO) in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html#cfn-resiliencehub-resiliencypolicy-failurepolicy-rpoinsecs
	//
	RpoInSecs *float64 `field:"optional" json:"rpoInSecs" yaml:"rpoInSecs"`
	// Recovery Time Objective (RTO) in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html#cfn-resiliencehub-resiliencypolicy-failurepolicy-rtoinsecs
	//
	RtoInSecs *float64 `field:"optional" json:"rtoInSecs" yaml:"rtoInSecs"`
}

