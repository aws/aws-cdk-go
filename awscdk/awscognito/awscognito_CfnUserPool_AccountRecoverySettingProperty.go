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
//   accountRecoverySettingProperty := &accountRecoverySettingProperty{
//   	recoveryMechanisms: []interface{}{
//   		&recoveryOptionProperty{
//   			name: jsii.String("name"),
//   			priority: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnUserPool_AccountRecoverySettingProperty struct {
	// The list of `RecoveryOptionTypes` .
	RecoveryMechanisms interface{} `field:"optional" json:"recoveryMechanisms" yaml:"recoveryMechanisms"`
}

