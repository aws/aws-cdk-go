package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiAzTargetsProperty := &MultiAzTargetsProperty{
//   	DisasterRecoveryApproach: jsii.String("disasterRecoveryApproach"),
//   	RpoInMinutes: jsii.Number(123),
//   	RtoInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-multiaztargets.html
//
type CfnPolicy_MultiAzTargetsProperty struct {
	// Multi-AZ disaster recovery approach.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-multiaztargets.html#cfn-resiliencehubv2-policy-multiaztargets-disasterrecoveryapproach
	//
	DisasterRecoveryApproach *string `field:"optional" json:"disasterRecoveryApproach" yaml:"disasterRecoveryApproach"`
	// Recovery Point Objective in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-multiaztargets.html#cfn-resiliencehubv2-policy-multiaztargets-rpoinminutes
	//
	RpoInMinutes *float64 `field:"optional" json:"rpoInMinutes" yaml:"rpoInMinutes"`
	// Recovery Time Objective in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-multiaztargets.html#cfn-resiliencehubv2-policy-multiaztargets-rtoinminutes
	//
	RtoInMinutes *float64 `field:"optional" json:"rtoInMinutes" yaml:"rtoInMinutes"`
}

