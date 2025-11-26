package previewawsimagebuildermixins


// Defines criteria to exclude AMIs from lifecycle actions based on the last time they were used to launch an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lastLaunchedProperty := &LastLaunchedProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-lastlaunched.html
//
type CfnLifecyclePolicyPropsMixin_LastLaunchedProperty struct {
	// Defines the unit of time that the lifecycle policy uses to calculate elapsed time since the last instance launched from the AMI.
	//
	// For example: days, weeks, months, or years.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-lastlaunched.html#cfn-imagebuilder-lifecyclepolicy-lastlaunched-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The integer number of units for the time period.
	//
	// For example `6` (months).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-lastlaunched.html#cfn-imagebuilder-lifecyclepolicy-lastlaunched-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

