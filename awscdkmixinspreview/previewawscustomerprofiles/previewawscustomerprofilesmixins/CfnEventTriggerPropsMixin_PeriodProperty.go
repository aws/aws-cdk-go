package previewawscustomerprofilesmixins


// Defines a limit and the time period during which it is enforced.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   periodProperty := &PeriodProperty{
//   	MaxInvocationsPerProfile: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   	Unlimited: jsii.Boolean(false),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html
//
type CfnEventTriggerPropsMixin_PeriodProperty struct {
	// The maximum allowed number of destination invocations per profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html#cfn-customerprofiles-eventtrigger-period-maxinvocationsperprofile
	//
	MaxInvocationsPerProfile *float64 `field:"optional" json:"maxInvocationsPerProfile" yaml:"maxInvocationsPerProfile"`
	// The unit of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html#cfn-customerprofiles-eventtrigger-period-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// If set to true, there is no limit on the number of destination invocations per profile.
	//
	// The default is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html#cfn-customerprofiles-eventtrigger-period-unlimited
	//
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
	// The amount of time of the specified unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-period.html#cfn-customerprofiles-eventtrigger-period-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

