package awscognito


// The available verified method a user can use to recover their password when they call `ForgotPassword` .
//
// You can use this setting to define a preferred method when a user has more than one method available. With this setting, SMS doesn't qualify for a valid password recovery mechanism if the user also has SMS multi-factor authentication (MFA) activated. In the absence of this setting, Amazon Cognito uses the legacy behavior to determine the recovery method where SMS is preferred through email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountRecoverySettingProperty := &AccountRecoverySettingProperty{
//   	RecoveryMechanisms: []interface{}{
//   		&RecoveryOptionProperty{
//   			Name: jsii.String("name"),
//   			Priority: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-accountrecoverysetting.html
//
type CfnUserPool_AccountRecoverySettingProperty struct {
	// The list of options and priorities for user message delivery in forgot-password operations.
	//
	// Sets or displays user pool preferences for email or SMS message priority, whether users should fall back to a second delivery method, and whether passwords should only be reset by administrators.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-accountrecoverysetting.html#cfn-cognito-userpool-accountrecoverysetting-recoverymechanisms
	//
	RecoveryMechanisms interface{} `field:"optional" json:"recoveryMechanisms" yaml:"recoveryMechanisms"`
}

