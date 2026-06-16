package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiRegionTargetsProperty := &MultiRegionTargetsProperty{
//   	DisasterRecoveryApproach: jsii.String("disasterRecoveryApproach"),
//   	RpoInMinutes: jsii.Number(123),
//   	RtoInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-multiregiontargets.html
//
type CfnPolicy_MultiRegionTargetsProperty struct {
	// Multi-Region disaster recovery approach.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-multiregiontargets.html#cfn-resiliencehubv2-policy-multiregiontargets-disasterrecoveryapproach
	//
	DisasterRecoveryApproach *string `field:"optional" json:"disasterRecoveryApproach" yaml:"disasterRecoveryApproach"`
	// Recovery Point Objective in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-multiregiontargets.html#cfn-resiliencehubv2-policy-multiregiontargets-rpoinminutes
	//
	RpoInMinutes *float64 `field:"optional" json:"rpoInMinutes" yaml:"rpoInMinutes"`
	// Recovery Time Objective in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-multiregiontargets.html#cfn-resiliencehubv2-policy-multiregiontargets-rtoinminutes
	//
	RtoInMinutes *float64 `field:"optional" json:"rtoInMinutes" yaml:"rtoInMinutes"`
}

