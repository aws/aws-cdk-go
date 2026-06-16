package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availabilitySloProperty := &AvailabilitySloProperty{
//   	Target: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-availabilityslo.html
//
type CfnPolicy_AvailabilitySloProperty struct {
	// Availability target percentage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-availabilityslo.html#cfn-resiliencehubv2-policy-availabilityslo-target
	//
	Target *float64 `field:"optional" json:"target" yaml:"target"`
}

