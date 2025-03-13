package awscustomerprofiles


// Defines a limit and the time period during which it is enforced.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   periodProperty := &PeriodProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//
//   	// the properties below are optional
//   	MaxInvocationsPerProfile: jsii.Number(123),
//   	Unlimited: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html
//
type CfnEventTrigger_PeriodProperty struct {
	// The unit of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html#cfn-customerprofiles-eventtrigger-period-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// The amount of time of the specified unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html#cfn-customerprofiles-eventtrigger-period-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// The maximum allowed number of destination invocations per profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html#cfn-customerprofiles-eventtrigger-period-maxinvocationsperprofile
	//
	MaxInvocationsPerProfile *float64 `field:"optional" json:"maxInvocationsPerProfile" yaml:"maxInvocationsPerProfile"`
	// If set to true, there is no limit on the number of destination invocations per profile.
	//
	// The default is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html#cfn-customerprofiles-eventtrigger-period-unlimited
	//
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

