package awscognito


// Use this setting to define which verified available method a user can use to recover their password when they call `ForgotPassword` .
//
// It allows you to define a preferred method when a user has more than one method available. With this setting, SMS does not qualify for a valid password recovery mechanism if the user also has SMS MFA enabled. In the absence of this setting, Cognito uses the legacy behavior to determine the recovery method where SMS is preferred over email.
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
	// The list of `RecoveryOptionTypes` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-accountrecoverysetting.html#cfn-cognito-userpool-accountrecoverysetting-recoverymechanisms
	//
	RecoveryMechanisms interface{} `field:"optional" json:"recoveryMechanisms" yaml:"recoveryMechanisms"`
}

