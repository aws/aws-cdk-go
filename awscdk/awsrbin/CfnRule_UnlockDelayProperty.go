package awsrbin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   unlockDelayProperty := &UnlockDelayProperty{
//   	UnlockDelayUnit: jsii.String("unlockDelayUnit"),
//   	UnlockDelayValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-unlockdelay.html
//
type CfnRule_UnlockDelayProperty struct {
	// The unit of time in which to measure the unlock delay.
	//
	// Currently, the unlock delay can be measure only in days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-unlockdelay.html#cfn-rbin-rule-unlockdelay-unlockdelayunit
	//
	UnlockDelayUnit *string `field:"optional" json:"unlockDelayUnit" yaml:"unlockDelayUnit"`
	// The unlock delay period, measured in the unit specified for UnlockDelayUnit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-unlockdelay.html#cfn-rbin-rule-unlockdelay-unlockdelayvalue
	//
	UnlockDelayValue *float64 `field:"optional" json:"unlockDelayValue" yaml:"unlockDelayValue"`
}

