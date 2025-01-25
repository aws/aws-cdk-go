package awscognito


// A recovery option for a user.
//
// The `AccountRecoverySettingType` data type is an array of this object. Each `RecoveryOptionType` has a priority property that determines whether it is a primary or secondary option.
//
// For example, if `verified_email` has a priority of `1` and `verified_phone_number` has a priority of `2` , your user pool sends account-recovery messages to a verified email address but falls back to an SMS message if the user has a verified phone number. The `admin_only` option prevents self-service account recovery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recoveryOptionProperty := &RecoveryOptionProperty{
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-recoveryoption.html
//
type CfnUserPool_RecoveryOptionProperty struct {
	// The recovery method that this object sets a recovery option for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-recoveryoption.html#cfn-cognito-userpool-recoveryoption-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Your priority preference for using the specified attribute in account recovery.
	//
	// The highest priority is `1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-recoveryoption.html#cfn-cognito-userpool-recoveryoption-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

